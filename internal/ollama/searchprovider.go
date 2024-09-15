package ollama

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/godbus/dbus/v5"
	"github.com/xyproto/ollamaclient/v2"
)

const busName = "com.ollama"
const objectPath = "/com/ollama/SearchProvider"
const systemPromt = "Give short 300 character response. Answer with just a list of apps. Suggest what Fedora linux user is trying to find. Recommend relevant applications as a list. Do not use markdown in answer. User request:"

type SearchProvider struct {
	logger       *slog.Logger
	debounceTime time.Duration
	client       *ollamaclient.Config
	cancelFn     context.CancelFunc
	mu           sync.RWMutex
}

func NewSearchProvider(debounceTime time.Duration, modelName string, l *slog.Logger) *SearchProvider {
	client := ollamaclient.New()
	client.ModelName = modelName
	client.Verbose = true
	sp := &SearchProvider{
		client:       client,
		logger:       l,
		debounceTime: debounceTime,
	}

	return sp
}

func (sp *SearchProvider) Serve(conn *dbus.Conn) error {
	reply, err := conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("Failed to request name: %v", err)
	}

	sp.logger.Info(fmt.Sprintf("code=%v", reply))
	conn.Export(sp, objectPath, "org.gnome.Shell.SearchProvider2")

	return nil
}

func (sp *SearchProvider) GetInitialResultSet(terms []string) ([]string, *dbus.Error) {
	sp.logger.Info(fmt.Sprintf("Got terms from D-Bus: %s\n", strings.Join(terms, " ")))

	if sp.cancelFn != nil {
		sp.cancelFn()
	}

	ctx, cancel := context.WithCancel(context.Background())

	sp.mu.Lock()
	sp.cancelFn = cancel
	sp.mu.Unlock()

	select {
	case <-ctx.Done():
		return []string{"Thinking..."}, nil
	case <-time.After(sp.debounceTime):
		return []string{sp.performSearch(terms)}, nil
	}
}

func (sp *SearchProvider) performSearch(terms []string) string {
	sp.logger.Info(fmt.Sprintf("Performing search with Ollama: %s\n", strings.Join(terms, " ")))
	promt := fmt.Sprintf("%s %s", systemPromt, strings.Join(terms, " "))
	result, err := sp.client.GetOutput(promt)

	sp.logger.Info(result)

	if err != nil {
		sp.logger.Info(fmt.Sprintf("failed to perfom search using Ollama Search Provider: %v", err.Error()))
	}

	return result
}

func (sp *SearchProvider) GetSubsearchResultSet(previousResults, terms []string) ([]string, *dbus.Error) {
	return sp.GetInitialResultSet(terms)
}

func (sp *SearchProvider) GetResultMetas(identifiers []string) ([]map[string]dbus.Variant, *dbus.Error) {
	metas := make([]map[string]dbus.Variant, len(identifiers))

	for i, id := range identifiers {
		metas[i] = map[string]dbus.Variant{
			"id":          dbus.MakeVariant(id),
			"name":        dbus.MakeVariant(id),
			"description": dbus.MakeVariant("Search result"),
		}
	}

	return metas, nil
}

func (sp *SearchProvider) ActivateResult(identifier string, terms []string, timestamp uint32) *dbus.Error {
	sp.logger.Info(fmt.Sprintf("Activating result: %s\n", identifier))
	return nil
}

func (sp *SearchProvider) LaunchSearch(terms []string, timestamp uint32) *dbus.Error {
	sp.logger.Info(fmt.Sprintf("Launching search for: %s\n", terms))
	return nil
}
