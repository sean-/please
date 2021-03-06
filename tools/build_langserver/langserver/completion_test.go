package langserver

import (
	"context"
	"strings"
	"testing"

	"core"
	"tools/build_langserver/lsp"

	"github.com/stretchr/testify/assert"
)

func TestCompletionWithCONFIG(t *testing.T) {
	ctx := context.Background()

	err := storeFile(ctx, completionPropURI)
	assert.Equal(t, nil, err)

	// Test completion on CONFIG with no starting character
	items, err := handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 0, Character: 7})
	assert.Equal(t, nil, err)
	assert.Equal(t, len(analyzer.State.Config.TagsToFields()), len(items))
	for _, i := range items {
		assert.Equal(t, i.Kind, lsp.Property)
	}

	// Test completion on CONFIG with 1 starting character
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 1, Character: 8})
	assert.Equal(t, nil, err)
	assert.True(t, len(analyzer.State.Config.TagsToFields()) > len(items))
	assert.True(t, itemInList(items, "JARCAT_TOOL"))
	assert.False(t, itemInList(items, "PLZ_VERSION"))

	// Test completion on CONFIG with a word
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 2, Character: 11})
	assert.Equal(t, nil, err)
	assert.True(t, len(analyzer.State.Config.TagsToFields()) > len(items))
	assert.True(t, itemInList(items, "JAVAC_TOOL"))
	for _, i := range items {
		assert.True(t, strings.Contains(i.Label, "JAVA"))
	}

	// Test completion with assignment
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 3, Character: 18})
	assert.Equal(t, nil, err)
	assert.True(t, len(analyzer.State.Config.TagsToFields()) > len(items))
	assert.True(t, itemInList(items, "JAVAC_TOOL"))
	for _, i := range items {
		assert.True(t, strings.Contains(i.Label, "JAVA"))
	}

	// Test completion on empty line
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 5, Character: 13})
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(items))

	// Test config should be empty
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 4, Character: 14})
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(items))
}

func TestCompletionWithStringMethods(t *testing.T) {
	ctx := context.Background()
	context.Background()

	// Tests completion on no letters follows after dot(.)
	items, err := handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 6, Character: 19})
	assert.Equal(t, nil, err)
	assert.Equal(t, len(analyzer.Attributes["str"]), len(items))
	assert.True(t, itemInList(items, "replace"))
	assert.True(t, itemInList(items, "format"))

	// Test completion with 1 starting character: f
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 7, Character: 20})
	assert.Equal(t, nil, err)
	assert.True(t, itemInList(items, "format"))
	assert.True(t, itemInList(items, "find"))
	assert.True(t, itemInList(items, "rfind"))

	// Test completion with a three letters: for
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 8, Character: 22})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "format", items[0].Label)
	//
	// Test completion with assignment
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 9, Character: 19})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "format", items[0].Label)
}

func TestCompletionWithDictMethods(t *testing.T) {
	ctx := context.Background()

	// Tests completion on no letters follows after dot(.)
	items, err := handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 11, Character: 25})
	assert.Equal(t, nil, err)
	assert.Equal(t, len(analyzer.Attributes["dict"]), len(items))
	assert.True(t, itemInList(items, "get"))
	assert.True(t, itemInList(items, "keys"))
	assert.True(t, itemInList(items, "items"))

	// Test completion with 1 starting character: k
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 12, Character: 16})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "keys", items[0].Label)

	// Test completion with a three letters: get
	items, err = handler.getCompletionItemsList(ctx, completionPropURI, lsp.Position{Line: 13, Character: 18})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "get", items[0].Label)
}

func TestCompletionWithBuildLabels(t *testing.T) {
	ctx := context.Background()

	err := storeFile(ctx, completionLabelURI)
	assert.Equal(t, nil, err)

	items, err := handler.getCompletionItemsList(ctx, completionLabelURI, lsp.Position{Line: 0, Character: 6})
	assert.Equal(t, nil, err)
	assert.True(t, itemInList(items, "src/cache"))
	for _, i := range items {
		assert.True(t, strings.HasPrefix(i.Label, "src"))
	}

	items, err = handler.getCompletionItemsList(ctx, completionLabelURI, lsp.Position{Line: 1, Character: 13})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "query", items[0].Label)
	//assert.Equal(t, items[0].TextEdit)
	t.Log(items[0].TextEdit)

	items, err = handler.getCompletionItemsList(ctx, completionLabelURI, lsp.Position{Line: 2, Character: 14})
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, "query", items[0].Label)
	t.Log(items[0].Label)
}

func TestCompletionWithBuildLabels2(t *testing.T) {
	ctx := context.Background()

	items, err := handler.getCompletionItemsList(ctx, completionLabelURI, lsp.Position{Line: 4, Character: 7})
	assert.Equal(t, nil, err)
	t.Log(items)
}

func TestCompletionIncompleteFile(t *testing.T) {
	//TODO(BNM)
	t.Log(core.LooksLikeABuildLabel("//bkag//bh"))
	//stmt, err := analyzer.AspStatementFromFile(completionURI)
	//t.Log(stmt)
	//t.Log(err)
	t.Log(analyzer.BuildFileURIFromPackage(""))
}

/***************************************
 * Helpers
 ***************************************/
func itemInList(itemList []*lsp.CompletionItem, targetLabel string) bool {
	for _, item := range itemList {
		if item.Label == targetLabel {
			return true
		}
	}
	return false
}

func storeFile(ctx context.Context, uri lsp.DocumentURI) error {
	content, err := ReadFile(ctx, uri)
	if err != nil {
		return err
	}
	text := strings.Join(content, "\n")

	handler.workspace.Store(uri, text)
	return nil
}
