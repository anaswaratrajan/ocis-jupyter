package svc

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	v0proto "github.com/anaswaratrajan/ocis-jupyter/pkg/proto/v0"
	mclient "github.com/micro/go-micro/v2/client"
	olog "github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/ocis-pkg/middleware"
	settings "github.com/owncloud/ocis/settings/pkg/proto/v0"
	ssvc "github.com/owncloud/ocis/settings/pkg/service/v0"
)

var (
	// ErrMissingName defines the error if name is missing.
	ErrMissingName = errors.New("missing a name")

	bundleIDGreeting       = "21fb587b-7b69-4aa6-b0a7-93c74af1918f"
	settingIDGreeterPhrase = "b3584ea8-caec-4951-a2c1-92cbc70071b7"

	// maxRetries indicates how many times to try a request for network reasons.
	maxRetries = 5
)

// NewService returns a service implementation for HelloHandler.
func NewService() v0proto.HelloHandler {
	return Hello{}
}

// Hello defines implements the business logic for HelloHandler.
type Hello struct {
	// Add database handlers here.
}

// Greet implements the HelloHandler interface.
func (s Hello) Greet(ctx context.Context, req *v0proto.GreetRequest, rsp *v0proto.GreetResponse) error {
	if req.Name == "" {
		return ErrMissingName
	}

	phrase := getGreetingPhrase(ctx)
	rsp.Message = fmt.Sprintf(phrase, req.Name)

	return nil
}

func getGreetingPhrase(ctx context.Context) string {
	ownAccountUUID := ctx.Value(middleware.UUIDKey)
	if ownAccountUUID != nil {
		// request to the settings service requires to have the account uuid of the authenticated user available in the context
		rq := settings.GetValueByUniqueIdentifiersRequest{
			AccountUuid: ownAccountUUID.(string),
			SettingId:   settingIDGreeterPhrase,
		}

		// TODO this won't work with a registry other than mdns. Look into Micro's client initialization.
		// https://github.com/owncloud/ocis-hello/issues/74
		valueService := settings.NewValueService("com.owncloud.api.settings", mclient.DefaultClient)
		response, err := valueService.GetValueByUniqueIdentifiers(ctx, &rq)
		if err == nil {
			value, ok := response.Value.Value.Value.(*settings.Value_StringValue)
			if ok {
				trimmedPhrase := strings.Trim(
					value.StringValue,
					" \t",
				)
				if trimmedPhrase != "" {
					return trimmedPhrase + " %s"
				}
			}
		}
	}

	return `<!DOCTYPE html>
	<html>
	<head><meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
	<title>Notebook</title><script src="https://cdnjs.cloudflare.com/ajax/libs/require.js/2.1.10/require.min.js"></script>
	
	
	
	
	<style type="text/css">
		pre { line-height: 125%; }
	td.linenos .normal { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
	span.linenos { color: inherit; background-color: transparent; padding-left: 5px; padding-right: 5px; }
	td.linenos .special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
	span.linenos.special { color: #000000; background-color: #ffffc0; padding-left: 5px; padding-right: 5px; }
	.highlight .hll { background-color: var(--jp-cell-editor-active-background) }
	.highlight { background: var(--jp-cell-editor-background); color: var(--jp-mirror-editor-variable-color) }
	.highlight .c { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment */
	.highlight .err { color: var(--jp-mirror-editor-error-color) } /* Error */
	.highlight .k { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword */
	.highlight .o { color: var(--jp-mirror-editor-operator-color); font-weight: bold } /* Operator */
	.highlight .p { color: var(--jp-mirror-editor-punctuation-color) } /* Punctuation */
	.highlight .ch { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.Hashbang */
	.highlight .cm { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.Multiline */
	.highlight .cp { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.Preproc */
	.highlight .cpf { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.PreprocFile */
	.highlight .c1 { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.Single */
	.highlight .cs { color: var(--jp-mirror-editor-comment-color); font-style: italic } /* Comment.Special */
	.highlight .kc { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Constant */
	.highlight .kd { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Declaration */
	.highlight .kn { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Namespace */
	.highlight .kp { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Pseudo */
	.highlight .kr { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Reserved */
	.highlight .kt { color: var(--jp-mirror-editor-keyword-color); font-weight: bold } /* Keyword.Type */
	.highlight .m { color: var(--jp-mirror-editor-number-color) } /* Literal.Number */
	.highlight .s { color: var(--jp-mirror-editor-string-color) } /* Literal.String */
	.highlight .ow { color: var(--jp-mirror-editor-operator-color); font-weight: bold } /* Operator.Word */
	.highlight .w { color: var(--jp-mirror-editor-variable-color) } /* Text.Whitespace */
	.highlight .mb { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Bin */
	.highlight .mf { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Float */
	.highlight .mh { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Hex */
	.highlight .mi { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Integer */
	.highlight .mo { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Oct */
	.highlight .sa { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Affix */
	.highlight .sb { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Backtick */
	.highlight .sc { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Char */
	.highlight .dl { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Delimiter */
	.highlight .sd { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Doc */
	.highlight .s2 { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Double */
	.highlight .se { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Escape */
	.highlight .sh { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Heredoc */
	.highlight .si { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Interpol */
	.highlight .sx { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Other */
	.highlight .sr { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Regex */
	.highlight .s1 { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Single */
	.highlight .ss { color: var(--jp-mirror-editor-string-color) } /* Literal.String.Symbol */
	.highlight .il { color: var(--jp-mirror-editor-number-color) } /* Literal.Number.Integer.Long */
	  </style>
	
	
	
	<style type="text/css">
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*
	 * Mozilla scrollbar styling
	 */
	
	/* use standard opaque scrollbars for most nodes */
	[data-jp-theme-scrollbars='true'] {
	  scrollbar-color: rgb(var(--jp-scrollbar-thumb-color))
		var(--jp-scrollbar-background-color);
	}
	
	/* for code nodes, use a transparent style of scrollbar. These selectors
	 * will match lower in the tree, and so will override the above */
	[data-jp-theme-scrollbars='true'] .CodeMirror-hscrollbar,
	[data-jp-theme-scrollbars='true'] .CodeMirror-vscrollbar {
	  scrollbar-color: rgba(var(--jp-scrollbar-thumb-color), 0.5) transparent;
	}
	
	/* tiny scrollbar */
	
	.jp-scrollbar-tiny {
	  scrollbar-color: rgba(var(--jp-scrollbar-thumb-color), 0.5) transparent;
	  scrollbar-width: thin;
	}
	
	/*
	 * Webkit scrollbar styling
	 */
	
	/* use standard opaque scrollbars for most nodes */
	
	[data-jp-theme-scrollbars='true'] ::-webkit-scrollbar,
	[data-jp-theme-scrollbars='true'] ::-webkit-scrollbar-corner {
	  background: var(--jp-scrollbar-background-color);
	}
	
	[data-jp-theme-scrollbars='true'] ::-webkit-scrollbar-thumb {
	  background: rgb(var(--jp-scrollbar-thumb-color));
	  border: var(--jp-scrollbar-thumb-margin) solid transparent;
	  background-clip: content-box;
	  border-radius: var(--jp-scrollbar-thumb-radius);
	}
	
	[data-jp-theme-scrollbars='true'] ::-webkit-scrollbar-track:horizontal {
	  border-left: var(--jp-scrollbar-endpad) solid
		var(--jp-scrollbar-background-color);
	  border-right: var(--jp-scrollbar-endpad) solid
		var(--jp-scrollbar-background-color);
	}
	
	[data-jp-theme-scrollbars='true'] ::-webkit-scrollbar-track:vertical {
	  border-top: var(--jp-scrollbar-endpad) solid
		var(--jp-scrollbar-background-color);
	  border-bottom: var(--jp-scrollbar-endpad) solid
		var(--jp-scrollbar-background-color);
	}
	
	/* for code nodes, use a transparent style of scrollbar */
	
	[data-jp-theme-scrollbars='true'] .CodeMirror-hscrollbar::-webkit-scrollbar,
	[data-jp-theme-scrollbars='true'] .CodeMirror-vscrollbar::-webkit-scrollbar,
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-hscrollbar::-webkit-scrollbar-corner,
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-vscrollbar::-webkit-scrollbar-corner {
	  background-color: transparent;
	}
	
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-hscrollbar::-webkit-scrollbar-thumb,
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-vscrollbar::-webkit-scrollbar-thumb {
	  background: rgba(var(--jp-scrollbar-thumb-color), 0.5);
	  border: var(--jp-scrollbar-thumb-margin) solid transparent;
	  background-clip: content-box;
	  border-radius: var(--jp-scrollbar-thumb-radius);
	}
	
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-hscrollbar::-webkit-scrollbar-track:horizontal {
	  border-left: var(--jp-scrollbar-endpad) solid transparent;
	  border-right: var(--jp-scrollbar-endpad) solid transparent;
	}
	
	[data-jp-theme-scrollbars='true']
	  .CodeMirror-vscrollbar::-webkit-scrollbar-track:vertical {
	  border-top: var(--jp-scrollbar-endpad) solid transparent;
	  border-bottom: var(--jp-scrollbar-endpad) solid transparent;
	}
	
	/* tiny scrollbar */
	
	.jp-scrollbar-tiny::-webkit-scrollbar,
	.jp-scrollbar-tiny::-webkit-scrollbar-corner {
	  background-color: transparent;
	  height: 4px;
	  width: 4px;
	}
	
	.jp-scrollbar-tiny::-webkit-scrollbar-thumb {
	  background: rgba(var(--jp-scrollbar-thumb-color), 0.5);
	}
	
	.jp-scrollbar-tiny::-webkit-scrollbar-track:horizontal {
	  border-left: 0px solid transparent;
	  border-right: 0px solid transparent;
	}
	
	.jp-scrollbar-tiny::-webkit-scrollbar-track:vertical {
	  border-top: 0px solid transparent;
	  border-bottom: 0px solid transparent;
	}
	
	/*
	 * Phosphor
	 */
	
	.lm-ScrollBar[data-orientation='horizontal'] {
	  min-height: 16px;
	  max-height: 16px;
	  min-width: 45px;
	  border-top: 1px solid #a0a0a0;
	}
	
	.lm-ScrollBar[data-orientation='vertical'] {
	  min-width: 16px;
	  max-width: 16px;
	  min-height: 45px;
	  border-left: 1px solid #a0a0a0;
	}
	
	.lm-ScrollBar-button {
	  background-color: #f0f0f0;
	  background-position: center center;
	  min-height: 15px;
	  max-height: 15px;
	  min-width: 15px;
	  max-width: 15px;
	}
	
	.lm-ScrollBar-button:hover {
	  background-color: #dadada;
	}
	
	.lm-ScrollBar-button.lm-mod-active {
	  background-color: #cdcdcd;
	}
	
	.lm-ScrollBar-track {
	  background: #f0f0f0;
	}
	
	.lm-ScrollBar-thumb {
	  background: #cdcdcd;
	}
	
	.lm-ScrollBar-thumb:hover {
	  background: #bababa;
	}
	
	.lm-ScrollBar-thumb.lm-mod-active {
	  background: #a0a0a0;
	}
	
	.lm-ScrollBar[data-orientation='horizontal'] .lm-ScrollBar-thumb {
	  height: 100%;
	  min-width: 15px;
	  border-left: 1px solid #a0a0a0;
	  border-right: 1px solid #a0a0a0;
	}
	
	.lm-ScrollBar[data-orientation='vertical'] .lm-ScrollBar-thumb {
	  width: 100%;
	  min-height: 15px;
	  border-top: 1px solid #a0a0a0;
	  border-bottom: 1px solid #a0a0a0;
	}
	
	.lm-ScrollBar[data-orientation='horizontal']
	  .lm-ScrollBar-button[data-action='decrement'] {
	  background-image: var(--jp-icon-caret-left);
	  background-size: 17px;
	}
	
	.lm-ScrollBar[data-orientation='horizontal']
	  .lm-ScrollBar-button[data-action='increment'] {
	  background-image: var(--jp-icon-caret-right);
	  background-size: 17px;
	}
	
	.lm-ScrollBar[data-orientation='vertical']
	  .lm-ScrollBar-button[data-action='decrement'] {
	  background-image: var(--jp-icon-caret-up);
	  background-size: 17px;
	}
	
	.lm-ScrollBar[data-orientation='vertical']
	  .lm-ScrollBar-button[data-action='increment'] {
	  background-image: var(--jp-icon-caret-down);
	  background-size: 17px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-Widget, /* </DEPRECATED> */
	.lm-Widget {
	  box-sizing: border-box;
	  position: relative;
	  overflow: hidden;
	  cursor: default;
	}
	
	
	/* <DEPRECATED> */ .p-Widget.p-mod-hidden, /* </DEPRECATED> */
	.lm-Widget.lm-mod-hidden {
	  display: none !important;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-CommandPalette, /* </DEPRECATED> */
	.lm-CommandPalette {
	  display: flex;
	  flex-direction: column;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-search, /* </DEPRECATED> */
	.lm-CommandPalette-search {
	  flex: 0 0 auto;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-content, /* </DEPRECATED> */
	.lm-CommandPalette-content {
	  flex: 1 1 auto;
	  margin: 0;
	  padding: 0;
	  min-height: 0;
	  overflow: auto;
	  list-style-type: none;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-header, /* </DEPRECATED> */
	.lm-CommandPalette-header {
	  overflow: hidden;
	  white-space: nowrap;
	  text-overflow: ellipsis;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-item, /* </DEPRECATED> */
	.lm-CommandPalette-item {
	  display: flex;
	  flex-direction: row;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-itemIcon, /* </DEPRECATED> */
	.lm-CommandPalette-itemIcon {
	  flex: 0 0 auto;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-itemContent, /* </DEPRECATED> */
	.lm-CommandPalette-itemContent {
	  flex: 1 1 auto;
	  overflow: hidden;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-itemShortcut, /* </DEPRECATED> */
	.lm-CommandPalette-itemShortcut {
	  flex: 0 0 auto;
	}
	
	
	/* <DEPRECATED> */ .p-CommandPalette-itemLabel, /* </DEPRECATED> */
	.lm-CommandPalette-itemLabel {
	  overflow: hidden;
	  white-space: nowrap;
	  text-overflow: ellipsis;
	}
	
	.lm-close-icon {
		border:1px solid transparent;
	  background-color: transparent;
	  position: absolute;
		z-index:1;
		right:3%;
		top: 0;
		bottom: 0;
		margin: auto;
		padding: 7px 0;
		display: none;
		vertical-align: middle;
	  outline: 0;
	  cursor: pointer;
	}
	.lm-close-icon:after {
		content: "X";
		display: block;
		width: 15px;
		height: 15px;
		text-align: center;
		color:#000;
		font-weight: normal;
		font-size: 12px;
		cursor: pointer;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-DockPanel, /* </DEPRECATED> */
	.lm-DockPanel {
	  z-index: 0;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-widget, /* </DEPRECATED> */
	.lm-DockPanel-widget {
	  z-index: 0;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-tabBar, /* </DEPRECATED> */
	.lm-DockPanel-tabBar {
	  z-index: 1;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-handle, /* </DEPRECATED> */
	.lm-DockPanel-handle {
	  z-index: 2;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-handle.p-mod-hidden, /* </DEPRECATED> */
	.lm-DockPanel-handle.lm-mod-hidden {
	  display: none !important;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-handle:after, /* </DEPRECATED> */
	.lm-DockPanel-handle:after {
	  position: absolute;
	  top: 0;
	  left: 0;
	  width: 100%;
	  height: 100%;
	  content: '';
	}
	
	
	/* <DEPRECATED> */
	.p-DockPanel-handle[data-orientation='horizontal'],
	/* </DEPRECATED> */
	.lm-DockPanel-handle[data-orientation='horizontal'] {
	  cursor: ew-resize;
	}
	
	
	/* <DEPRECATED> */
	.p-DockPanel-handle[data-orientation='vertical'],
	/* </DEPRECATED> */
	.lm-DockPanel-handle[data-orientation='vertical'] {
	  cursor: ns-resize;
	}
	
	
	/* <DEPRECATED> */
	.p-DockPanel-handle[data-orientation='horizontal']:after,
	/* </DEPRECATED> */
	.lm-DockPanel-handle[data-orientation='horizontal']:after {
	  left: 50%;
	  min-width: 8px;
	  transform: translateX(-50%);
	}
	
	
	/* <DEPRECATED> */
	.p-DockPanel-handle[data-orientation='vertical']:after,
	/* </DEPRECATED> */
	.lm-DockPanel-handle[data-orientation='vertical']:after {
	  top: 50%;
	  min-height: 8px;
	  transform: translateY(-50%);
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-overlay, /* </DEPRECATED> */
	.lm-DockPanel-overlay {
	  z-index: 3;
	  box-sizing: border-box;
	  pointer-events: none;
	}
	
	
	/* <DEPRECATED> */ .p-DockPanel-overlay.p-mod-hidden, /* </DEPRECATED> */
	.lm-DockPanel-overlay.lm-mod-hidden {
	  display: none !important;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-Menu, /* </DEPRECATED> */
	.lm-Menu {
	  z-index: 10000;
	  position: absolute;
	  white-space: nowrap;
	  overflow-x: hidden;
	  overflow-y: auto;
	  outline: none;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	
	/* <DEPRECATED> */ .p-Menu-content, /* </DEPRECATED> */
	.lm-Menu-content {
	  margin: 0;
	  padding: 0;
	  display: table;
	  list-style-type: none;
	}
	
	
	/* <DEPRECATED> */ .p-Menu-item, /* </DEPRECATED> */
	.lm-Menu-item {
	  display: table-row;
	}
	
	
	/* <DEPRECATED> */
	.p-Menu-item.p-mod-hidden,
	.p-Menu-item.p-mod-collapsed,
	/* </DEPRECATED> */
	.lm-Menu-item.lm-mod-hidden,
	.lm-Menu-item.lm-mod-collapsed {
	  display: none !important;
	}
	
	
	/* <DEPRECATED> */
	.p-Menu-itemIcon,
	.p-Menu-itemSubmenuIcon,
	/* </DEPRECATED> */
	.lm-Menu-itemIcon,
	.lm-Menu-itemSubmenuIcon {
	  display: table-cell;
	  text-align: center;
	}
	
	
	/* <DEPRECATED> */ .p-Menu-itemLabel, /* </DEPRECATED> */
	.lm-Menu-itemLabel {
	  display: table-cell;
	  text-align: left;
	}
	
	
	/* <DEPRECATED> */ .p-Menu-itemShortcut, /* </DEPRECATED> */
	.lm-Menu-itemShortcut {
	  display: table-cell;
	  text-align: right;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-MenuBar, /* </DEPRECATED> */
	.lm-MenuBar {
	  outline: none;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	
	/* <DEPRECATED> */ .p-MenuBar-content, /* </DEPRECATED> */
	.lm-MenuBar-content {
	  margin: 0;
	  padding: 0;
	  display: flex;
	  flex-direction: row;
	  list-style-type: none;
	}
	
	
	/* <DEPRECATED> */ .p--MenuBar-item, /* </DEPRECATED> */
	.lm-MenuBar-item {
	  box-sizing: border-box;
	}
	
	
	/* <DEPRECATED> */
	.p-MenuBar-itemIcon,
	.p-MenuBar-itemLabel,
	/* </DEPRECATED> */
	.lm-MenuBar-itemIcon,
	.lm-MenuBar-itemLabel {
	  display: inline-block;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-ScrollBar, /* </DEPRECATED> */
	.lm-ScrollBar {
	  display: flex;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	
	/* <DEPRECATED> */
	.p-ScrollBar[data-orientation='horizontal'],
	/* </DEPRECATED> */
	.lm-ScrollBar[data-orientation='horizontal'] {
	  flex-direction: row;
	}
	
	
	/* <DEPRECATED> */
	.p-ScrollBar[data-orientation='vertical'],
	/* </DEPRECATED> */
	.lm-ScrollBar[data-orientation='vertical'] {
	  flex-direction: column;
	}
	
	
	/* <DEPRECATED> */ .p-ScrollBar-button, /* </DEPRECATED> */
	.lm-ScrollBar-button {
	  box-sizing: border-box;
	  flex: 0 0 auto;
	}
	
	
	/* <DEPRECATED> */ .p-ScrollBar-track, /* </DEPRECATED> */
	.lm-ScrollBar-track {
	  box-sizing: border-box;
	  position: relative;
	  overflow: hidden;
	  flex: 1 1 auto;
	}
	
	
	/* <DEPRECATED> */ .p-ScrollBar-thumb, /* </DEPRECATED> */
	.lm-ScrollBar-thumb {
	  box-sizing: border-box;
	  position: absolute;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-SplitPanel-child, /* </DEPRECATED> */
	.lm-SplitPanel-child {
	  z-index: 0;
	}
	
	
	/* <DEPRECATED> */ .p-SplitPanel-handle, /* </DEPRECATED> */
	.lm-SplitPanel-handle {
	  z-index: 1;
	}
	
	
	/* <DEPRECATED> */ .p-SplitPanel-handle.p-mod-hidden, /* </DEPRECATED> */
	.lm-SplitPanel-handle.lm-mod-hidden {
	  display: none !important;
	}
	
	
	/* <DEPRECATED> */ .p-SplitPanel-handle:after, /* </DEPRECATED> */
	.lm-SplitPanel-handle:after {
	  position: absolute;
	  top: 0;
	  left: 0;
	  width: 100%;
	  height: 100%;
	  content: '';
	}
	
	
	/* <DEPRECATED> */
	.p-SplitPanel[data-orientation='horizontal'] > .p-SplitPanel-handle,
	/* </DEPRECATED> */
	.lm-SplitPanel[data-orientation='horizontal'] > .lm-SplitPanel-handle {
	  cursor: ew-resize;
	}
	
	
	/* <DEPRECATED> */
	.p-SplitPanel[data-orientation='vertical'] > .p-SplitPanel-handle,
	/* </DEPRECATED> */
	.lm-SplitPanel[data-orientation='vertical'] > .lm-SplitPanel-handle {
	  cursor: ns-resize;
	}
	
	
	/* <DEPRECATED> */
	.p-SplitPanel[data-orientation='horizontal'] > .p-SplitPanel-handle:after,
	/* </DEPRECATED> */
	.lm-SplitPanel[data-orientation='horizontal'] > .lm-SplitPanel-handle:after {
	  left: 50%;
	  min-width: 8px;
	  transform: translateX(-50%);
	}
	
	
	/* <DEPRECATED> */
	.p-SplitPanel[data-orientation='vertical'] > .p-SplitPanel-handle:after,
	/* </DEPRECATED> */
	.lm-SplitPanel[data-orientation='vertical'] > .lm-SplitPanel-handle:after {
	  top: 50%;
	  min-height: 8px;
	  transform: translateY(-50%);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-TabBar, /* </DEPRECATED> */
	.lm-TabBar {
	  display: flex;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar[data-orientation='horizontal'], /* </DEPRECATED> */
	.lm-TabBar[data-orientation='horizontal'] {
	  flex-direction: row;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar[data-orientation='vertical'], /* </DEPRECATED> */
	.lm-TabBar[data-orientation='vertical'] {
	  flex-direction: column;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar-content, /* </DEPRECATED> */
	.lm-TabBar-content {
	  margin: 0;
	  padding: 0;
	  display: flex;
	  flex: 1 1 auto;
	  list-style-type: none;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar[data-orientation='horizontal'] > .p-TabBar-content,
	/* </DEPRECATED> */
	.lm-TabBar[data-orientation='horizontal'] > .lm-TabBar-content {
	  flex-direction: row;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar[data-orientation='vertical'] > .p-TabBar-content,
	/* </DEPRECATED> */
	.lm-TabBar[data-orientation='vertical'] > .lm-TabBar-content {
	  flex-direction: column;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar-tab, /* </DEPRECATED> */
	.lm-TabBar-tab {
	  display: flex;
	  flex-direction: row;
	  box-sizing: border-box;
	  overflow: hidden;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar-tabIcon,
	.p-TabBar-tabCloseIcon,
	/* </DEPRECATED> */
	.lm-TabBar-tabIcon,
	.lm-TabBar-tabCloseIcon {
	  flex: 0 0 auto;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar-tabLabel, /* </DEPRECATED> */
	.lm-TabBar-tabLabel {
	  flex: 1 1 auto;
	  overflow: hidden;
	  white-space: nowrap;
	}
	
	
	.lm-TabBar-tabInput {
	  user-select: all;
	  width: 100%;
	  box-sizing : border-box;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar-tab.p-mod-hidden, /* </DEPRECATED> */
	.lm-TabBar-tab.lm-mod-hidden {
	  display: none !important;
	}
	
	
	/* <DEPRECATED> */ .p-TabBar.p-mod-dragging .p-TabBar-tab, /* </DEPRECATED> */
	.lm-TabBar.lm-mod-dragging .lm-TabBar-tab {
	  position: relative;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar.p-mod-dragging[data-orientation='horizontal'] .p-TabBar-tab,
	/* </DEPRECATED> */
	.lm-TabBar.lm-mod-dragging[data-orientation='horizontal'] .lm-TabBar-tab {
	  left: 0;
	  transition: left 150ms ease;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar.p-mod-dragging[data-orientation='vertical'] .p-TabBar-tab,
	/* </DEPRECATED> */
	.lm-TabBar.lm-mod-dragging[data-orientation='vertical'] .lm-TabBar-tab {
	  top: 0;
	  transition: top 150ms ease;
	}
	
	
	/* <DEPRECATED> */
	.p-TabBar.p-mod-dragging .p-TabBar-tab.p-mod-dragging,
	/* </DEPRECATED> */
	.lm-TabBar.lm-mod-dragging .lm-TabBar-tab.lm-mod-dragging {
	  transition: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ .p-TabPanel-tabBar, /* </DEPRECATED> */
	.lm-TabPanel-tabBar {
	  z-index: 1;
	}
	
	
	/* <DEPRECATED> */ .p-TabPanel-stackedPanel, /* </DEPRECATED> */
	.lm-TabPanel-stackedPanel {
	  z-index: 0;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	@charset "UTF-8";
	html{
	  -webkit-box-sizing:border-box;
			  box-sizing:border-box; }
	
	*,
	*::before,
	*::after{
	  -webkit-box-sizing:inherit;
			  box-sizing:inherit; }
	
	body{
	  font-size:14px;
	  font-weight:400;
	  letter-spacing:0;
	  line-height:1.28581;
	  text-transform:none;
	  color:#182026;
	  font-family:-apple-system, "BlinkMacSystemFont", "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Open Sans", "Helvetica Neue", "Icons16", sans-serif; }
	
	p{
	  margin-bottom:10px;
	  margin-top:0; }
	
	small{
	  font-size:12px; }
	
	strong{
	  font-weight:600; }
	
	::-moz-selection{
	  background:rgba(125, 188, 255, 0.6); }
	
	::selection{
	  background:rgba(125, 188, 255, 0.6); }
	.bp3-heading{
	  color:#182026;
	  font-weight:600;
	  margin:0 0 10px;
	  padding:0; }
	  .bp3-dark .bp3-heading{
		color:#f5f8fa; }
	
	h1.bp3-heading, .bp3-running-text h1{
	  font-size:36px;
	  line-height:40px; }
	
	h2.bp3-heading, .bp3-running-text h2{
	  font-size:28px;
	  line-height:32px; }
	
	h3.bp3-heading, .bp3-running-text h3{
	  font-size:22px;
	  line-height:25px; }
	
	h4.bp3-heading, .bp3-running-text h4{
	  font-size:18px;
	  line-height:21px; }
	
	h5.bp3-heading, .bp3-running-text h5{
	  font-size:16px;
	  line-height:19px; }
	
	h6.bp3-heading, .bp3-running-text h6{
	  font-size:14px;
	  line-height:16px; }
	.bp3-ui-text{
	  font-size:14px;
	  font-weight:400;
	  letter-spacing:0;
	  line-height:1.28581;
	  text-transform:none; }
	
	.bp3-monospace-text{
	  font-family:monospace;
	  text-transform:none; }
	
	.bp3-text-muted{
	  color:#5c7080; }
	  .bp3-dark .bp3-text-muted{
		color:#a7b6c2; }
	
	.bp3-text-disabled{
	  color:rgba(92, 112, 128, 0.6); }
	  .bp3-dark .bp3-text-disabled{
		color:rgba(167, 182, 194, 0.6); }
	
	.bp3-text-overflow-ellipsis{
	  overflow:hidden;
	  text-overflow:ellipsis;
	  white-space:nowrap;
	  word-wrap:normal; }
	.bp3-running-text{
	  font-size:14px;
	  line-height:1.5; }
	  .bp3-running-text h1{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h1{
		  color:#f5f8fa; }
	  .bp3-running-text h2{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h2{
		  color:#f5f8fa; }
	  .bp3-running-text h3{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h3{
		  color:#f5f8fa; }
	  .bp3-running-text h4{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h4{
		  color:#f5f8fa; }
	  .bp3-running-text h5{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h5{
		  color:#f5f8fa; }
	  .bp3-running-text h6{
		color:#182026;
		font-weight:600;
		margin-bottom:20px;
		margin-top:40px; }
		.bp3-dark .bp3-running-text h6{
		  color:#f5f8fa; }
	  .bp3-running-text hr{
		border:none;
		border-bottom:1px solid rgba(16, 22, 26, 0.15);
		margin:20px 0; }
		.bp3-dark .bp3-running-text hr{
		  border-color:rgba(255, 255, 255, 0.15); }
	  .bp3-running-text p{
		margin:0 0 10px;
		padding:0; }
	
	.bp3-text-large{
	  font-size:16px; }
	
	.bp3-text-small{
	  font-size:12px; }
	a{
	  color:#106ba3;
	  text-decoration:none; }
	  a:hover{
		color:#106ba3;
		cursor:pointer;
		text-decoration:underline; }
	  a .bp3-icon, a .bp3-icon-standard, a .bp3-icon-large{
		color:inherit; }
	  a code,
	  .bp3-dark a code{
		color:inherit; }
	  .bp3-dark a,
	  .bp3-dark a:hover{
		color:#48aff0; }
		.bp3-dark a .bp3-icon, .bp3-dark a .bp3-icon-standard, .bp3-dark a .bp3-icon-large,
		.bp3-dark a:hover .bp3-icon,
		.bp3-dark a:hover .bp3-icon-standard,
		.bp3-dark a:hover .bp3-icon-large{
		  color:inherit; }
	.bp3-running-text code, .bp3-code{
	  font-family:monospace;
	  text-transform:none;
	  background:rgba(255, 255, 255, 0.7);
	  border-radius:3px;
	  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2);
			  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2);
	  color:#5c7080;
	  font-size:smaller;
	  padding:2px 5px; }
	  .bp3-dark .bp3-running-text code, .bp3-running-text .bp3-dark code, .bp3-dark .bp3-code{
		background:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
		color:#a7b6c2; }
	  .bp3-running-text a > code, a > .bp3-code{
		color:#137cbd; }
		.bp3-dark .bp3-running-text a > code, .bp3-running-text .bp3-dark a > code, .bp3-dark a > .bp3-code{
		  color:inherit; }
	
	.bp3-running-text pre, .bp3-code-block{
	  font-family:monospace;
	  text-transform:none;
	  background:rgba(255, 255, 255, 0.7);
	  border-radius:3px;
	  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15);
			  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15);
	  color:#182026;
	  display:block;
	  font-size:13px;
	  line-height:1.4;
	  margin:10px 0;
	  padding:13px 15px 12px;
	  word-break:break-all;
	  word-wrap:break-word; }
	  .bp3-dark .bp3-running-text pre, .bp3-running-text .bp3-dark pre, .bp3-dark .bp3-code-block{
		background:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
	  .bp3-running-text pre > code, .bp3-code-block > code{
		background:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		color:inherit;
		font-size:inherit;
		padding:0; }
	
	.bp3-running-text kbd, .bp3-key{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  background:#ffffff;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
	  color:#5c7080;
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex;
	  font-family:inherit;
	  font-size:12px;
	  height:24px;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  line-height:24px;
	  min-width:24px;
	  padding:3px 6px;
	  vertical-align:middle; }
	  .bp3-running-text kbd .bp3-icon, .bp3-key .bp3-icon, .bp3-running-text kbd .bp3-icon-standard, .bp3-key .bp3-icon-standard, .bp3-running-text kbd .bp3-icon-large, .bp3-key .bp3-icon-large{
		margin-right:5px; }
	  .bp3-dark .bp3-running-text kbd, .bp3-running-text .bp3-dark kbd, .bp3-dark .bp3-key{
		background:#394b59;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
		color:#a7b6c2; }
	.bp3-running-text blockquote, .bp3-blockquote{
	  border-left:solid 4px rgba(167, 182, 194, 0.5);
	  margin:0 0 10px;
	  padding:0 20px; }
	  .bp3-dark .bp3-running-text blockquote, .bp3-running-text .bp3-dark blockquote, .bp3-dark .bp3-blockquote{
		border-color:rgba(115, 134, 148, 0.5); }
	.bp3-running-text ul,
	.bp3-running-text ol, .bp3-list{
	  margin:10px 0;
	  padding-left:30px; }
	  .bp3-running-text ul li:not(:last-child), .bp3-running-text ol li:not(:last-child), .bp3-list li:not(:last-child){
		margin-bottom:5px; }
	  .bp3-running-text ul ol, .bp3-running-text ol ol, .bp3-list ol,
	  .bp3-running-text ul ul,
	  .bp3-running-text ol ul,
	  .bp3-list ul{
		margin-top:5px; }
	
	.bp3-list-unstyled{
	  list-style:none;
	  margin:0;
	  padding:0; }
	  .bp3-list-unstyled li{
		padding:0; }
	.bp3-rtl{
	  text-align:right; }
	
	.bp3-dark{
	  color:#f5f8fa; }
	
	:focus{
	  outline:rgba(19, 124, 189, 0.6) auto 2px;
	  outline-offset:2px;
	  -moz-outline-radius:6px; }
	
	.bp3-focus-disabled :focus{
	  outline:none !important; }
	  .bp3-focus-disabled :focus ~ .bp3-control-indicator{
		outline:none !important; }
	
	.bp3-alert{
	  max-width:400px;
	  padding:20px; }
	
	.bp3-alert-body{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex; }
	  .bp3-alert-body .bp3-icon{
		font-size:40px;
		margin-right:20px;
		margin-top:0; }
	
	.bp3-alert-contents{
	  word-break:break-word; }
	
	.bp3-alert-footer{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:reverse;
		  -ms-flex-direction:row-reverse;
			  flex-direction:row-reverse;
	  margin-top:10px; }
	  .bp3-alert-footer .bp3-button{
		margin-left:10px; }
	.bp3-breadcrumbs{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  cursor:default;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -ms-flex-wrap:wrap;
		  flex-wrap:wrap;
	  height:30px;
	  list-style:none;
	  margin:0;
	  padding:0; }
	  .bp3-breadcrumbs > li{
		-webkit-box-align:center;
			-ms-flex-align:center;
				align-items:center;
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex; }
		.bp3-breadcrumbs > li::after{
		  background:url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'%3e%3cpath fill-rule='evenodd' clip-rule='evenodd' d='M10.71 7.29l-4-4a1.003 1.003 0 00-1.42 1.42L8.59 8 5.3 11.29c-.19.18-.3.43-.3.71a1.003 1.003 0 001.71.71l4-4c.18-.18.29-.43.29-.71 0-.28-.11-.53-.29-.71z' fill='%235C7080'/%3e%3c/svg%3e");
		  content:"";
		  display:block;
		  height:16px;
		  margin:0 5px;
		  width:16px; }
		.bp3-breadcrumbs > li:last-of-type::after{
		  display:none; }
	
	.bp3-breadcrumb,
	.bp3-breadcrumb-current,
	.bp3-breadcrumbs-collapsed{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex;
	  font-size:16px; }
	
	.bp3-breadcrumb,
	.bp3-breadcrumbs-collapsed{
	  color:#5c7080; }
	
	.bp3-breadcrumb:hover{
	  text-decoration:none; }
	
	.bp3-breadcrumb.bp3-disabled{
	  color:rgba(92, 112, 128, 0.6);
	  cursor:not-allowed; }
	
	.bp3-breadcrumb .bp3-icon{
	  margin-right:5px; }
	
	.bp3-breadcrumb-current{
	  color:inherit;
	  font-weight:600; }
	  .bp3-breadcrumb-current .bp3-input{
		font-size:inherit;
		font-weight:inherit;
		vertical-align:baseline; }
	
	.bp3-breadcrumbs-collapsed{
	  background:#ced9e0;
	  border:none;
	  border-radius:3px;
	  cursor:pointer;
	  margin-right:2px;
	  padding:1px 5px;
	  vertical-align:text-bottom; }
	  .bp3-breadcrumbs-collapsed::before{
		background:url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'%3e%3cg fill='%235C7080'%3e%3ccircle cx='2' cy='8.03' r='2'/%3e%3ccircle cx='14' cy='8.03' r='2'/%3e%3ccircle cx='8' cy='8.03' r='2'/%3e%3c/g%3e%3c/svg%3e") center no-repeat;
		content:"";
		display:block;
		height:16px;
		width:16px; }
	  .bp3-breadcrumbs-collapsed:hover{
		background:#bfccd6;
		color:#182026;
		text-decoration:none; }
	
	.bp3-dark .bp3-breadcrumb,
	.bp3-dark .bp3-breadcrumbs-collapsed{
	  color:#a7b6c2; }
	
	.bp3-dark .bp3-breadcrumbs > li::after{
	  color:#a7b6c2; }
	
	.bp3-dark .bp3-breadcrumb.bp3-disabled{
	  color:rgba(167, 182, 194, 0.6); }
	
	.bp3-dark .bp3-breadcrumb-current{
	  color:#f5f8fa; }
	
	.bp3-dark .bp3-breadcrumbs-collapsed{
	  background:rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-breadcrumbs-collapsed:hover{
		background:rgba(16, 22, 26, 0.6);
		color:#f5f8fa; }
	.bp3-button{
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  border:none;
	  border-radius:3px;
	  cursor:pointer;
	  font-size:14px;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  padding:5px 10px;
	  text-align:left;
	  vertical-align:middle;
	  min-height:30px;
	  min-width:30px; }
	  .bp3-button > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-button > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-button::before,
	  .bp3-button > *{
		margin-right:7px; }
	  .bp3-button:empty::before,
	  .bp3-button > :last-child{
		margin-right:0; }
	  .bp3-button:empty{
		padding:0 !important; }
	  .bp3-button:disabled, .bp3-button.bp3-disabled{
		cursor:not-allowed; }
	  .bp3-button.bp3-fill{
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		width:100%; }
	  .bp3-button.bp3-align-right,
	  .bp3-align-right .bp3-button{
		text-align:right; }
	  .bp3-button.bp3-align-left,
	  .bp3-align-left .bp3-button{
		text-align:left; }
	  .bp3-button:not([class*="bp3-intent-"]){
		background-color:#f5f8fa;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.8)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
		color:#182026; }
		.bp3-button:not([class*="bp3-intent-"]):hover{
		  background-clip:padding-box;
		  background-color:#ebf1f5;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
		.bp3-button:not([class*="bp3-intent-"]):active, .bp3-button:not([class*="bp3-intent-"]).bp3-active{
		  background-color:#d8e1e8;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-button:not([class*="bp3-intent-"]):disabled, .bp3-button:not([class*="bp3-intent-"]).bp3-disabled{
		  background-color:rgba(206, 217, 224, 0.5);
		  background-image:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed;
		  outline:none; }
		  .bp3-button:not([class*="bp3-intent-"]):disabled.bp3-active, .bp3-button:not([class*="bp3-intent-"]):disabled.bp3-active:hover, .bp3-button:not([class*="bp3-intent-"]).bp3-disabled.bp3-active, .bp3-button:not([class*="bp3-intent-"]).bp3-disabled.bp3-active:hover{
			background:rgba(206, 217, 224, 0.7); }
	  .bp3-button.bp3-intent-primary{
		background-color:#137cbd;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
		.bp3-button.bp3-intent-primary:hover, .bp3-button.bp3-intent-primary:active, .bp3-button.bp3-intent-primary.bp3-active{
		  color:#ffffff; }
		.bp3-button.bp3-intent-primary:hover{
		  background-color:#106ba3;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-primary:active, .bp3-button.bp3-intent-primary.bp3-active{
		  background-color:#0e5a8a;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-primary:disabled, .bp3-button.bp3-intent-primary.bp3-disabled{
		  background-color:rgba(19, 124, 189, 0.5);
		  background-image:none;
		  border-color:transparent;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(255, 255, 255, 0.6); }
	  .bp3-button.bp3-intent-success{
		background-color:#0f9960;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
		.bp3-button.bp3-intent-success:hover, .bp3-button.bp3-intent-success:active, .bp3-button.bp3-intent-success.bp3-active{
		  color:#ffffff; }
		.bp3-button.bp3-intent-success:hover{
		  background-color:#0d8050;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-success:active, .bp3-button.bp3-intent-success.bp3-active{
		  background-color:#0a6640;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-success:disabled, .bp3-button.bp3-intent-success.bp3-disabled{
		  background-color:rgba(15, 153, 96, 0.5);
		  background-image:none;
		  border-color:transparent;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(255, 255, 255, 0.6); }
	  .bp3-button.bp3-intent-warning{
		background-color:#d9822b;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
		.bp3-button.bp3-intent-warning:hover, .bp3-button.bp3-intent-warning:active, .bp3-button.bp3-intent-warning.bp3-active{
		  color:#ffffff; }
		.bp3-button.bp3-intent-warning:hover{
		  background-color:#bf7326;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-warning:active, .bp3-button.bp3-intent-warning.bp3-active{
		  background-color:#a66321;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-warning:disabled, .bp3-button.bp3-intent-warning.bp3-disabled{
		  background-color:rgba(217, 130, 43, 0.5);
		  background-image:none;
		  border-color:transparent;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(255, 255, 255, 0.6); }
	  .bp3-button.bp3-intent-danger{
		background-color:#db3737;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
		.bp3-button.bp3-intent-danger:hover, .bp3-button.bp3-intent-danger:active, .bp3-button.bp3-intent-danger.bp3-active{
		  color:#ffffff; }
		.bp3-button.bp3-intent-danger:hover{
		  background-color:#c23030;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-danger:active, .bp3-button.bp3-intent-danger.bp3-active{
		  background-color:#a82a2a;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-button.bp3-intent-danger:disabled, .bp3-button.bp3-intent-danger.bp3-disabled{
		  background-color:rgba(219, 55, 55, 0.5);
		  background-image:none;
		  border-color:transparent;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(255, 255, 255, 0.6); }
	  .bp3-button[class*="bp3-intent-"] .bp3-button-spinner .bp3-spinner-head{
		stroke:#ffffff; }
	  .bp3-button.bp3-large,
	  .bp3-large .bp3-button{
		min-height:40px;
		min-width:40px;
		font-size:16px;
		padding:5px 15px; }
		.bp3-button.bp3-large::before,
		.bp3-button.bp3-large > *,
		.bp3-large .bp3-button::before,
		.bp3-large .bp3-button > *{
		  margin-right:10px; }
		.bp3-button.bp3-large:empty::before,
		.bp3-button.bp3-large > :last-child,
		.bp3-large .bp3-button:empty::before,
		.bp3-large .bp3-button > :last-child{
		  margin-right:0; }
	  .bp3-button.bp3-small,
	  .bp3-small .bp3-button{
		min-height:24px;
		min-width:24px;
		padding:0 7px; }
	  .bp3-button.bp3-loading{
		position:relative; }
		.bp3-button.bp3-loading[class*="bp3-icon-"]::before{
		  visibility:hidden; }
		.bp3-button.bp3-loading .bp3-button-spinner{
		  margin:0;
		  position:absolute; }
		.bp3-button.bp3-loading > :not(.bp3-button-spinner){
		  visibility:hidden; }
	  .bp3-button[class*="bp3-icon-"]::before{
		font-family:"Icons16", sans-serif;
		font-size:16px;
		font-style:normal;
		font-weight:400;
		line-height:1;
		-moz-osx-font-smoothing:grayscale;
		-webkit-font-smoothing:antialiased;
		color:#5c7080; }
	  .bp3-button .bp3-icon, .bp3-button .bp3-icon-standard, .bp3-button .bp3-icon-large{
		color:#5c7080; }
		.bp3-button .bp3-icon.bp3-align-right, .bp3-button .bp3-icon-standard.bp3-align-right, .bp3-button .bp3-icon-large.bp3-align-right{
		  margin-left:7px; }
	  .bp3-button .bp3-icon:first-child:last-child,
	  .bp3-button .bp3-spinner + .bp3-icon:last-child{
		margin:0 -7px; }
	  .bp3-dark .bp3-button:not([class*="bp3-intent-"]){
		background-color:#394b59;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.05)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0));
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]):hover, .bp3-dark .bp3-button:not([class*="bp3-intent-"]):active, .bp3-dark .bp3-button:not([class*="bp3-intent-"]).bp3-active{
		  color:#f5f8fa; }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]):hover{
		  background-color:#30404d;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]):active, .bp3-dark .bp3-button:not([class*="bp3-intent-"]).bp3-active{
		  background-color:#202b33;
		  background-image:none;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]):disabled, .bp3-dark .bp3-button:not([class*="bp3-intent-"]).bp3-disabled{
		  background-color:rgba(57, 75, 89, 0.5);
		  background-image:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(167, 182, 194, 0.6); }
		  .bp3-dark .bp3-button:not([class*="bp3-intent-"]):disabled.bp3-active, .bp3-dark .bp3-button:not([class*="bp3-intent-"]).bp3-disabled.bp3-active{
			background:rgba(57, 75, 89, 0.7); }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]) .bp3-button-spinner .bp3-spinner-head{
		  background:rgba(16, 22, 26, 0.5);
		  stroke:#8a9ba8; }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"])[class*="bp3-icon-"]::before{
		  color:#a7b6c2; }
		.bp3-dark .bp3-button:not([class*="bp3-intent-"]) .bp3-icon, .bp3-dark .bp3-button:not([class*="bp3-intent-"]) .bp3-icon-standard, .bp3-dark .bp3-button:not([class*="bp3-intent-"]) .bp3-icon-large{
		  color:#a7b6c2; }
	  .bp3-dark .bp3-button[class*="bp3-intent-"]{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-button[class*="bp3-intent-"]:hover{
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-button[class*="bp3-intent-"]:active, .bp3-dark .bp3-button[class*="bp3-intent-"].bp3-active{
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-dark .bp3-button[class*="bp3-intent-"]:disabled, .bp3-dark .bp3-button[class*="bp3-intent-"].bp3-disabled{
		  background-image:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(255, 255, 255, 0.3); }
		.bp3-dark .bp3-button[class*="bp3-intent-"] .bp3-button-spinner .bp3-spinner-head{
		  stroke:#8a9ba8; }
	  .bp3-button:disabled::before,
	  .bp3-button:disabled .bp3-icon, .bp3-button:disabled .bp3-icon-standard, .bp3-button:disabled .bp3-icon-large, .bp3-button.bp3-disabled::before,
	  .bp3-button.bp3-disabled .bp3-icon, .bp3-button.bp3-disabled .bp3-icon-standard, .bp3-button.bp3-disabled .bp3-icon-large, .bp3-button[class*="bp3-intent-"]::before,
	  .bp3-button[class*="bp3-intent-"] .bp3-icon, .bp3-button[class*="bp3-intent-"] .bp3-icon-standard, .bp3-button[class*="bp3-intent-"] .bp3-icon-large{
		color:inherit !important; }
	  .bp3-button.bp3-minimal{
		background:none;
		-webkit-box-shadow:none;
				box-shadow:none; }
		.bp3-button.bp3-minimal:hover{
		  background:rgba(167, 182, 194, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026;
		  text-decoration:none; }
		.bp3-button.bp3-minimal:active, .bp3-button.bp3-minimal.bp3-active{
		  background:rgba(115, 134, 148, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026; }
		.bp3-button.bp3-minimal:disabled, .bp3-button.bp3-minimal:disabled:hover, .bp3-button.bp3-minimal.bp3-disabled, .bp3-button.bp3-minimal.bp3-disabled:hover{
		  background:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed; }
		  .bp3-button.bp3-minimal:disabled.bp3-active, .bp3-button.bp3-minimal:disabled:hover.bp3-active, .bp3-button.bp3-minimal.bp3-disabled.bp3-active, .bp3-button.bp3-minimal.bp3-disabled:hover.bp3-active{
			background:rgba(115, 134, 148, 0.3); }
		.bp3-dark .bp3-button.bp3-minimal{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:inherit; }
		  .bp3-dark .bp3-button.bp3-minimal:hover, .bp3-dark .bp3-button.bp3-minimal:active, .bp3-dark .bp3-button.bp3-minimal.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none; }
		  .bp3-dark .bp3-button.bp3-minimal:hover{
			background:rgba(138, 155, 168, 0.15); }
		  .bp3-dark .bp3-button.bp3-minimal:active, .bp3-dark .bp3-button.bp3-minimal.bp3-active{
			background:rgba(138, 155, 168, 0.3);
			color:#f5f8fa; }
		  .bp3-dark .bp3-button.bp3-minimal:disabled, .bp3-dark .bp3-button.bp3-minimal:disabled:hover, .bp3-dark .bp3-button.bp3-minimal.bp3-disabled, .bp3-dark .bp3-button.bp3-minimal.bp3-disabled:hover{
			background:none;
			color:rgba(167, 182, 194, 0.6);
			cursor:not-allowed; }
			.bp3-dark .bp3-button.bp3-minimal:disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal:disabled:hover.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-disabled:hover.bp3-active{
			  background:rgba(138, 155, 168, 0.3); }
		.bp3-button.bp3-minimal.bp3-intent-primary{
		  color:#106ba3; }
		  .bp3-button.bp3-minimal.bp3-intent-primary:hover, .bp3-button.bp3-minimal.bp3-intent-primary:active, .bp3-button.bp3-minimal.bp3-intent-primary.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#106ba3; }
		  .bp3-button.bp3-minimal.bp3-intent-primary:hover{
			background:rgba(19, 124, 189, 0.15);
			color:#106ba3; }
		  .bp3-button.bp3-minimal.bp3-intent-primary:active, .bp3-button.bp3-minimal.bp3-intent-primary.bp3-active{
			background:rgba(19, 124, 189, 0.3);
			color:#106ba3; }
		  .bp3-button.bp3-minimal.bp3-intent-primary:disabled, .bp3-button.bp3-minimal.bp3-intent-primary.bp3-disabled{
			background:none;
			color:rgba(16, 107, 163, 0.5); }
			.bp3-button.bp3-minimal.bp3-intent-primary:disabled.bp3-active, .bp3-button.bp3-minimal.bp3-intent-primary.bp3-disabled.bp3-active{
			  background:rgba(19, 124, 189, 0.3); }
		  .bp3-button.bp3-minimal.bp3-intent-primary .bp3-button-spinner .bp3-spinner-head{
			stroke:#106ba3; }
		  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary{
			color:#48aff0; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary:hover{
			  background:rgba(19, 124, 189, 0.2);
			  color:#48aff0; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary:active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary.bp3-active{
			  background:rgba(19, 124, 189, 0.3);
			  color:#48aff0; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary:disabled, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary.bp3-disabled{
			  background:none;
			  color:rgba(72, 175, 240, 0.5); }
			  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary:disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-primary.bp3-disabled.bp3-active{
				background:rgba(19, 124, 189, 0.3); }
		.bp3-button.bp3-minimal.bp3-intent-success{
		  color:#0d8050; }
		  .bp3-button.bp3-minimal.bp3-intent-success:hover, .bp3-button.bp3-minimal.bp3-intent-success:active, .bp3-button.bp3-minimal.bp3-intent-success.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#0d8050; }
		  .bp3-button.bp3-minimal.bp3-intent-success:hover{
			background:rgba(15, 153, 96, 0.15);
			color:#0d8050; }
		  .bp3-button.bp3-minimal.bp3-intent-success:active, .bp3-button.bp3-minimal.bp3-intent-success.bp3-active{
			background:rgba(15, 153, 96, 0.3);
			color:#0d8050; }
		  .bp3-button.bp3-minimal.bp3-intent-success:disabled, .bp3-button.bp3-minimal.bp3-intent-success.bp3-disabled{
			background:none;
			color:rgba(13, 128, 80, 0.5); }
			.bp3-button.bp3-minimal.bp3-intent-success:disabled.bp3-active, .bp3-button.bp3-minimal.bp3-intent-success.bp3-disabled.bp3-active{
			  background:rgba(15, 153, 96, 0.3); }
		  .bp3-button.bp3-minimal.bp3-intent-success .bp3-button-spinner .bp3-spinner-head{
			stroke:#0d8050; }
		  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-success{
			color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-success:hover{
			  background:rgba(15, 153, 96, 0.2);
			  color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-success:active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-success.bp3-active{
			  background:rgba(15, 153, 96, 0.3);
			  color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-success:disabled, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-success.bp3-disabled{
			  background:none;
			  color:rgba(61, 204, 145, 0.5); }
			  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-success:disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-success.bp3-disabled.bp3-active{
				background:rgba(15, 153, 96, 0.3); }
		.bp3-button.bp3-minimal.bp3-intent-warning{
		  color:#bf7326; }
		  .bp3-button.bp3-minimal.bp3-intent-warning:hover, .bp3-button.bp3-minimal.bp3-intent-warning:active, .bp3-button.bp3-minimal.bp3-intent-warning.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#bf7326; }
		  .bp3-button.bp3-minimal.bp3-intent-warning:hover{
			background:rgba(217, 130, 43, 0.15);
			color:#bf7326; }
		  .bp3-button.bp3-minimal.bp3-intent-warning:active, .bp3-button.bp3-minimal.bp3-intent-warning.bp3-active{
			background:rgba(217, 130, 43, 0.3);
			color:#bf7326; }
		  .bp3-button.bp3-minimal.bp3-intent-warning:disabled, .bp3-button.bp3-minimal.bp3-intent-warning.bp3-disabled{
			background:none;
			color:rgba(191, 115, 38, 0.5); }
			.bp3-button.bp3-minimal.bp3-intent-warning:disabled.bp3-active, .bp3-button.bp3-minimal.bp3-intent-warning.bp3-disabled.bp3-active{
			  background:rgba(217, 130, 43, 0.3); }
		  .bp3-button.bp3-minimal.bp3-intent-warning .bp3-button-spinner .bp3-spinner-head{
			stroke:#bf7326; }
		  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning{
			color:#ffb366; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning:hover{
			  background:rgba(217, 130, 43, 0.2);
			  color:#ffb366; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning:active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning.bp3-active{
			  background:rgba(217, 130, 43, 0.3);
			  color:#ffb366; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning:disabled, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning.bp3-disabled{
			  background:none;
			  color:rgba(255, 179, 102, 0.5); }
			  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning:disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-warning.bp3-disabled.bp3-active{
				background:rgba(217, 130, 43, 0.3); }
		.bp3-button.bp3-minimal.bp3-intent-danger{
		  color:#c23030; }
		  .bp3-button.bp3-minimal.bp3-intent-danger:hover, .bp3-button.bp3-minimal.bp3-intent-danger:active, .bp3-button.bp3-minimal.bp3-intent-danger.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#c23030; }
		  .bp3-button.bp3-minimal.bp3-intent-danger:hover{
			background:rgba(219, 55, 55, 0.15);
			color:#c23030; }
		  .bp3-button.bp3-minimal.bp3-intent-danger:active, .bp3-button.bp3-minimal.bp3-intent-danger.bp3-active{
			background:rgba(219, 55, 55, 0.3);
			color:#c23030; }
		  .bp3-button.bp3-minimal.bp3-intent-danger:disabled, .bp3-button.bp3-minimal.bp3-intent-danger.bp3-disabled{
			background:none;
			color:rgba(194, 48, 48, 0.5); }
			.bp3-button.bp3-minimal.bp3-intent-danger:disabled.bp3-active, .bp3-button.bp3-minimal.bp3-intent-danger.bp3-disabled.bp3-active{
			  background:rgba(219, 55, 55, 0.3); }
		  .bp3-button.bp3-minimal.bp3-intent-danger .bp3-button-spinner .bp3-spinner-head{
			stroke:#c23030; }
		  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger{
			color:#ff7373; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger:hover{
			  background:rgba(219, 55, 55, 0.2);
			  color:#ff7373; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger:active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger.bp3-active{
			  background:rgba(219, 55, 55, 0.3);
			  color:#ff7373; }
			.bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger:disabled, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger.bp3-disabled{
			  background:none;
			  color:rgba(255, 115, 115, 0.5); }
			  .bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger:disabled.bp3-active, .bp3-dark .bp3-button.bp3-minimal.bp3-intent-danger.bp3-disabled.bp3-active{
				background:rgba(219, 55, 55, 0.3); }
	  .bp3-button.bp3-outlined{
		background:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		border:1px solid rgba(24, 32, 38, 0.2);
		-webkit-box-sizing:border-box;
				box-sizing:border-box; }
		.bp3-button.bp3-outlined:hover{
		  background:rgba(167, 182, 194, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026;
		  text-decoration:none; }
		.bp3-button.bp3-outlined:active, .bp3-button.bp3-outlined.bp3-active{
		  background:rgba(115, 134, 148, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026; }
		.bp3-button.bp3-outlined:disabled, .bp3-button.bp3-outlined:disabled:hover, .bp3-button.bp3-outlined.bp3-disabled, .bp3-button.bp3-outlined.bp3-disabled:hover{
		  background:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed; }
		  .bp3-button.bp3-outlined:disabled.bp3-active, .bp3-button.bp3-outlined:disabled:hover.bp3-active, .bp3-button.bp3-outlined.bp3-disabled.bp3-active, .bp3-button.bp3-outlined.bp3-disabled:hover.bp3-active{
			background:rgba(115, 134, 148, 0.3); }
		.bp3-dark .bp3-button.bp3-outlined{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:inherit; }
		  .bp3-dark .bp3-button.bp3-outlined:hover, .bp3-dark .bp3-button.bp3-outlined:active, .bp3-dark .bp3-button.bp3-outlined.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none; }
		  .bp3-dark .bp3-button.bp3-outlined:hover{
			background:rgba(138, 155, 168, 0.15); }
		  .bp3-dark .bp3-button.bp3-outlined:active, .bp3-dark .bp3-button.bp3-outlined.bp3-active{
			background:rgba(138, 155, 168, 0.3);
			color:#f5f8fa; }
		  .bp3-dark .bp3-button.bp3-outlined:disabled, .bp3-dark .bp3-button.bp3-outlined:disabled:hover, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled:hover{
			background:none;
			color:rgba(167, 182, 194, 0.6);
			cursor:not-allowed; }
			.bp3-dark .bp3-button.bp3-outlined:disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined:disabled:hover.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled:hover.bp3-active{
			  background:rgba(138, 155, 168, 0.3); }
		.bp3-button.bp3-outlined.bp3-intent-primary{
		  color:#106ba3; }
		  .bp3-button.bp3-outlined.bp3-intent-primary:hover, .bp3-button.bp3-outlined.bp3-intent-primary:active, .bp3-button.bp3-outlined.bp3-intent-primary.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#106ba3; }
		  .bp3-button.bp3-outlined.bp3-intent-primary:hover{
			background:rgba(19, 124, 189, 0.15);
			color:#106ba3; }
		  .bp3-button.bp3-outlined.bp3-intent-primary:active, .bp3-button.bp3-outlined.bp3-intent-primary.bp3-active{
			background:rgba(19, 124, 189, 0.3);
			color:#106ba3; }
		  .bp3-button.bp3-outlined.bp3-intent-primary:disabled, .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled{
			background:none;
			color:rgba(16, 107, 163, 0.5); }
			.bp3-button.bp3-outlined.bp3-intent-primary:disabled.bp3-active, .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled.bp3-active{
			  background:rgba(19, 124, 189, 0.3); }
		  .bp3-button.bp3-outlined.bp3-intent-primary .bp3-button-spinner .bp3-spinner-head{
			stroke:#106ba3; }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary{
			color:#48aff0; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary:hover{
			  background:rgba(19, 124, 189, 0.2);
			  color:#48aff0; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary:active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary.bp3-active{
			  background:rgba(19, 124, 189, 0.3);
			  color:#48aff0; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled{
			  background:none;
			  color:rgba(72, 175, 240, 0.5); }
			  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary:disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled.bp3-active{
				background:rgba(19, 124, 189, 0.3); }
		.bp3-button.bp3-outlined.bp3-intent-success{
		  color:#0d8050; }
		  .bp3-button.bp3-outlined.bp3-intent-success:hover, .bp3-button.bp3-outlined.bp3-intent-success:active, .bp3-button.bp3-outlined.bp3-intent-success.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#0d8050; }
		  .bp3-button.bp3-outlined.bp3-intent-success:hover{
			background:rgba(15, 153, 96, 0.15);
			color:#0d8050; }
		  .bp3-button.bp3-outlined.bp3-intent-success:active, .bp3-button.bp3-outlined.bp3-intent-success.bp3-active{
			background:rgba(15, 153, 96, 0.3);
			color:#0d8050; }
		  .bp3-button.bp3-outlined.bp3-intent-success:disabled, .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled{
			background:none;
			color:rgba(13, 128, 80, 0.5); }
			.bp3-button.bp3-outlined.bp3-intent-success:disabled.bp3-active, .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled.bp3-active{
			  background:rgba(15, 153, 96, 0.3); }
		  .bp3-button.bp3-outlined.bp3-intent-success .bp3-button-spinner .bp3-spinner-head{
			stroke:#0d8050; }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success{
			color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-success:hover{
			  background:rgba(15, 153, 96, 0.2);
			  color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-success:active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success.bp3-active{
			  background:rgba(15, 153, 96, 0.3);
			  color:#3dcc91; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-success:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled{
			  background:none;
			  color:rgba(61, 204, 145, 0.5); }
			  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success:disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled.bp3-active{
				background:rgba(15, 153, 96, 0.3); }
		.bp3-button.bp3-outlined.bp3-intent-warning{
		  color:#bf7326; }
		  .bp3-button.bp3-outlined.bp3-intent-warning:hover, .bp3-button.bp3-outlined.bp3-intent-warning:active, .bp3-button.bp3-outlined.bp3-intent-warning.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#bf7326; }
		  .bp3-button.bp3-outlined.bp3-intent-warning:hover{
			background:rgba(217, 130, 43, 0.15);
			color:#bf7326; }
		  .bp3-button.bp3-outlined.bp3-intent-warning:active, .bp3-button.bp3-outlined.bp3-intent-warning.bp3-active{
			background:rgba(217, 130, 43, 0.3);
			color:#bf7326; }
		  .bp3-button.bp3-outlined.bp3-intent-warning:disabled, .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled{
			background:none;
			color:rgba(191, 115, 38, 0.5); }
			.bp3-button.bp3-outlined.bp3-intent-warning:disabled.bp3-active, .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled.bp3-active{
			  background:rgba(217, 130, 43, 0.3); }
		  .bp3-button.bp3-outlined.bp3-intent-warning .bp3-button-spinner .bp3-spinner-head{
			stroke:#bf7326; }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning{
			color:#ffb366; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning:hover{
			  background:rgba(217, 130, 43, 0.2);
			  color:#ffb366; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning:active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning.bp3-active{
			  background:rgba(217, 130, 43, 0.3);
			  color:#ffb366; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled{
			  background:none;
			  color:rgba(255, 179, 102, 0.5); }
			  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning:disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled.bp3-active{
				background:rgba(217, 130, 43, 0.3); }
		.bp3-button.bp3-outlined.bp3-intent-danger{
		  color:#c23030; }
		  .bp3-button.bp3-outlined.bp3-intent-danger:hover, .bp3-button.bp3-outlined.bp3-intent-danger:active, .bp3-button.bp3-outlined.bp3-intent-danger.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#c23030; }
		  .bp3-button.bp3-outlined.bp3-intent-danger:hover{
			background:rgba(219, 55, 55, 0.15);
			color:#c23030; }
		  .bp3-button.bp3-outlined.bp3-intent-danger:active, .bp3-button.bp3-outlined.bp3-intent-danger.bp3-active{
			background:rgba(219, 55, 55, 0.3);
			color:#c23030; }
		  .bp3-button.bp3-outlined.bp3-intent-danger:disabled, .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled{
			background:none;
			color:rgba(194, 48, 48, 0.5); }
			.bp3-button.bp3-outlined.bp3-intent-danger:disabled.bp3-active, .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled.bp3-active{
			  background:rgba(219, 55, 55, 0.3); }
		  .bp3-button.bp3-outlined.bp3-intent-danger .bp3-button-spinner .bp3-spinner-head{
			stroke:#c23030; }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger{
			color:#ff7373; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger:hover{
			  background:rgba(219, 55, 55, 0.2);
			  color:#ff7373; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger:active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger.bp3-active{
			  background:rgba(219, 55, 55, 0.3);
			  color:#ff7373; }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled{
			  background:none;
			  color:rgba(255, 115, 115, 0.5); }
			  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger:disabled.bp3-active, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled.bp3-active{
				background:rgba(219, 55, 55, 0.3); }
		.bp3-button.bp3-outlined:disabled, .bp3-button.bp3-outlined.bp3-disabled, .bp3-button.bp3-outlined:disabled:hover, .bp3-button.bp3-outlined.bp3-disabled:hover{
		  border-color:rgba(92, 112, 128, 0.1); }
		.bp3-dark .bp3-button.bp3-outlined{
		  border-color:rgba(255, 255, 255, 0.4); }
		  .bp3-dark .bp3-button.bp3-outlined:disabled, .bp3-dark .bp3-button.bp3-outlined:disabled:hover, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-disabled:hover{
			border-color:rgba(255, 255, 255, 0.2); }
		.bp3-button.bp3-outlined.bp3-intent-primary{
		  border-color:rgba(16, 107, 163, 0.6); }
		  .bp3-button.bp3-outlined.bp3-intent-primary:disabled, .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled{
			border-color:rgba(16, 107, 163, 0.2); }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary{
			border-color:rgba(72, 175, 240, 0.6); }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-primary.bp3-disabled{
			  border-color:rgba(72, 175, 240, 0.2); }
		.bp3-button.bp3-outlined.bp3-intent-success{
		  border-color:rgba(13, 128, 80, 0.6); }
		  .bp3-button.bp3-outlined.bp3-intent-success:disabled, .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled{
			border-color:rgba(13, 128, 80, 0.2); }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success{
			border-color:rgba(61, 204, 145, 0.6); }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-success:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-success.bp3-disabled{
			  border-color:rgba(61, 204, 145, 0.2); }
		.bp3-button.bp3-outlined.bp3-intent-warning{
		  border-color:rgba(191, 115, 38, 0.6); }
		  .bp3-button.bp3-outlined.bp3-intent-warning:disabled, .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled{
			border-color:rgba(191, 115, 38, 0.2); }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning{
			border-color:rgba(255, 179, 102, 0.6); }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-warning.bp3-disabled{
			  border-color:rgba(255, 179, 102, 0.2); }
		.bp3-button.bp3-outlined.bp3-intent-danger{
		  border-color:rgba(194, 48, 48, 0.6); }
		  .bp3-button.bp3-outlined.bp3-intent-danger:disabled, .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled{
			border-color:rgba(194, 48, 48, 0.2); }
		  .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger{
			border-color:rgba(255, 115, 115, 0.6); }
			.bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger:disabled, .bp3-dark .bp3-button.bp3-outlined.bp3-intent-danger.bp3-disabled{
			  border-color:rgba(255, 115, 115, 0.2); }
	
	a.bp3-button{
	  text-align:center;
	  text-decoration:none;
	  -webkit-transition:none;
	  transition:none; }
	  a.bp3-button, a.bp3-button:hover, a.bp3-button:active{
		color:#182026; }
	  a.bp3-button.bp3-disabled{
		color:rgba(92, 112, 128, 0.6); }
	
	.bp3-button-text{
	  -webkit-box-flex:0;
		  -ms-flex:0 1 auto;
			  flex:0 1 auto; }
	
	.bp3-button.bp3-align-left .bp3-button-text, .bp3-button.bp3-align-right .bp3-button-text,
	.bp3-button-group.bp3-align-left .bp3-button-text,
	.bp3-button-group.bp3-align-right .bp3-button-text{
	  -webkit-box-flex:1;
		  -ms-flex:1 1 auto;
			  flex:1 1 auto; }
	.bp3-button-group{
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex; }
	  .bp3-button-group .bp3-button{
		-webkit-box-flex:0;
			-ms-flex:0 0 auto;
				flex:0 0 auto;
		position:relative;
		z-index:4; }
		.bp3-button-group .bp3-button:focus{
		  z-index:5; }
		.bp3-button-group .bp3-button:hover{
		  z-index:6; }
		.bp3-button-group .bp3-button:active, .bp3-button-group .bp3-button.bp3-active{
		  z-index:7; }
		.bp3-button-group .bp3-button:disabled, .bp3-button-group .bp3-button.bp3-disabled{
		  z-index:3; }
		.bp3-button-group .bp3-button[class*="bp3-intent-"]{
		  z-index:9; }
		  .bp3-button-group .bp3-button[class*="bp3-intent-"]:focus{
			z-index:10; }
		  .bp3-button-group .bp3-button[class*="bp3-intent-"]:hover{
			z-index:11; }
		  .bp3-button-group .bp3-button[class*="bp3-intent-"]:active, .bp3-button-group .bp3-button[class*="bp3-intent-"].bp3-active{
			z-index:12; }
		  .bp3-button-group .bp3-button[class*="bp3-intent-"]:disabled, .bp3-button-group .bp3-button[class*="bp3-intent-"].bp3-disabled{
			z-index:8; }
	  .bp3-button-group:not(.bp3-minimal) > .bp3-popover-wrapper:not(:first-child) .bp3-button,
	  .bp3-button-group:not(.bp3-minimal) > .bp3-button:not(:first-child){
		border-bottom-left-radius:0;
		border-top-left-radius:0; }
	  .bp3-button-group:not(.bp3-minimal) > .bp3-popover-wrapper:not(:last-child) .bp3-button,
	  .bp3-button-group:not(.bp3-minimal) > .bp3-button:not(:last-child){
		border-bottom-right-radius:0;
		border-top-right-radius:0;
		margin-right:-1px; }
	  .bp3-button-group.bp3-minimal .bp3-button{
		background:none;
		-webkit-box-shadow:none;
				box-shadow:none; }
		.bp3-button-group.bp3-minimal .bp3-button:hover{
		  background:rgba(167, 182, 194, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026;
		  text-decoration:none; }
		.bp3-button-group.bp3-minimal .bp3-button:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-active{
		  background:rgba(115, 134, 148, 0.3);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#182026; }
		.bp3-button-group.bp3-minimal .bp3-button:disabled, .bp3-button-group.bp3-minimal .bp3-button:disabled:hover, .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled, .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled:hover{
		  background:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed; }
		  .bp3-button-group.bp3-minimal .bp3-button:disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button:disabled:hover.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled:hover.bp3-active{
			background:rgba(115, 134, 148, 0.3); }
		.bp3-dark .bp3-button-group.bp3-minimal .bp3-button{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:inherit; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:hover, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:hover{
			background:rgba(138, 155, 168, 0.15); }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-active{
			background:rgba(138, 155, 168, 0.3);
			color:#f5f8fa; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:disabled:hover, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled:hover{
			background:none;
			color:rgba(167, 182, 194, 0.6);
			cursor:not-allowed; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button:disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button:disabled:hover.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-disabled:hover.bp3-active{
			  background:rgba(138, 155, 168, 0.3); }
		.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary{
		  color:#106ba3; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:hover, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#106ba3; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:hover{
			background:rgba(19, 124, 189, 0.15);
			color:#106ba3; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-active{
			background:rgba(19, 124, 189, 0.3);
			color:#106ba3; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:disabled, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-disabled{
			background:none;
			color:rgba(16, 107, 163, 0.5); }
			.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-disabled.bp3-active{
			  background:rgba(19, 124, 189, 0.3); }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary .bp3-button-spinner .bp3-spinner-head{
			stroke:#106ba3; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary{
			color:#48aff0; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:hover{
			  background:rgba(19, 124, 189, 0.2);
			  color:#48aff0; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-active{
			  background:rgba(19, 124, 189, 0.3);
			  color:#48aff0; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-disabled{
			  background:none;
			  color:rgba(72, 175, 240, 0.5); }
			  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary:disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-primary.bp3-disabled.bp3-active{
				background:rgba(19, 124, 189, 0.3); }
		.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success{
		  color:#0d8050; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:hover, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#0d8050; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:hover{
			background:rgba(15, 153, 96, 0.15);
			color:#0d8050; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-active{
			background:rgba(15, 153, 96, 0.3);
			color:#0d8050; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:disabled, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-disabled{
			background:none;
			color:rgba(13, 128, 80, 0.5); }
			.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-disabled.bp3-active{
			  background:rgba(15, 153, 96, 0.3); }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success .bp3-button-spinner .bp3-spinner-head{
			stroke:#0d8050; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success{
			color:#3dcc91; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:hover{
			  background:rgba(15, 153, 96, 0.2);
			  color:#3dcc91; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-active{
			  background:rgba(15, 153, 96, 0.3);
			  color:#3dcc91; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-disabled{
			  background:none;
			  color:rgba(61, 204, 145, 0.5); }
			  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success:disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-success.bp3-disabled.bp3-active{
				background:rgba(15, 153, 96, 0.3); }
		.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning{
		  color:#bf7326; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:hover, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#bf7326; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:hover{
			background:rgba(217, 130, 43, 0.15);
			color:#bf7326; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-active{
			background:rgba(217, 130, 43, 0.3);
			color:#bf7326; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:disabled, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-disabled{
			background:none;
			color:rgba(191, 115, 38, 0.5); }
			.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-disabled.bp3-active{
			  background:rgba(217, 130, 43, 0.3); }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning .bp3-button-spinner .bp3-spinner-head{
			stroke:#bf7326; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning{
			color:#ffb366; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:hover{
			  background:rgba(217, 130, 43, 0.2);
			  color:#ffb366; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-active{
			  background:rgba(217, 130, 43, 0.3);
			  color:#ffb366; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-disabled{
			  background:none;
			  color:rgba(255, 179, 102, 0.5); }
			  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning:disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-warning.bp3-disabled.bp3-active{
				background:rgba(217, 130, 43, 0.3); }
		.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger{
		  color:#c23030; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:hover, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-active{
			background:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:#c23030; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:hover{
			background:rgba(219, 55, 55, 0.15);
			color:#c23030; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-active{
			background:rgba(219, 55, 55, 0.3);
			color:#c23030; }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:disabled, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-disabled{
			background:none;
			color:rgba(194, 48, 48, 0.5); }
			.bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:disabled.bp3-active, .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-disabled.bp3-active{
			  background:rgba(219, 55, 55, 0.3); }
		  .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger .bp3-button-spinner .bp3-spinner-head{
			stroke:#c23030; }
		  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger{
			color:#ff7373; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:hover{
			  background:rgba(219, 55, 55, 0.2);
			  color:#ff7373; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-active{
			  background:rgba(219, 55, 55, 0.3);
			  color:#ff7373; }
			.bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:disabled, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-disabled{
			  background:none;
			  color:rgba(255, 115, 115, 0.5); }
			  .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger:disabled.bp3-active, .bp3-dark .bp3-button-group.bp3-minimal .bp3-button.bp3-intent-danger.bp3-disabled.bp3-active{
				background:rgba(219, 55, 55, 0.3); }
	  .bp3-button-group .bp3-popover-wrapper,
	  .bp3-button-group .bp3-popover-target{
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto; }
	  .bp3-button-group.bp3-fill{
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		width:100%; }
	  .bp3-button-group .bp3-button.bp3-fill,
	  .bp3-button-group.bp3-fill .bp3-button:not(.bp3-fixed){
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto; }
	  .bp3-button-group.bp3-vertical{
		-webkit-box-align:stretch;
			-ms-flex-align:stretch;
				align-items:stretch;
		-webkit-box-orient:vertical;
		-webkit-box-direction:normal;
			-ms-flex-direction:column;
				flex-direction:column;
		vertical-align:top; }
		.bp3-button-group.bp3-vertical.bp3-fill{
		  height:100%;
		  width:unset; }
		.bp3-button-group.bp3-vertical .bp3-button{
		  margin-right:0 !important;
		  width:100%; }
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-popover-wrapper:first-child .bp3-button,
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-button:first-child{
		  border-radius:3px 3px 0 0; }
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-popover-wrapper:last-child .bp3-button,
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-button:last-child{
		  border-radius:0 0 3px 3px; }
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-popover-wrapper:not(:last-child) .bp3-button,
		.bp3-button-group.bp3-vertical:not(.bp3-minimal) > .bp3-button:not(:last-child){
		  margin-bottom:-1px; }
	  .bp3-button-group.bp3-align-left .bp3-button{
		text-align:left; }
	  .bp3-dark .bp3-button-group:not(.bp3-minimal) > .bp3-popover-wrapper:not(:last-child) .bp3-button,
	  .bp3-dark .bp3-button-group:not(.bp3-minimal) > .bp3-button:not(:last-child){
		margin-right:1px; }
	  .bp3-dark .bp3-button-group.bp3-vertical > .bp3-popover-wrapper:not(:last-child) .bp3-button,
	  .bp3-dark .bp3-button-group.bp3-vertical > .bp3-button:not(:last-child){
		margin-bottom:1px; }
	.bp3-callout{
	  font-size:14px;
	  line-height:1.5;
	  background-color:rgba(138, 155, 168, 0.15);
	  border-radius:3px;
	  padding:10px 12px 9px;
	  position:relative;
	  width:100%; }
	  .bp3-callout[class*="bp3-icon-"]{
		padding-left:40px; }
		.bp3-callout[class*="bp3-icon-"]::before{
		  font-family:"Icons20", sans-serif;
		  font-size:20px;
		  font-style:normal;
		  font-weight:400;
		  line-height:1;
		  -moz-osx-font-smoothing:grayscale;
		  -webkit-font-smoothing:antialiased;
		  color:#5c7080;
		  left:10px;
		  position:absolute;
		  top:10px; }
	  .bp3-callout.bp3-callout-icon{
		padding-left:40px; }
		.bp3-callout.bp3-callout-icon > .bp3-icon:first-child{
		  color:#5c7080;
		  left:10px;
		  position:absolute;
		  top:10px; }
	  .bp3-callout .bp3-heading{
		line-height:20px;
		margin-bottom:5px;
		margin-top:0; }
		.bp3-callout .bp3-heading:last-child{
		  margin-bottom:0; }
	  .bp3-dark .bp3-callout{
		background-color:rgba(138, 155, 168, 0.2); }
		.bp3-dark .bp3-callout[class*="bp3-icon-"]::before{
		  color:#a7b6c2; }
	  .bp3-callout.bp3-intent-primary{
		background-color:rgba(19, 124, 189, 0.15); }
		.bp3-callout.bp3-intent-primary[class*="bp3-icon-"]::before,
		.bp3-callout.bp3-intent-primary > .bp3-icon:first-child,
		.bp3-callout.bp3-intent-primary .bp3-heading{
		  color:#106ba3; }
		.bp3-dark .bp3-callout.bp3-intent-primary{
		  background-color:rgba(19, 124, 189, 0.25); }
		  .bp3-dark .bp3-callout.bp3-intent-primary[class*="bp3-icon-"]::before,
		  .bp3-dark .bp3-callout.bp3-intent-primary > .bp3-icon:first-child,
		  .bp3-dark .bp3-callout.bp3-intent-primary .bp3-heading{
			color:#48aff0; }
	  .bp3-callout.bp3-intent-success{
		background-color:rgba(15, 153, 96, 0.15); }
		.bp3-callout.bp3-intent-success[class*="bp3-icon-"]::before,
		.bp3-callout.bp3-intent-success > .bp3-icon:first-child,
		.bp3-callout.bp3-intent-success .bp3-heading{
		  color:#0d8050; }
		.bp3-dark .bp3-callout.bp3-intent-success{
		  background-color:rgba(15, 153, 96, 0.25); }
		  .bp3-dark .bp3-callout.bp3-intent-success[class*="bp3-icon-"]::before,
		  .bp3-dark .bp3-callout.bp3-intent-success > .bp3-icon:first-child,
		  .bp3-dark .bp3-callout.bp3-intent-success .bp3-heading{
			color:#3dcc91; }
	  .bp3-callout.bp3-intent-warning{
		background-color:rgba(217, 130, 43, 0.15); }
		.bp3-callout.bp3-intent-warning[class*="bp3-icon-"]::before,
		.bp3-callout.bp3-intent-warning > .bp3-icon:first-child,
		.bp3-callout.bp3-intent-warning .bp3-heading{
		  color:#bf7326; }
		.bp3-dark .bp3-callout.bp3-intent-warning{
		  background-color:rgba(217, 130, 43, 0.25); }
		  .bp3-dark .bp3-callout.bp3-intent-warning[class*="bp3-icon-"]::before,
		  .bp3-dark .bp3-callout.bp3-intent-warning > .bp3-icon:first-child,
		  .bp3-dark .bp3-callout.bp3-intent-warning .bp3-heading{
			color:#ffb366; }
	  .bp3-callout.bp3-intent-danger{
		background-color:rgba(219, 55, 55, 0.15); }
		.bp3-callout.bp3-intent-danger[class*="bp3-icon-"]::before,
		.bp3-callout.bp3-intent-danger > .bp3-icon:first-child,
		.bp3-callout.bp3-intent-danger .bp3-heading{
		  color:#c23030; }
		.bp3-dark .bp3-callout.bp3-intent-danger{
		  background-color:rgba(219, 55, 55, 0.25); }
		  .bp3-dark .bp3-callout.bp3-intent-danger[class*="bp3-icon-"]::before,
		  .bp3-dark .bp3-callout.bp3-intent-danger > .bp3-icon:first-child,
		  .bp3-dark .bp3-callout.bp3-intent-danger .bp3-heading{
			color:#ff7373; }
	  .bp3-running-text .bp3-callout{
		margin:20px 0; }
	.bp3-card{
	  background-color:#ffffff;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.15), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.15), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0);
	  padding:20px;
	  -webkit-transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), box-shadow 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), box-shadow 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 200ms cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-card.bp3-dark,
	  .bp3-dark .bp3-card{
		background-color:#30404d;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0); }
	
	.bp3-elevation-0{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.15), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.15), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0); }
	  .bp3-elevation-0.bp3-dark,
	  .bp3-dark .bp3-elevation-0{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), 0 0 0 rgba(16, 22, 26, 0), 0 0 0 rgba(16, 22, 26, 0); }
	
	.bp3-elevation-1{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-elevation-1.bp3-dark,
	  .bp3-dark .bp3-elevation-1{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4); }
	
	.bp3-elevation-2{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 1px 1px rgba(16, 22, 26, 0.2), 0 2px 6px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 1px 1px rgba(16, 22, 26, 0.2), 0 2px 6px rgba(16, 22, 26, 0.2); }
	  .bp3-elevation-2.bp3-dark,
	  .bp3-dark .bp3-elevation-2{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.4), 0 2px 6px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.4), 0 2px 6px rgba(16, 22, 26, 0.4); }
	
	.bp3-elevation-3{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2); }
	  .bp3-elevation-3.bp3-dark,
	  .bp3-dark .bp3-elevation-3{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
	
	.bp3-elevation-4{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2); }
	  .bp3-elevation-4.bp3-dark,
	  .bp3-dark .bp3-elevation-4{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4); }
	
	.bp3-card.bp3-interactive:hover{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
	  cursor:pointer; }
	  .bp3-card.bp3-interactive:hover.bp3-dark,
	  .bp3-dark .bp3-card.bp3-interactive:hover{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
	
	.bp3-card.bp3-interactive:active{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
	  opacity:0.9;
	  -webkit-transition-duration:0;
			  transition-duration:0; }
	  .bp3-card.bp3-interactive:active.bp3-dark,
	  .bp3-dark .bp3-card.bp3-interactive:active{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4); }
	
	.bp3-collapse{
	  height:0;
	  overflow-y:hidden;
	  -webkit-transition:height 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:height 200ms cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-collapse .bp3-collapse-body{
		-webkit-transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-collapse .bp3-collapse-body[aria-hidden="true"]{
		  display:none; }
	
	.bp3-context-menu .bp3-popover-target{
	  display:block; }
	
	.bp3-context-menu-popover-target{
	  position:fixed; }
	
	.bp3-divider{
	  border-bottom:1px solid rgba(16, 22, 26, 0.15);
	  border-right:1px solid rgba(16, 22, 26, 0.15);
	  margin:5px; }
	  .bp3-dark .bp3-divider{
		border-color:rgba(16, 22, 26, 0.4); }
	.bp3-dialog-container{
	  opacity:1;
	  -webkit-transform:scale(1);
			  transform:scale(1);
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  min-height:100%;
	  pointer-events:none;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none;
	  width:100%; }
	  .bp3-dialog-container.bp3-overlay-enter > .bp3-dialog, .bp3-dialog-container.bp3-overlay-appear > .bp3-dialog{
		opacity:0;
		-webkit-transform:scale(0.5);
				transform:scale(0.5); }
	  .bp3-dialog-container.bp3-overlay-enter-active > .bp3-dialog, .bp3-dialog-container.bp3-overlay-appear-active > .bp3-dialog{
		opacity:1;
		-webkit-transform:scale(1);
				transform:scale(1);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:opacity, -webkit-transform;
		transition-property:opacity, -webkit-transform;
		transition-property:opacity, transform;
		transition-property:opacity, transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	  .bp3-dialog-container.bp3-overlay-exit > .bp3-dialog{
		opacity:1;
		-webkit-transform:scale(1);
				transform:scale(1); }
	  .bp3-dialog-container.bp3-overlay-exit-active > .bp3-dialog{
		opacity:0;
		-webkit-transform:scale(0.5);
				transform:scale(0.5);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:opacity, -webkit-transform;
		transition-property:opacity, -webkit-transform;
		transition-property:opacity, transform;
		transition-property:opacity, transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	
	.bp3-dialog{
	  background:#ebf1f5;
	  border-radius:6px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  margin:30px 0;
	  padding-bottom:20px;
	  pointer-events:all;
	  -webkit-user-select:text;
		 -moz-user-select:text;
		  -ms-user-select:text;
			  user-select:text;
	  width:500px; }
	  .bp3-dialog:focus{
		outline:0; }
	  .bp3-dialog.bp3-dark,
	  .bp3-dark .bp3-dialog{
		background:#293742;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
	
	.bp3-dialog-header{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  background:#ffffff;
	  border-radius:6px 6px 0 0;
	  -webkit-box-shadow:0 1px 0 rgba(16, 22, 26, 0.15);
			  box-shadow:0 1px 0 rgba(16, 22, 26, 0.15);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  min-height:40px;
	  padding-left:20px;
	  padding-right:5px; }
	  .bp3-dialog-header .bp3-icon-large,
	  .bp3-dialog-header .bp3-icon{
		color:#5c7080;
		-webkit-box-flex:0;
			-ms-flex:0 0 auto;
				flex:0 0 auto;
		margin-right:10px; }
	  .bp3-dialog-header .bp3-heading{
		overflow:hidden;
		text-overflow:ellipsis;
		white-space:nowrap;
		word-wrap:normal;
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto;
		line-height:inherit;
		margin:0; }
		.bp3-dialog-header .bp3-heading:last-child{
		  margin-right:20px; }
	  .bp3-dark .bp3-dialog-header{
		background:#30404d;
		-webkit-box-shadow:0 1px 0 rgba(16, 22, 26, 0.4);
				box-shadow:0 1px 0 rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-dialog-header .bp3-icon-large,
		.bp3-dark .bp3-dialog-header .bp3-icon{
		  color:#a7b6c2; }
	
	.bp3-dialog-body{
	  -webkit-box-flex:1;
		  -ms-flex:1 1 auto;
			  flex:1 1 auto;
	  line-height:18px;
	  margin:20px; }
	
	.bp3-dialog-footer{
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  margin:0 20px; }
	
	.bp3-dialog-footer-actions{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-pack:end;
		  -ms-flex-pack:end;
			  justify-content:flex-end; }
	  .bp3-dialog-footer-actions .bp3-button{
		margin-left:10px; }
	.bp3-drawer{
	  background:#ffffff;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  margin:0;
	  padding:0; }
	  .bp3-drawer:focus{
		outline:0; }
	  .bp3-drawer.bp3-position-top{
		height:50%;
		left:0;
		right:0;
		top:0; }
		.bp3-drawer.bp3-position-top.bp3-overlay-enter, .bp3-drawer.bp3-position-top.bp3-overlay-appear{
		  -webkit-transform:translateY(-100%);
				  transform:translateY(-100%); }
		.bp3-drawer.bp3-position-top.bp3-overlay-enter-active, .bp3-drawer.bp3-position-top.bp3-overlay-appear-active{
		  -webkit-transform:translateY(0);
				  transform:translateY(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer.bp3-position-top.bp3-overlay-exit{
		  -webkit-transform:translateY(0);
				  transform:translateY(0); }
		.bp3-drawer.bp3-position-top.bp3-overlay-exit-active{
		  -webkit-transform:translateY(-100%);
				  transform:translateY(-100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer.bp3-position-bottom{
		bottom:0;
		height:50%;
		left:0;
		right:0; }
		.bp3-drawer.bp3-position-bottom.bp3-overlay-enter, .bp3-drawer.bp3-position-bottom.bp3-overlay-appear{
		  -webkit-transform:translateY(100%);
				  transform:translateY(100%); }
		.bp3-drawer.bp3-position-bottom.bp3-overlay-enter-active, .bp3-drawer.bp3-position-bottom.bp3-overlay-appear-active{
		  -webkit-transform:translateY(0);
				  transform:translateY(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer.bp3-position-bottom.bp3-overlay-exit{
		  -webkit-transform:translateY(0);
				  transform:translateY(0); }
		.bp3-drawer.bp3-position-bottom.bp3-overlay-exit-active{
		  -webkit-transform:translateY(100%);
				  transform:translateY(100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer.bp3-position-left{
		bottom:0;
		left:0;
		top:0;
		width:50%; }
		.bp3-drawer.bp3-position-left.bp3-overlay-enter, .bp3-drawer.bp3-position-left.bp3-overlay-appear{
		  -webkit-transform:translateX(-100%);
				  transform:translateX(-100%); }
		.bp3-drawer.bp3-position-left.bp3-overlay-enter-active, .bp3-drawer.bp3-position-left.bp3-overlay-appear-active{
		  -webkit-transform:translateX(0);
				  transform:translateX(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer.bp3-position-left.bp3-overlay-exit{
		  -webkit-transform:translateX(0);
				  transform:translateX(0); }
		.bp3-drawer.bp3-position-left.bp3-overlay-exit-active{
		  -webkit-transform:translateX(-100%);
				  transform:translateX(-100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer.bp3-position-right{
		bottom:0;
		right:0;
		top:0;
		width:50%; }
		.bp3-drawer.bp3-position-right.bp3-overlay-enter, .bp3-drawer.bp3-position-right.bp3-overlay-appear{
		  -webkit-transform:translateX(100%);
				  transform:translateX(100%); }
		.bp3-drawer.bp3-position-right.bp3-overlay-enter-active, .bp3-drawer.bp3-position-right.bp3-overlay-appear-active{
		  -webkit-transform:translateX(0);
				  transform:translateX(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer.bp3-position-right.bp3-overlay-exit{
		  -webkit-transform:translateX(0);
				  transform:translateX(0); }
		.bp3-drawer.bp3-position-right.bp3-overlay-exit-active{
		  -webkit-transform:translateX(100%);
				  transform:translateX(100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
	  .bp3-position-right):not(.bp3-vertical){
		bottom:0;
		right:0;
		top:0;
		width:50%; }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-enter, .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-appear{
		  -webkit-transform:translateX(100%);
				  transform:translateX(100%); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-enter-active, .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-appear-active{
		  -webkit-transform:translateX(0);
				  transform:translateX(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-exit{
		  -webkit-transform:translateX(0);
				  transform:translateX(0); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right):not(.bp3-vertical).bp3-overlay-exit-active{
		  -webkit-transform:translateX(100%);
				  transform:translateX(100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
	  .bp3-position-right).bp3-vertical{
		bottom:0;
		height:50%;
		left:0;
		right:0; }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-enter, .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-appear{
		  -webkit-transform:translateY(100%);
				  transform:translateY(100%); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-enter-active, .bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-appear-active{
		  -webkit-transform:translateY(0);
				  transform:translateY(0);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:200ms;
				  transition-duration:200ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-exit{
		  -webkit-transform:translateY(0);
				  transform:translateY(0); }
		.bp3-drawer:not(.bp3-position-top):not(.bp3-position-bottom):not(.bp3-position-left):not(
		.bp3-position-right).bp3-vertical.bp3-overlay-exit-active{
		  -webkit-transform:translateY(100%);
				  transform:translateY(100%);
		  -webkit-transition-delay:0;
				  transition-delay:0;
		  -webkit-transition-duration:100ms;
				  transition-duration:100ms;
		  -webkit-transition-property:-webkit-transform;
		  transition-property:-webkit-transform;
		  transition-property:transform;
		  transition-property:transform, -webkit-transform;
		  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-drawer.bp3-dark,
	  .bp3-dark .bp3-drawer{
		background:#30404d;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
	
	.bp3-drawer-header{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  border-radius:0;
	  -webkit-box-shadow:0 1px 0 rgba(16, 22, 26, 0.15);
			  box-shadow:0 1px 0 rgba(16, 22, 26, 0.15);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  min-height:40px;
	  padding:5px;
	  padding-left:20px;
	  position:relative; }
	  .bp3-drawer-header .bp3-icon-large,
	  .bp3-drawer-header .bp3-icon{
		color:#5c7080;
		-webkit-box-flex:0;
			-ms-flex:0 0 auto;
				flex:0 0 auto;
		margin-right:10px; }
	  .bp3-drawer-header .bp3-heading{
		overflow:hidden;
		text-overflow:ellipsis;
		white-space:nowrap;
		word-wrap:normal;
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto;
		line-height:inherit;
		margin:0; }
		.bp3-drawer-header .bp3-heading:last-child{
		  margin-right:20px; }
	  .bp3-dark .bp3-drawer-header{
		-webkit-box-shadow:0 1px 0 rgba(16, 22, 26, 0.4);
				box-shadow:0 1px 0 rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-drawer-header .bp3-icon-large,
		.bp3-dark .bp3-drawer-header .bp3-icon{
		  color:#a7b6c2; }
	
	.bp3-drawer-body{
	  -webkit-box-flex:1;
		  -ms-flex:1 1 auto;
			  flex:1 1 auto;
	  line-height:18px;
	  overflow:auto; }
	
	.bp3-drawer-footer{
	  -webkit-box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.15);
			  box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.15);
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  padding:10px 20px;
	  position:relative; }
	  .bp3-dark .bp3-drawer-footer{
		-webkit-box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.4);
				box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.4); }
	.bp3-editable-text{
	  cursor:text;
	  display:inline-block;
	  max-width:100%;
	  position:relative;
	  vertical-align:top;
	  white-space:nowrap; }
	  .bp3-editable-text::before{
		bottom:-3px;
		left:-3px;
		position:absolute;
		right:-3px;
		top:-3px;
		border-radius:3px;
		content:"";
		-webkit-transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9), box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9), box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-editable-text:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15); }
	  .bp3-editable-text.bp3-editable-text-editing::before{
		background-color:#ffffff;
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-editable-text.bp3-disabled::before{
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-editable-text.bp3-intent-primary .bp3-editable-text-input,
	  .bp3-editable-text.bp3-intent-primary .bp3-editable-text-content{
		color:#137cbd; }
	  .bp3-editable-text.bp3-intent-primary:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(19, 124, 189, 0.4);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(19, 124, 189, 0.4); }
	  .bp3-editable-text.bp3-intent-primary.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-editable-text.bp3-intent-success .bp3-editable-text-input,
	  .bp3-editable-text.bp3-intent-success .bp3-editable-text-content{
		color:#0f9960; }
	  .bp3-editable-text.bp3-intent-success:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px rgba(15, 153, 96, 0.4);
				box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px rgba(15, 153, 96, 0.4); }
	  .bp3-editable-text.bp3-intent-success.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-editable-text.bp3-intent-warning .bp3-editable-text-input,
	  .bp3-editable-text.bp3-intent-warning .bp3-editable-text-content{
		color:#d9822b; }
	  .bp3-editable-text.bp3-intent-warning:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px rgba(217, 130, 43, 0.4);
				box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px rgba(217, 130, 43, 0.4); }
	  .bp3-editable-text.bp3-intent-warning.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-editable-text.bp3-intent-danger .bp3-editable-text-input,
	  .bp3-editable-text.bp3-intent-danger .bp3-editable-text-content{
		color:#db3737; }
	  .bp3-editable-text.bp3-intent-danger:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px rgba(219, 55, 55, 0.4);
				box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px rgba(219, 55, 55, 0.4); }
	  .bp3-editable-text.bp3-intent-danger.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-editable-text:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(255, 255, 255, 0.15);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(255, 255, 255, 0.15); }
	  .bp3-dark .bp3-editable-text.bp3-editable-text-editing::before{
		background-color:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-disabled::before{
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-dark .bp3-editable-text.bp3-intent-primary .bp3-editable-text-content{
		color:#48aff0; }
	  .bp3-dark .bp3-editable-text.bp3-intent-primary:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(72, 175, 240, 0), 0 0 0 0 rgba(72, 175, 240, 0), inset 0 0 0 1px rgba(72, 175, 240, 0.4);
				box-shadow:0 0 0 0 rgba(72, 175, 240, 0), 0 0 0 0 rgba(72, 175, 240, 0), inset 0 0 0 1px rgba(72, 175, 240, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-primary.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #48aff0, 0 0 0 3px rgba(72, 175, 240, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #48aff0, 0 0 0 3px rgba(72, 175, 240, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-success .bp3-editable-text-content{
		color:#3dcc91; }
	  .bp3-dark .bp3-editable-text.bp3-intent-success:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(61, 204, 145, 0), 0 0 0 0 rgba(61, 204, 145, 0), inset 0 0 0 1px rgba(61, 204, 145, 0.4);
				box-shadow:0 0 0 0 rgba(61, 204, 145, 0), 0 0 0 0 rgba(61, 204, 145, 0), inset 0 0 0 1px rgba(61, 204, 145, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-success.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #3dcc91, 0 0 0 3px rgba(61, 204, 145, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #3dcc91, 0 0 0 3px rgba(61, 204, 145, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-warning .bp3-editable-text-content{
		color:#ffb366; }
	  .bp3-dark .bp3-editable-text.bp3-intent-warning:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(255, 179, 102, 0), 0 0 0 0 rgba(255, 179, 102, 0), inset 0 0 0 1px rgba(255, 179, 102, 0.4);
				box-shadow:0 0 0 0 rgba(255, 179, 102, 0), 0 0 0 0 rgba(255, 179, 102, 0), inset 0 0 0 1px rgba(255, 179, 102, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-warning.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #ffb366, 0 0 0 3px rgba(255, 179, 102, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #ffb366, 0 0 0 3px rgba(255, 179, 102, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-danger .bp3-editable-text-content{
		color:#ff7373; }
	  .bp3-dark .bp3-editable-text.bp3-intent-danger:hover::before{
		-webkit-box-shadow:0 0 0 0 rgba(255, 115, 115, 0), 0 0 0 0 rgba(255, 115, 115, 0), inset 0 0 0 1px rgba(255, 115, 115, 0.4);
				box-shadow:0 0 0 0 rgba(255, 115, 115, 0), 0 0 0 0 rgba(255, 115, 115, 0), inset 0 0 0 1px rgba(255, 115, 115, 0.4); }
	  .bp3-dark .bp3-editable-text.bp3-intent-danger.bp3-editable-text-editing::before{
		-webkit-box-shadow:0 0 0 1px #ff7373, 0 0 0 3px rgba(255, 115, 115, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #ff7373, 0 0 0 3px rgba(255, 115, 115, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	
	.bp3-editable-text-input,
	.bp3-editable-text-content{
	  color:inherit;
	  display:inherit;
	  font:inherit;
	  letter-spacing:inherit;
	  max-width:inherit;
	  min-width:inherit;
	  position:relative;
	  resize:none;
	  text-transform:inherit;
	  vertical-align:top; }
	
	.bp3-editable-text-input{
	  background:none;
	  border:none;
	  -webkit-box-shadow:none;
			  box-shadow:none;
	  padding:0;
	  white-space:pre-wrap;
	  width:100%; }
	  .bp3-editable-text-input::-webkit-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-editable-text-input::-moz-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-editable-text-input:-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-editable-text-input::-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-editable-text-input::placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-editable-text-input:focus{
		outline:none; }
	  .bp3-editable-text-input::-ms-clear{
		display:none; }
	
	.bp3-editable-text-content{
	  overflow:hidden;
	  padding-right:2px;
	  text-overflow:ellipsis;
	  white-space:pre; }
	  .bp3-editable-text-editing > .bp3-editable-text-content{
		left:0;
		position:absolute;
		visibility:hidden; }
	  .bp3-editable-text-placeholder > .bp3-editable-text-content{
		color:rgba(92, 112, 128, 0.6); }
		.bp3-dark .bp3-editable-text-placeholder > .bp3-editable-text-content{
		  color:rgba(167, 182, 194, 0.6); }
	
	.bp3-editable-text.bp3-multiline{
	  display:block; }
	  .bp3-editable-text.bp3-multiline .bp3-editable-text-content{
		overflow:auto;
		white-space:pre-wrap;
		word-wrap:break-word; }
	.bp3-divider{
	  border-bottom:1px solid rgba(16, 22, 26, 0.15);
	  border-right:1px solid rgba(16, 22, 26, 0.15);
	  margin:5px; }
	  .bp3-dark .bp3-divider{
		border-color:rgba(16, 22, 26, 0.4); }
	.bp3-control-group{
	  -webkit-transform:translateZ(0);
			  transform:translateZ(0);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:stretch;
		  -ms-flex-align:stretch;
			  align-items:stretch; }
	  .bp3-control-group > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-control-group > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-control-group .bp3-button,
	  .bp3-control-group .bp3-html-select,
	  .bp3-control-group .bp3-input,
	  .bp3-control-group .bp3-select{
		position:relative; }
	  .bp3-control-group .bp3-input{
		border-radius:inherit;
		z-index:2; }
		.bp3-control-group .bp3-input:focus{
		  border-radius:3px;
		  z-index:14; }
		.bp3-control-group .bp3-input[class*="bp3-intent"]{
		  z-index:13; }
		  .bp3-control-group .bp3-input[class*="bp3-intent"]:focus{
			z-index:15; }
		.bp3-control-group .bp3-input[readonly], .bp3-control-group .bp3-input:disabled, .bp3-control-group .bp3-input.bp3-disabled{
		  z-index:1; }
	  .bp3-control-group .bp3-input-group[class*="bp3-intent"] .bp3-input{
		z-index:13; }
		.bp3-control-group .bp3-input-group[class*="bp3-intent"] .bp3-input:focus{
		  z-index:15; }
	  .bp3-control-group .bp3-button,
	  .bp3-control-group .bp3-html-select select,
	  .bp3-control-group .bp3-select select{
		-webkit-transform:translateZ(0);
				transform:translateZ(0);
		border-radius:inherit;
		z-index:4; }
		.bp3-control-group .bp3-button:focus,
		.bp3-control-group .bp3-html-select select:focus,
		.bp3-control-group .bp3-select select:focus{
		  z-index:5; }
		.bp3-control-group .bp3-button:hover,
		.bp3-control-group .bp3-html-select select:hover,
		.bp3-control-group .bp3-select select:hover{
		  z-index:6; }
		.bp3-control-group .bp3-button:active,
		.bp3-control-group .bp3-html-select select:active,
		.bp3-control-group .bp3-select select:active{
		  z-index:7; }
		.bp3-control-group .bp3-button[readonly], .bp3-control-group .bp3-button:disabled, .bp3-control-group .bp3-button.bp3-disabled,
		.bp3-control-group .bp3-html-select select[readonly],
		.bp3-control-group .bp3-html-select select:disabled,
		.bp3-control-group .bp3-html-select select.bp3-disabled,
		.bp3-control-group .bp3-select select[readonly],
		.bp3-control-group .bp3-select select:disabled,
		.bp3-control-group .bp3-select select.bp3-disabled{
		  z-index:3; }
		.bp3-control-group .bp3-button[class*="bp3-intent"],
		.bp3-control-group .bp3-html-select select[class*="bp3-intent"],
		.bp3-control-group .bp3-select select[class*="bp3-intent"]{
		  z-index:9; }
		  .bp3-control-group .bp3-button[class*="bp3-intent"]:focus,
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"]:focus,
		  .bp3-control-group .bp3-select select[class*="bp3-intent"]:focus{
			z-index:10; }
		  .bp3-control-group .bp3-button[class*="bp3-intent"]:hover,
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"]:hover,
		  .bp3-control-group .bp3-select select[class*="bp3-intent"]:hover{
			z-index:11; }
		  .bp3-control-group .bp3-button[class*="bp3-intent"]:active,
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"]:active,
		  .bp3-control-group .bp3-select select[class*="bp3-intent"]:active{
			z-index:12; }
		  .bp3-control-group .bp3-button[class*="bp3-intent"][readonly], .bp3-control-group .bp3-button[class*="bp3-intent"]:disabled, .bp3-control-group .bp3-button[class*="bp3-intent"].bp3-disabled,
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"][readonly],
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"]:disabled,
		  .bp3-control-group .bp3-html-select select[class*="bp3-intent"].bp3-disabled,
		  .bp3-control-group .bp3-select select[class*="bp3-intent"][readonly],
		  .bp3-control-group .bp3-select select[class*="bp3-intent"]:disabled,
		  .bp3-control-group .bp3-select select[class*="bp3-intent"].bp3-disabled{
			z-index:8; }
	  .bp3-control-group .bp3-input-group > .bp3-icon,
	  .bp3-control-group .bp3-input-group > .bp3-button,
	  .bp3-control-group .bp3-input-group > .bp3-input-action{
		z-index:16; }
	  .bp3-control-group .bp3-select::after,
	  .bp3-control-group .bp3-html-select::after,
	  .bp3-control-group .bp3-select > .bp3-icon,
	  .bp3-control-group .bp3-html-select > .bp3-icon{
		z-index:17; }
	  .bp3-control-group .bp3-select:focus-within{
		z-index:5; }
	  .bp3-control-group:not(.bp3-vertical) > *:not(.bp3-divider){
		margin-right:-1px; }
	  .bp3-control-group:not(.bp3-vertical) > .bp3-divider:not(:first-child){
		margin-left:6px; }
	  .bp3-dark .bp3-control-group:not(.bp3-vertical) > *:not(.bp3-divider){
		margin-right:0; }
	  .bp3-dark .bp3-control-group:not(.bp3-vertical) > .bp3-button + .bp3-button{
		margin-left:1px; }
	  .bp3-control-group .bp3-popover-wrapper,
	  .bp3-control-group .bp3-popover-target{
		border-radius:inherit; }
	  .bp3-control-group > :first-child{
		border-radius:3px 0 0 3px; }
	  .bp3-control-group > :last-child{
		border-radius:0 3px 3px 0;
		margin-right:0; }
	  .bp3-control-group > :only-child{
		border-radius:3px;
		margin-right:0; }
	  .bp3-control-group .bp3-input-group .bp3-button{
		border-radius:3px; }
	  .bp3-control-group .bp3-numeric-input:not(:first-child) .bp3-input-group{
		border-bottom-left-radius:0;
		border-top-left-radius:0; }
	  .bp3-control-group.bp3-fill{
		width:100%; }
	  .bp3-control-group > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto; }
	  .bp3-control-group.bp3-fill > *:not(.bp3-fixed){
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto; }
	  .bp3-control-group.bp3-vertical{
		-webkit-box-orient:vertical;
		-webkit-box-direction:normal;
			-ms-flex-direction:column;
				flex-direction:column; }
		.bp3-control-group.bp3-vertical > *{
		  margin-top:-1px; }
		.bp3-control-group.bp3-vertical > :first-child{
		  border-radius:3px 3px 0 0;
		  margin-top:0; }
		.bp3-control-group.bp3-vertical > :last-child{
		  border-radius:0 0 3px 3px; }
	.bp3-control{
	  cursor:pointer;
	  display:block;
	  margin-bottom:10px;
	  position:relative;
	  text-transform:none; }
	  .bp3-control input:checked ~ .bp3-control-indicator{
		background-color:#137cbd;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
	  .bp3-control:hover input:checked ~ .bp3-control-indicator{
		background-color:#106ba3;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
	  .bp3-control input:not(:disabled):active:checked ~ .bp3-control-indicator{
		background:#0e5a8a;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-control input:disabled:checked ~ .bp3-control-indicator{
		background:rgba(19, 124, 189, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-dark .bp3-control input:checked ~ .bp3-control-indicator{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control:hover input:checked ~ .bp3-control-indicator{
		background-color:#106ba3;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control input:not(:disabled):active:checked ~ .bp3-control-indicator{
		background-color:#0e5a8a;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-control input:disabled:checked ~ .bp3-control-indicator{
		background:rgba(14, 90, 138, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-control:not(.bp3-align-right){
		padding-left:26px; }
		.bp3-control:not(.bp3-align-right) .bp3-control-indicator{
		  margin-left:-26px; }
	  .bp3-control.bp3-align-right{
		padding-right:26px; }
		.bp3-control.bp3-align-right .bp3-control-indicator{
		  margin-right:-26px; }
	  .bp3-control.bp3-disabled{
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed; }
	  .bp3-control.bp3-inline{
		display:inline-block;
		margin-right:20px; }
	  .bp3-control input{
		left:0;
		opacity:0;
		position:absolute;
		top:0;
		z-index:-1; }
	  .bp3-control .bp3-control-indicator{
		background-clip:padding-box;
		background-color:#f5f8fa;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.8)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0));
		border:none;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
		cursor:pointer;
		display:inline-block;
		font-size:16px;
		height:1em;
		margin-right:10px;
		margin-top:-3px;
		position:relative;
		-webkit-user-select:none;
		   -moz-user-select:none;
			-ms-user-select:none;
				user-select:none;
		vertical-align:middle;
		width:1em; }
		.bp3-control .bp3-control-indicator::before{
		  content:"";
		  display:block;
		  height:1em;
		  width:1em; }
	  .bp3-control:hover .bp3-control-indicator{
		background-color:#ebf1f5; }
	  .bp3-control input:not(:disabled):active ~ .bp3-control-indicator{
		background:#d8e1e8;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-control input:disabled ~ .bp3-control-indicator{
		background:rgba(206, 217, 224, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none;
		cursor:not-allowed; }
	  .bp3-control input:focus ~ .bp3-control-indicator{
		outline:rgba(19, 124, 189, 0.6) auto 2px;
		outline-offset:2px;
		-moz-outline-radius:6px; }
	  .bp3-control.bp3-align-right .bp3-control-indicator{
		float:right;
		margin-left:10px;
		margin-top:1px; }
	  .bp3-control.bp3-large{
		font-size:16px; }
		.bp3-control.bp3-large:not(.bp3-align-right){
		  padding-left:30px; }
		  .bp3-control.bp3-large:not(.bp3-align-right) .bp3-control-indicator{
			margin-left:-30px; }
		.bp3-control.bp3-large.bp3-align-right{
		  padding-right:30px; }
		  .bp3-control.bp3-large.bp3-align-right .bp3-control-indicator{
			margin-right:-30px; }
		.bp3-control.bp3-large .bp3-control-indicator{
		  font-size:20px; }
		.bp3-control.bp3-large.bp3-align-right .bp3-control-indicator{
		  margin-top:0; }
	  .bp3-control.bp3-checkbox input:indeterminate ~ .bp3-control-indicator{
		background-color:#137cbd;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.1)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
		color:#ffffff; }
	  .bp3-control.bp3-checkbox:hover input:indeterminate ~ .bp3-control-indicator{
		background-color:#106ba3;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 -1px 0 rgba(16, 22, 26, 0.2); }
	  .bp3-control.bp3-checkbox input:not(:disabled):active:indeterminate ~ .bp3-control-indicator{
		background:#0e5a8a;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-control.bp3-checkbox input:disabled:indeterminate ~ .bp3-control-indicator{
		background:rgba(19, 124, 189, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-dark .bp3-control.bp3-checkbox input:indeterminate ~ .bp3-control-indicator{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control.bp3-checkbox:hover input:indeterminate ~ .bp3-control-indicator{
		background-color:#106ba3;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control.bp3-checkbox input:not(:disabled):active:indeterminate ~ .bp3-control-indicator{
		background-color:#0e5a8a;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-control.bp3-checkbox input:disabled:indeterminate ~ .bp3-control-indicator{
		background:rgba(14, 90, 138, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-control.bp3-checkbox .bp3-control-indicator{
		border-radius:3px; }
	  .bp3-control.bp3-checkbox input:checked ~ .bp3-control-indicator::before{
		background-image:url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'%3e%3cpath fill-rule='evenodd' clip-rule='evenodd' d='M12 5c-.28 0-.53.11-.71.29L7 9.59l-2.29-2.3a1.003 1.003 0 00-1.42 1.42l3 3c.18.18.43.29.71.29s.53-.11.71-.29l5-5A1.003 1.003 0 0012 5z' fill='white'/%3e%3c/svg%3e"); }
	  .bp3-control.bp3-checkbox input:indeterminate ~ .bp3-control-indicator::before{
		background-image:url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'%3e%3cpath fill-rule='evenodd' clip-rule='evenodd' d='M11 7H5c-.55 0-1 .45-1 1s.45 1 1 1h6c.55 0 1-.45 1-1s-.45-1-1-1z' fill='white'/%3e%3c/svg%3e"); }
	  .bp3-control.bp3-radio .bp3-control-indicator{
		border-radius:50%; }
	  .bp3-control.bp3-radio input:checked ~ .bp3-control-indicator::before{
		background-image:radial-gradient(#ffffff, #ffffff 28%, transparent 32%); }
	  .bp3-control.bp3-radio input:checked:disabled ~ .bp3-control-indicator::before{
		opacity:0.5; }
	  .bp3-control.bp3-radio input:focus ~ .bp3-control-indicator{
		-moz-outline-radius:16px; }
	  .bp3-control.bp3-switch input ~ .bp3-control-indicator{
		background:rgba(167, 182, 194, 0.5); }
	  .bp3-control.bp3-switch:hover input ~ .bp3-control-indicator{
		background:rgba(115, 134, 148, 0.5); }
	  .bp3-control.bp3-switch input:not(:disabled):active ~ .bp3-control-indicator{
		background:rgba(92, 112, 128, 0.5); }
	  .bp3-control.bp3-switch input:disabled ~ .bp3-control-indicator{
		background:rgba(206, 217, 224, 0.5); }
		.bp3-control.bp3-switch input:disabled ~ .bp3-control-indicator::before{
		  background:rgba(255, 255, 255, 0.8); }
	  .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator{
		background:#137cbd; }
	  .bp3-control.bp3-switch:hover input:checked ~ .bp3-control-indicator{
		background:#106ba3; }
	  .bp3-control.bp3-switch input:checked:not(:disabled):active ~ .bp3-control-indicator{
		background:#0e5a8a; }
	  .bp3-control.bp3-switch input:checked:disabled ~ .bp3-control-indicator{
		background:rgba(19, 124, 189, 0.5); }
		.bp3-control.bp3-switch input:checked:disabled ~ .bp3-control-indicator::before{
		  background:rgba(255, 255, 255, 0.8); }
	  .bp3-control.bp3-switch:not(.bp3-align-right){
		padding-left:38px; }
		.bp3-control.bp3-switch:not(.bp3-align-right) .bp3-control-indicator{
		  margin-left:-38px; }
	  .bp3-control.bp3-switch.bp3-align-right{
		padding-right:38px; }
		.bp3-control.bp3-switch.bp3-align-right .bp3-control-indicator{
		  margin-right:-38px; }
	  .bp3-control.bp3-switch .bp3-control-indicator{
		border:none;
		border-radius:1.75em;
		-webkit-box-shadow:none !important;
				box-shadow:none !important;
		min-width:1.75em;
		-webkit-transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:background-color 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		width:auto; }
		.bp3-control.bp3-switch .bp3-control-indicator::before{
		  background:#ffffff;
		  border-radius:50%;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
		  height:calc(1em - 4px);
		  left:0;
		  margin:2px;
		  position:absolute;
		  -webkit-transition:left 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		  transition:left 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
		  width:calc(1em - 4px); }
	  .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator::before{
		left:calc(100% - 1em); }
	  .bp3-control.bp3-switch.bp3-large:not(.bp3-align-right){
		padding-left:45px; }
		.bp3-control.bp3-switch.bp3-large:not(.bp3-align-right) .bp3-control-indicator{
		  margin-left:-45px; }
	  .bp3-control.bp3-switch.bp3-large.bp3-align-right{
		padding-right:45px; }
		.bp3-control.bp3-switch.bp3-large.bp3-align-right .bp3-control-indicator{
		  margin-right:-45px; }
	  .bp3-dark .bp3-control.bp3-switch input ~ .bp3-control-indicator{
		background:rgba(16, 22, 26, 0.5); }
	  .bp3-dark .bp3-control.bp3-switch:hover input ~ .bp3-control-indicator{
		background:rgba(16, 22, 26, 0.7); }
	  .bp3-dark .bp3-control.bp3-switch input:not(:disabled):active ~ .bp3-control-indicator{
		background:rgba(16, 22, 26, 0.9); }
	  .bp3-dark .bp3-control.bp3-switch input:disabled ~ .bp3-control-indicator{
		background:rgba(57, 75, 89, 0.5); }
		.bp3-dark .bp3-control.bp3-switch input:disabled ~ .bp3-control-indicator::before{
		  background:rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator{
		background:#137cbd; }
	  .bp3-dark .bp3-control.bp3-switch:hover input:checked ~ .bp3-control-indicator{
		background:#106ba3; }
	  .bp3-dark .bp3-control.bp3-switch input:checked:not(:disabled):active ~ .bp3-control-indicator{
		background:#0e5a8a; }
	  .bp3-dark .bp3-control.bp3-switch input:checked:disabled ~ .bp3-control-indicator{
		background:rgba(14, 90, 138, 0.5); }
		.bp3-dark .bp3-control.bp3-switch input:checked:disabled ~ .bp3-control-indicator::before{
		  background:rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control.bp3-switch .bp3-control-indicator::before{
		background:#394b59;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator::before{
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-control.bp3-switch .bp3-switch-inner-text{
		font-size:0.7em;
		text-align:center; }
	  .bp3-control.bp3-switch .bp3-control-indicator-child:first-child{
		line-height:0;
		margin-left:0.5em;
		margin-right:1.2em;
		visibility:hidden; }
	  .bp3-control.bp3-switch .bp3-control-indicator-child:last-child{
		line-height:1em;
		margin-left:1.2em;
		margin-right:0.5em;
		visibility:visible; }
	  .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator .bp3-control-indicator-child:first-child{
		line-height:1em;
		visibility:visible; }
	  .bp3-control.bp3-switch input:checked ~ .bp3-control-indicator .bp3-control-indicator-child:last-child{
		line-height:0;
		visibility:hidden; }
	  .bp3-dark .bp3-control{
		color:#f5f8fa; }
		.bp3-dark .bp3-control.bp3-disabled{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-control .bp3-control-indicator{
		  background-color:#394b59;
		  background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.05)), to(rgba(255, 255, 255, 0)));
		  background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0));
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-control:hover .bp3-control-indicator{
		  background-color:#30404d; }
		.bp3-dark .bp3-control input:not(:disabled):active ~ .bp3-control-indicator{
		  background:#202b33;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-dark .bp3-control input:disabled ~ .bp3-control-indicator{
		  background:rgba(57, 75, 89, 0.5);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  cursor:not-allowed; }
		.bp3-dark .bp3-control.bp3-checkbox input:disabled:checked ~ .bp3-control-indicator, .bp3-dark .bp3-control.bp3-checkbox input:disabled:indeterminate ~ .bp3-control-indicator{
		  color:rgba(167, 182, 194, 0.6); }
	.bp3-file-input{
	  cursor:pointer;
	  display:inline-block;
	  height:30px;
	  position:relative; }
	  .bp3-file-input input{
		margin:0;
		min-width:200px;
		opacity:0; }
		.bp3-file-input input:disabled + .bp3-file-upload-input,
		.bp3-file-input input.bp3-disabled + .bp3-file-upload-input{
		  background:rgba(206, 217, 224, 0.5);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed;
		  resize:none; }
		  .bp3-file-input input:disabled + .bp3-file-upload-input::after,
		  .bp3-file-input input.bp3-disabled + .bp3-file-upload-input::after{
			background-color:rgba(206, 217, 224, 0.5);
			background-image:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:rgba(92, 112, 128, 0.6);
			cursor:not-allowed;
			outline:none; }
			.bp3-file-input input:disabled + .bp3-file-upload-input::after.bp3-active, .bp3-file-input input:disabled + .bp3-file-upload-input::after.bp3-active:hover,
			.bp3-file-input input.bp3-disabled + .bp3-file-upload-input::after.bp3-active,
			.bp3-file-input input.bp3-disabled + .bp3-file-upload-input::after.bp3-active:hover{
			  background:rgba(206, 217, 224, 0.7); }
		  .bp3-dark .bp3-file-input input:disabled + .bp3-file-upload-input, .bp3-dark
		  .bp3-file-input input.bp3-disabled + .bp3-file-upload-input{
			background:rgba(57, 75, 89, 0.5);
			-webkit-box-shadow:none;
					box-shadow:none;
			color:rgba(167, 182, 194, 0.6); }
			.bp3-dark .bp3-file-input input:disabled + .bp3-file-upload-input::after, .bp3-dark
			.bp3-file-input input.bp3-disabled + .bp3-file-upload-input::after{
			  background-color:rgba(57, 75, 89, 0.5);
			  background-image:none;
			  -webkit-box-shadow:none;
					  box-shadow:none;
			  color:rgba(167, 182, 194, 0.6); }
			  .bp3-dark .bp3-file-input input:disabled + .bp3-file-upload-input::after.bp3-active, .bp3-dark
			  .bp3-file-input input.bp3-disabled + .bp3-file-upload-input::after.bp3-active{
				background:rgba(57, 75, 89, 0.7); }
	  .bp3-file-input.bp3-file-input-has-selection .bp3-file-upload-input{
		color:#182026; }
	  .bp3-dark .bp3-file-input.bp3-file-input-has-selection .bp3-file-upload-input{
		color:#f5f8fa; }
	  .bp3-file-input.bp3-fill{
		width:100%; }
	  .bp3-file-input.bp3-large,
	  .bp3-large .bp3-file-input{
		height:40px; }
	  .bp3-file-input .bp3-file-upload-input-custom-text::after{
		content:attr(bp3-button-text); }
	
	.bp3-file-upload-input{
	  -webkit-appearance:none;
		 -moz-appearance:none;
			  appearance:none;
	  background:#ffffff;
	  border:none;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
	  color:#182026;
	  font-size:14px;
	  font-weight:400;
	  height:30px;
	  line-height:30px;
	  outline:none;
	  padding:0 10px;
	  -webkit-transition:-webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:-webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  vertical-align:middle;
	  overflow:hidden;
	  text-overflow:ellipsis;
	  white-space:nowrap;
	  word-wrap:normal;
	  color:rgba(92, 112, 128, 0.6);
	  left:0;
	  padding-right:80px;
	  position:absolute;
	  right:0;
	  top:0;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-file-upload-input::-webkit-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-file-upload-input::-moz-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-file-upload-input:-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-file-upload-input::-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-file-upload-input::placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-file-upload-input:focus, .bp3-file-upload-input.bp3-active{
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-file-upload-input[type="search"], .bp3-file-upload-input.bp3-round{
		border-radius:30px;
		-webkit-box-sizing:border-box;
				box-sizing:border-box;
		padding-left:10px; }
	  .bp3-file-upload-input[readonly]{
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15); }
	  .bp3-file-upload-input:disabled, .bp3-file-upload-input.bp3-disabled{
		background:rgba(206, 217, 224, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed;
		resize:none; }
	  .bp3-file-upload-input::after{
		background-color:#f5f8fa;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.8)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0));
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
		color:#182026;
		min-height:24px;
		min-width:24px;
		overflow:hidden;
		text-overflow:ellipsis;
		white-space:nowrap;
		word-wrap:normal;
		border-radius:3px;
		content:"Browse";
		line-height:24px;
		margin:3px;
		position:absolute;
		right:0;
		text-align:center;
		top:0;
		width:70px; }
		.bp3-file-upload-input::after:hover{
		  background-clip:padding-box;
		  background-color:#ebf1f5;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
		.bp3-file-upload-input::after:active, .bp3-file-upload-input::after.bp3-active{
		  background-color:#d8e1e8;
		  background-image:none;
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-file-upload-input::after:disabled, .bp3-file-upload-input::after.bp3-disabled{
		  background-color:rgba(206, 217, 224, 0.5);
		  background-image:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(92, 112, 128, 0.6);
		  cursor:not-allowed;
		  outline:none; }
		  .bp3-file-upload-input::after:disabled.bp3-active, .bp3-file-upload-input::after:disabled.bp3-active:hover, .bp3-file-upload-input::after.bp3-disabled.bp3-active, .bp3-file-upload-input::after.bp3-disabled.bp3-active:hover{
			background:rgba(206, 217, 224, 0.7); }
	  .bp3-file-upload-input:hover::after{
		background-clip:padding-box;
		background-color:#ebf1f5;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
	  .bp3-file-upload-input:active::after{
		background-color:#d8e1e8;
		background-image:none;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-large .bp3-file-upload-input{
		font-size:16px;
		height:40px;
		line-height:40px;
		padding-right:95px; }
		.bp3-large .bp3-file-upload-input[type="search"], .bp3-large .bp3-file-upload-input.bp3-round{
		  padding:0 15px; }
		.bp3-large .bp3-file-upload-input::after{
		  min-height:30px;
		  min-width:30px;
		  line-height:30px;
		  margin:5px;
		  width:85px; }
	  .bp3-dark .bp3-file-upload-input{
		background:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa;
		color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input::-webkit-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input::-moz-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input:-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input::-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input::placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input:focus{
		  -webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-file-upload-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-file-upload-input:disabled, .bp3-dark .bp3-file-upload-input.bp3-disabled{
		  background:rgba(57, 75, 89, 0.5);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-file-upload-input::after{
		  background-color:#394b59;
		  background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.05)), to(rgba(255, 255, 255, 0)));
		  background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0));
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
		  color:#f5f8fa; }
		  .bp3-dark .bp3-file-upload-input::after:hover, .bp3-dark .bp3-file-upload-input::after:active, .bp3-dark .bp3-file-upload-input::after.bp3-active{
			color:#f5f8fa; }
		  .bp3-dark .bp3-file-upload-input::after:hover{
			background-color:#30404d;
			-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
					box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-file-upload-input::after:active, .bp3-dark .bp3-file-upload-input::after.bp3-active{
			background-color:#202b33;
			background-image:none;
			-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
					box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		  .bp3-dark .bp3-file-upload-input::after:disabled, .bp3-dark .bp3-file-upload-input::after.bp3-disabled{
			background-color:rgba(57, 75, 89, 0.5);
			background-image:none;
			-webkit-box-shadow:none;
					box-shadow:none;
			color:rgba(167, 182, 194, 0.6); }
			.bp3-dark .bp3-file-upload-input::after:disabled.bp3-active, .bp3-dark .bp3-file-upload-input::after.bp3-disabled.bp3-active{
			  background:rgba(57, 75, 89, 0.7); }
		  .bp3-dark .bp3-file-upload-input::after .bp3-button-spinner .bp3-spinner-head{
			background:rgba(16, 22, 26, 0.5);
			stroke:#8a9ba8; }
		.bp3-dark .bp3-file-upload-input:hover::after{
		  background-color:#30404d;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-file-upload-input:active::after{
		  background-color:#202b33;
		  background-image:none;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	.bp3-file-upload-input::after{
	  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
			  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
	.bp3-form-group{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  margin:0 0 15px; }
	  .bp3-form-group label.bp3-label{
		margin-bottom:5px; }
	  .bp3-form-group .bp3-control{
		margin-top:7px; }
	  .bp3-form-group .bp3-form-helper-text{
		color:#5c7080;
		font-size:12px;
		margin-top:5px; }
	  .bp3-form-group.bp3-intent-primary .bp3-form-helper-text{
		color:#106ba3; }
	  .bp3-form-group.bp3-intent-success .bp3-form-helper-text{
		color:#0d8050; }
	  .bp3-form-group.bp3-intent-warning .bp3-form-helper-text{
		color:#bf7326; }
	  .bp3-form-group.bp3-intent-danger .bp3-form-helper-text{
		color:#c23030; }
	  .bp3-form-group.bp3-inline{
		-webkit-box-align:start;
			-ms-flex-align:start;
				align-items:flex-start;
		-webkit-box-orient:horizontal;
		-webkit-box-direction:normal;
			-ms-flex-direction:row;
				flex-direction:row; }
		.bp3-form-group.bp3-inline.bp3-large label.bp3-label{
		  line-height:40px;
		  margin:0 10px 0 0; }
		.bp3-form-group.bp3-inline label.bp3-label{
		  line-height:30px;
		  margin:0 10px 0 0; }
	  .bp3-form-group.bp3-disabled .bp3-label,
	  .bp3-form-group.bp3-disabled .bp3-text-muted,
	  .bp3-form-group.bp3-disabled .bp3-form-helper-text{
		color:rgba(92, 112, 128, 0.6) !important; }
	  .bp3-dark .bp3-form-group.bp3-intent-primary .bp3-form-helper-text{
		color:#48aff0; }
	  .bp3-dark .bp3-form-group.bp3-intent-success .bp3-form-helper-text{
		color:#3dcc91; }
	  .bp3-dark .bp3-form-group.bp3-intent-warning .bp3-form-helper-text{
		color:#ffb366; }
	  .bp3-dark .bp3-form-group.bp3-intent-danger .bp3-form-helper-text{
		color:#ff7373; }
	  .bp3-dark .bp3-form-group .bp3-form-helper-text{
		color:#a7b6c2; }
	  .bp3-dark .bp3-form-group.bp3-disabled .bp3-label,
	  .bp3-dark .bp3-form-group.bp3-disabled .bp3-text-muted,
	  .bp3-dark .bp3-form-group.bp3-disabled .bp3-form-helper-text{
		color:rgba(167, 182, 194, 0.6) !important; }
	.bp3-input-group{
	  display:block;
	  position:relative; }
	  .bp3-input-group .bp3-input{
		position:relative;
		width:100%; }
		.bp3-input-group .bp3-input:not(:first-child){
		  padding-left:30px; }
		.bp3-input-group .bp3-input:not(:last-child){
		  padding-right:30px; }
	  .bp3-input-group .bp3-input-action,
	  .bp3-input-group > .bp3-input-left-container,
	  .bp3-input-group > .bp3-button,
	  .bp3-input-group > .bp3-icon{
		position:absolute;
		top:0; }
		.bp3-input-group .bp3-input-action:first-child,
		.bp3-input-group > .bp3-input-left-container:first-child,
		.bp3-input-group > .bp3-button:first-child,
		.bp3-input-group > .bp3-icon:first-child{
		  left:0; }
		.bp3-input-group .bp3-input-action:last-child,
		.bp3-input-group > .bp3-input-left-container:last-child,
		.bp3-input-group > .bp3-button:last-child,
		.bp3-input-group > .bp3-icon:last-child{
		  right:0; }
	  .bp3-input-group .bp3-button{
		min-height:24px;
		min-width:24px;
		margin:3px;
		padding:0 7px; }
		.bp3-input-group .bp3-button:empty{
		  padding:0; }
	  .bp3-input-group > .bp3-input-left-container,
	  .bp3-input-group > .bp3-icon{
		z-index:1; }
	  .bp3-input-group > .bp3-input-left-container > .bp3-icon,
	  .bp3-input-group > .bp3-icon{
		color:#5c7080; }
		.bp3-input-group > .bp3-input-left-container > .bp3-icon:empty,
		.bp3-input-group > .bp3-icon:empty{
		  font-family:"Icons16", sans-serif;
		  font-size:16px;
		  font-style:normal;
		  font-weight:400;
		  line-height:1;
		  -moz-osx-font-smoothing:grayscale;
		  -webkit-font-smoothing:antialiased; }
	  .bp3-input-group > .bp3-input-left-container > .bp3-icon,
	  .bp3-input-group > .bp3-icon,
	  .bp3-input-group .bp3-input-action > .bp3-spinner{
		margin:7px; }
	  .bp3-input-group .bp3-tag{
		margin:5px; }
	  .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:not(:hover):not(:focus),
	  .bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:not(:hover):not(:focus){
		color:#5c7080; }
		.bp3-dark .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:not(:hover):not(:focus), .bp3-dark
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:not(:hover):not(:focus){
		  color:#a7b6c2; }
		.bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon, .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon-standard, .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon-large,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon-standard,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:not(:hover):not(:focus) .bp3-icon-large{
		  color:#5c7080; }
	  .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:disabled,
	  .bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:disabled{
		color:rgba(92, 112, 128, 0.6) !important; }
		.bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:disabled .bp3-icon, .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:disabled .bp3-icon-standard, .bp3-input-group .bp3-input:not(:focus) + .bp3-button.bp3-minimal:disabled .bp3-icon-large,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:disabled .bp3-icon,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:disabled .bp3-icon-standard,
		.bp3-input-group .bp3-input:not(:focus) + .bp3-input-action .bp3-button.bp3-minimal:disabled .bp3-icon-large{
		  color:rgba(92, 112, 128, 0.6) !important; }
	  .bp3-input-group.bp3-disabled{
		cursor:not-allowed; }
		.bp3-input-group.bp3-disabled .bp3-icon{
		  color:rgba(92, 112, 128, 0.6); }
	  .bp3-input-group.bp3-large .bp3-button{
		min-height:30px;
		min-width:30px;
		margin:5px; }
	  .bp3-input-group.bp3-large > .bp3-input-left-container > .bp3-icon,
	  .bp3-input-group.bp3-large > .bp3-icon,
	  .bp3-input-group.bp3-large .bp3-input-action > .bp3-spinner{
		margin:12px; }
	  .bp3-input-group.bp3-large .bp3-input{
		font-size:16px;
		height:40px;
		line-height:40px; }
		.bp3-input-group.bp3-large .bp3-input[type="search"], .bp3-input-group.bp3-large .bp3-input.bp3-round{
		  padding:0 15px; }
		.bp3-input-group.bp3-large .bp3-input:not(:first-child){
		  padding-left:40px; }
		.bp3-input-group.bp3-large .bp3-input:not(:last-child){
		  padding-right:40px; }
	  .bp3-input-group.bp3-small .bp3-button{
		min-height:20px;
		min-width:20px;
		margin:2px; }
	  .bp3-input-group.bp3-small .bp3-tag{
		min-height:20px;
		min-width:20px;
		margin:2px; }
	  .bp3-input-group.bp3-small > .bp3-input-left-container > .bp3-icon,
	  .bp3-input-group.bp3-small > .bp3-icon,
	  .bp3-input-group.bp3-small .bp3-input-action > .bp3-spinner{
		margin:4px; }
	  .bp3-input-group.bp3-small .bp3-input{
		font-size:12px;
		height:24px;
		line-height:24px;
		padding-left:8px;
		padding-right:8px; }
		.bp3-input-group.bp3-small .bp3-input[type="search"], .bp3-input-group.bp3-small .bp3-input.bp3-round{
		  padding:0 12px; }
		.bp3-input-group.bp3-small .bp3-input:not(:first-child){
		  padding-left:24px; }
		.bp3-input-group.bp3-small .bp3-input:not(:last-child){
		  padding-right:24px; }
	  .bp3-input-group.bp3-fill{
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto;
		width:100%; }
	  .bp3-input-group.bp3-round .bp3-button,
	  .bp3-input-group.bp3-round .bp3-input,
	  .bp3-input-group.bp3-round .bp3-tag{
		border-radius:30px; }
	  .bp3-dark .bp3-input-group .bp3-icon{
		color:#a7b6c2; }
	  .bp3-dark .bp3-input-group.bp3-disabled .bp3-icon{
		color:rgba(167, 182, 194, 0.6); }
	  .bp3-input-group.bp3-intent-primary .bp3-input{
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-primary .bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-primary .bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #137cbd;
				  box-shadow:inset 0 0 0 1px #137cbd; }
		.bp3-input-group.bp3-intent-primary .bp3-input:disabled, .bp3-input-group.bp3-intent-primary .bp3-input.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-input-group.bp3-intent-primary > .bp3-icon{
		color:#106ba3; }
		.bp3-dark .bp3-input-group.bp3-intent-primary > .bp3-icon{
		  color:#48aff0; }
	  .bp3-input-group.bp3-intent-success .bp3-input{
		-webkit-box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-success .bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-success .bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #0f9960;
				  box-shadow:inset 0 0 0 1px #0f9960; }
		.bp3-input-group.bp3-intent-success .bp3-input:disabled, .bp3-input-group.bp3-intent-success .bp3-input.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-input-group.bp3-intent-success > .bp3-icon{
		color:#0d8050; }
		.bp3-dark .bp3-input-group.bp3-intent-success > .bp3-icon{
		  color:#3dcc91; }
	  .bp3-input-group.bp3-intent-warning .bp3-input{
		-webkit-box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-warning .bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-warning .bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #d9822b;
				  box-shadow:inset 0 0 0 1px #d9822b; }
		.bp3-input-group.bp3-intent-warning .bp3-input:disabled, .bp3-input-group.bp3-intent-warning .bp3-input.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-input-group.bp3-intent-warning > .bp3-icon{
		color:#bf7326; }
		.bp3-dark .bp3-input-group.bp3-intent-warning > .bp3-icon{
		  color:#ffb366; }
	  .bp3-input-group.bp3-intent-danger .bp3-input{
		-webkit-box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-danger .bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input-group.bp3-intent-danger .bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #db3737;
				  box-shadow:inset 0 0 0 1px #db3737; }
		.bp3-input-group.bp3-intent-danger .bp3-input:disabled, .bp3-input-group.bp3-intent-danger .bp3-input.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-input-group.bp3-intent-danger > .bp3-icon{
		color:#c23030; }
		.bp3-dark .bp3-input-group.bp3-intent-danger > .bp3-icon{
		  color:#ff7373; }
	.bp3-input{
	  -webkit-appearance:none;
		 -moz-appearance:none;
			  appearance:none;
	  background:#ffffff;
	  border:none;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
	  color:#182026;
	  font-size:14px;
	  font-weight:400;
	  height:30px;
	  line-height:30px;
	  outline:none;
	  padding:0 10px;
	  -webkit-transition:-webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:-webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-box-shadow 100ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  vertical-align:middle; }
	  .bp3-input::-webkit-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input::-moz-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input:-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input::-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input::placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input:focus, .bp3-input.bp3-active{
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-input[type="search"], .bp3-input.bp3-round{
		border-radius:30px;
		-webkit-box-sizing:border-box;
				box-sizing:border-box;
		padding-left:10px; }
	  .bp3-input[readonly]{
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.15); }
	  .bp3-input:disabled, .bp3-input.bp3-disabled{
		background:rgba(206, 217, 224, 0.5);
		-webkit-box-shadow:none;
				box-shadow:none;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed;
		resize:none; }
	  .bp3-input.bp3-large{
		font-size:16px;
		height:40px;
		line-height:40px; }
		.bp3-input.bp3-large[type="search"], .bp3-input.bp3-large.bp3-round{
		  padding:0 15px; }
	  .bp3-input.bp3-small{
		font-size:12px;
		height:24px;
		line-height:24px;
		padding-left:8px;
		padding-right:8px; }
		.bp3-input.bp3-small[type="search"], .bp3-input.bp3-small.bp3-round{
		  padding:0 12px; }
	  .bp3-input.bp3-fill{
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto;
		width:100%; }
	  .bp3-dark .bp3-input{
		background:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
		.bp3-dark .bp3-input::-webkit-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-input::-moz-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-input:-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-input::-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-input::placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-input:disabled, .bp3-dark .bp3-input.bp3-disabled{
		  background:rgba(57, 75, 89, 0.5);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(167, 182, 194, 0.6); }
	  .bp3-input.bp3-intent-primary{
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-primary:focus{
		  -webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-primary[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #137cbd;
				  box-shadow:inset 0 0 0 1px #137cbd; }
		.bp3-input.bp3-intent-primary:disabled, .bp3-input.bp3-intent-primary.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
		.bp3-dark .bp3-input.bp3-intent-primary{
		  -webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px #137cbd, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-primary:focus{
			-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
					box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-primary[readonly]{
			-webkit-box-shadow:inset 0 0 0 1px #137cbd;
					box-shadow:inset 0 0 0 1px #137cbd; }
		  .bp3-dark .bp3-input.bp3-intent-primary:disabled, .bp3-dark .bp3-input.bp3-intent-primary.bp3-disabled{
			-webkit-box-shadow:none;
					box-shadow:none; }
	  .bp3-input.bp3-intent-success{
		-webkit-box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-success:focus{
		  -webkit-box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-success[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #0f9960;
				  box-shadow:inset 0 0 0 1px #0f9960; }
		.bp3-input.bp3-intent-success:disabled, .bp3-input.bp3-intent-success.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
		.bp3-dark .bp3-input.bp3-intent-success{
		  -webkit-box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), 0 0 0 0 rgba(15, 153, 96, 0), inset 0 0 0 1px #0f9960, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-success:focus{
			-webkit-box-shadow:0 0 0 1px #0f9960, 0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
					box-shadow:0 0 0 1px #0f9960, 0 0 0 1px #0f9960, 0 0 0 3px rgba(15, 153, 96, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-success[readonly]{
			-webkit-box-shadow:inset 0 0 0 1px #0f9960;
					box-shadow:inset 0 0 0 1px #0f9960; }
		  .bp3-dark .bp3-input.bp3-intent-success:disabled, .bp3-dark .bp3-input.bp3-intent-success.bp3-disabled{
			-webkit-box-shadow:none;
					box-shadow:none; }
	  .bp3-input.bp3-intent-warning{
		-webkit-box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-warning:focus{
		  -webkit-box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-warning[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #d9822b;
				  box-shadow:inset 0 0 0 1px #d9822b; }
		.bp3-input.bp3-intent-warning:disabled, .bp3-input.bp3-intent-warning.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
		.bp3-dark .bp3-input.bp3-intent-warning{
		  -webkit-box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), 0 0 0 0 rgba(217, 130, 43, 0), inset 0 0 0 1px #d9822b, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-warning:focus{
			-webkit-box-shadow:0 0 0 1px #d9822b, 0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
					box-shadow:0 0 0 1px #d9822b, 0 0 0 1px #d9822b, 0 0 0 3px rgba(217, 130, 43, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-warning[readonly]{
			-webkit-box-shadow:inset 0 0 0 1px #d9822b;
					box-shadow:inset 0 0 0 1px #d9822b; }
		  .bp3-dark .bp3-input.bp3-intent-warning:disabled, .bp3-dark .bp3-input.bp3-intent-warning.bp3-disabled{
			-webkit-box-shadow:none;
					box-shadow:none; }
	  .bp3-input.bp3-intent-danger{
		-webkit-box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.15), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-danger:focus{
		  -webkit-box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-input.bp3-intent-danger[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px #db3737;
				  box-shadow:inset 0 0 0 1px #db3737; }
		.bp3-input.bp3-intent-danger:disabled, .bp3-input.bp3-intent-danger.bp3-disabled{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
		.bp3-dark .bp3-input.bp3-intent-danger{
		  -webkit-box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), 0 0 0 0 rgba(219, 55, 55, 0), inset 0 0 0 1px #db3737, inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-danger:focus{
			-webkit-box-shadow:0 0 0 1px #db3737, 0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
					box-shadow:0 0 0 1px #db3737, 0 0 0 1px #db3737, 0 0 0 3px rgba(219, 55, 55, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		  .bp3-dark .bp3-input.bp3-intent-danger[readonly]{
			-webkit-box-shadow:inset 0 0 0 1px #db3737;
					box-shadow:inset 0 0 0 1px #db3737; }
		  .bp3-dark .bp3-input.bp3-intent-danger:disabled, .bp3-dark .bp3-input.bp3-intent-danger.bp3-disabled{
			-webkit-box-shadow:none;
					box-shadow:none; }
	  .bp3-input::-ms-clear{
		display:none; }
	textarea.bp3-input{
	  max-width:100%;
	  padding:10px; }
	  textarea.bp3-input, textarea.bp3-input.bp3-large, textarea.bp3-input.bp3-small{
		height:auto;
		line-height:inherit; }
	  textarea.bp3-input.bp3-small{
		padding:8px; }
	  .bp3-dark textarea.bp3-input{
		background:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), 0 0 0 0 rgba(19, 124, 189, 0), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
		.bp3-dark textarea.bp3-input::-webkit-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark textarea.bp3-input::-moz-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark textarea.bp3-input:-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark textarea.bp3-input::-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark textarea.bp3-input::placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark textarea.bp3-input:focus{
		  -webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark textarea.bp3-input[readonly]{
		  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark textarea.bp3-input:disabled, .bp3-dark textarea.bp3-input.bp3-disabled{
		  background:rgba(57, 75, 89, 0.5);
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(167, 182, 194, 0.6); }
	label.bp3-label{
	  display:block;
	  margin-bottom:15px;
	  margin-top:0; }
	  label.bp3-label .bp3-html-select,
	  label.bp3-label .bp3-input,
	  label.bp3-label .bp3-select,
	  label.bp3-label .bp3-slider,
	  label.bp3-label .bp3-popover-wrapper{
		display:block;
		margin-top:5px;
		text-transform:none; }
	  label.bp3-label .bp3-button-group{
		margin-top:5px; }
	  label.bp3-label .bp3-select select,
	  label.bp3-label .bp3-html-select select{
		font-weight:400;
		vertical-align:top;
		width:100%; }
	  label.bp3-label.bp3-disabled,
	  label.bp3-label.bp3-disabled .bp3-text-muted{
		color:rgba(92, 112, 128, 0.6); }
	  label.bp3-label.bp3-inline{
		line-height:30px; }
		label.bp3-label.bp3-inline .bp3-html-select,
		label.bp3-label.bp3-inline .bp3-input,
		label.bp3-label.bp3-inline .bp3-input-group,
		label.bp3-label.bp3-inline .bp3-select,
		label.bp3-label.bp3-inline .bp3-popover-wrapper{
		  display:inline-block;
		  margin:0 0 0 5px;
		  vertical-align:top; }
		label.bp3-label.bp3-inline .bp3-button-group{
		  margin:0 0 0 5px; }
		label.bp3-label.bp3-inline .bp3-input-group .bp3-input{
		  margin-left:0; }
		label.bp3-label.bp3-inline.bp3-large{
		  line-height:40px; }
	  label.bp3-label:not(.bp3-inline) .bp3-popover-target{
		display:block; }
	  .bp3-dark label.bp3-label{
		color:#f5f8fa; }
		.bp3-dark label.bp3-label.bp3-disabled,
		.bp3-dark label.bp3-label.bp3-disabled .bp3-text-muted{
		  color:rgba(167, 182, 194, 0.6); }
	.bp3-numeric-input .bp3-button-group.bp3-vertical > .bp3-button{
	  -webkit-box-flex:1;
		  -ms-flex:1 1 14px;
			  flex:1 1 14px;
	  min-height:0;
	  padding:0;
	  width:30px; }
	  .bp3-numeric-input .bp3-button-group.bp3-vertical > .bp3-button:first-child{
		border-radius:0 3px 0 0; }
	  .bp3-numeric-input .bp3-button-group.bp3-vertical > .bp3-button:last-child{
		border-radius:0 0 3px 0; }
	
	.bp3-numeric-input .bp3-button-group.bp3-vertical:first-child > .bp3-button:first-child{
	  border-radius:3px 0 0 0; }
	
	.bp3-numeric-input .bp3-button-group.bp3-vertical:first-child > .bp3-button:last-child{
	  border-radius:0 0 0 3px; }
	
	.bp3-numeric-input.bp3-large .bp3-button-group.bp3-vertical > .bp3-button{
	  width:40px; }
	
	form{
	  display:block; }
	.bp3-html-select select,
	.bp3-select select{
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  border:none;
	  border-radius:3px;
	  cursor:pointer;
	  font-size:14px;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  padding:5px 10px;
	  text-align:left;
	  vertical-align:middle;
	  background-color:#f5f8fa;
	  background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.8)), to(rgba(255, 255, 255, 0)));
	  background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0));
	  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
			  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
	  color:#182026;
	  -moz-appearance:none;
	  -webkit-appearance:none;
	  border-radius:3px;
	  height:30px;
	  padding:0 25px 0 10px;
	  width:100%; }
	  .bp3-html-select select > *, .bp3-select select > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-html-select select > .bp3-fill, .bp3-select select > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-html-select select::before,
	  .bp3-select select::before, .bp3-html-select select > *, .bp3-select select > *{
		margin-right:7px; }
	  .bp3-html-select select:empty::before,
	  .bp3-select select:empty::before,
	  .bp3-html-select select > :last-child,
	  .bp3-select select > :last-child{
		margin-right:0; }
	  .bp3-html-select select:hover,
	  .bp3-select select:hover{
		background-clip:padding-box;
		background-color:#ebf1f5;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
	  .bp3-html-select select:active,
	  .bp3-select select:active, .bp3-html-select select.bp3-active,
	  .bp3-select select.bp3-active{
		background-color:#d8e1e8;
		background-image:none;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-html-select select:disabled,
	  .bp3-select select:disabled, .bp3-html-select select.bp3-disabled,
	  .bp3-select select.bp3-disabled{
		background-color:rgba(206, 217, 224, 0.5);
		background-image:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed;
		outline:none; }
		.bp3-html-select select:disabled.bp3-active,
		.bp3-select select:disabled.bp3-active, .bp3-html-select select:disabled.bp3-active:hover,
		.bp3-select select:disabled.bp3-active:hover, .bp3-html-select select.bp3-disabled.bp3-active,
		.bp3-select select.bp3-disabled.bp3-active, .bp3-html-select select.bp3-disabled.bp3-active:hover,
		.bp3-select select.bp3-disabled.bp3-active:hover{
		  background:rgba(206, 217, 224, 0.7); }
	
	.bp3-html-select.bp3-minimal select,
	.bp3-select.bp3-minimal select{
	  background:none;
	  -webkit-box-shadow:none;
			  box-shadow:none; }
	  .bp3-html-select.bp3-minimal select:hover,
	  .bp3-select.bp3-minimal select:hover{
		background:rgba(167, 182, 194, 0.3);
		-webkit-box-shadow:none;
				box-shadow:none;
		color:#182026;
		text-decoration:none; }
	  .bp3-html-select.bp3-minimal select:active,
	  .bp3-select.bp3-minimal select:active, .bp3-html-select.bp3-minimal select.bp3-active,
	  .bp3-select.bp3-minimal select.bp3-active{
		background:rgba(115, 134, 148, 0.3);
		-webkit-box-shadow:none;
				box-shadow:none;
		color:#182026; }
	  .bp3-html-select.bp3-minimal select:disabled,
	  .bp3-select.bp3-minimal select:disabled, .bp3-html-select.bp3-minimal select:disabled:hover,
	  .bp3-select.bp3-minimal select:disabled:hover, .bp3-html-select.bp3-minimal select.bp3-disabled,
	  .bp3-select.bp3-minimal select.bp3-disabled, .bp3-html-select.bp3-minimal select.bp3-disabled:hover,
	  .bp3-select.bp3-minimal select.bp3-disabled:hover{
		background:none;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed; }
		.bp3-html-select.bp3-minimal select:disabled.bp3-active,
		.bp3-select.bp3-minimal select:disabled.bp3-active, .bp3-html-select.bp3-minimal select:disabled:hover.bp3-active,
		.bp3-select.bp3-minimal select:disabled:hover.bp3-active, .bp3-html-select.bp3-minimal select.bp3-disabled.bp3-active,
		.bp3-select.bp3-minimal select.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal select.bp3-disabled:hover.bp3-active,
		.bp3-select.bp3-minimal select.bp3-disabled:hover.bp3-active{
		  background:rgba(115, 134, 148, 0.3); }
	  .bp3-dark .bp3-html-select.bp3-minimal select, .bp3-html-select.bp3-minimal .bp3-dark select,
	  .bp3-dark .bp3-select.bp3-minimal select, .bp3-select.bp3-minimal .bp3-dark select{
		background:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		color:inherit; }
		.bp3-dark .bp3-html-select.bp3-minimal select:hover, .bp3-html-select.bp3-minimal .bp3-dark select:hover,
		.bp3-dark .bp3-select.bp3-minimal select:hover, .bp3-select.bp3-minimal .bp3-dark select:hover, .bp3-dark .bp3-html-select.bp3-minimal select:active, .bp3-html-select.bp3-minimal .bp3-dark select:active,
		.bp3-dark .bp3-select.bp3-minimal select:active, .bp3-select.bp3-minimal .bp3-dark select:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-active,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-active{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none; }
		.bp3-dark .bp3-html-select.bp3-minimal select:hover, .bp3-html-select.bp3-minimal .bp3-dark select:hover,
		.bp3-dark .bp3-select.bp3-minimal select:hover, .bp3-select.bp3-minimal .bp3-dark select:hover{
		  background:rgba(138, 155, 168, 0.15); }
		.bp3-dark .bp3-html-select.bp3-minimal select:active, .bp3-html-select.bp3-minimal .bp3-dark select:active,
		.bp3-dark .bp3-select.bp3-minimal select:active, .bp3-select.bp3-minimal .bp3-dark select:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-active,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-active{
		  background:rgba(138, 155, 168, 0.3);
		  color:#f5f8fa; }
		.bp3-dark .bp3-html-select.bp3-minimal select:disabled, .bp3-html-select.bp3-minimal .bp3-dark select:disabled,
		.bp3-dark .bp3-select.bp3-minimal select:disabled, .bp3-select.bp3-minimal .bp3-dark select:disabled, .bp3-dark .bp3-html-select.bp3-minimal select:disabled:hover, .bp3-html-select.bp3-minimal .bp3-dark select:disabled:hover,
		.bp3-dark .bp3-select.bp3-minimal select:disabled:hover, .bp3-select.bp3-minimal .bp3-dark select:disabled:hover, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-disabled,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-disabled, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-disabled:hover, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-disabled:hover,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-disabled:hover, .bp3-select.bp3-minimal .bp3-dark select.bp3-disabled:hover{
		  background:none;
		  color:rgba(167, 182, 194, 0.6);
		  cursor:not-allowed; }
		  .bp3-dark .bp3-html-select.bp3-minimal select:disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select:disabled.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select:disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select:disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select:disabled:hover.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select:disabled:hover.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select:disabled:hover.bp3-active, .bp3-select.bp3-minimal .bp3-dark select:disabled:hover.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-disabled.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-disabled:hover.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-disabled:hover.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-disabled:hover.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-disabled:hover.bp3-active{
			background:rgba(138, 155, 168, 0.3); }
	  .bp3-html-select.bp3-minimal select.bp3-intent-primary,
	  .bp3-select.bp3-minimal select.bp3-intent-primary{
		color:#106ba3; }
		.bp3-html-select.bp3-minimal select.bp3-intent-primary:hover,
		.bp3-select.bp3-minimal select.bp3-intent-primary:hover, .bp3-html-select.bp3-minimal select.bp3-intent-primary:active,
		.bp3-select.bp3-minimal select.bp3-intent-primary:active, .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-primary.bp3-active{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#106ba3; }
		.bp3-html-select.bp3-minimal select.bp3-intent-primary:hover,
		.bp3-select.bp3-minimal select.bp3-intent-primary:hover{
		  background:rgba(19, 124, 189, 0.15);
		  color:#106ba3; }
		.bp3-html-select.bp3-minimal select.bp3-intent-primary:active,
		.bp3-select.bp3-minimal select.bp3-intent-primary:active, .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-primary.bp3-active{
		  background:rgba(19, 124, 189, 0.3);
		  color:#106ba3; }
		.bp3-html-select.bp3-minimal select.bp3-intent-primary:disabled,
		.bp3-select.bp3-minimal select.bp3-intent-primary:disabled, .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-disabled,
		.bp3-select.bp3-minimal select.bp3-intent-primary.bp3-disabled{
		  background:none;
		  color:rgba(16, 107, 163, 0.5); }
		  .bp3-html-select.bp3-minimal select.bp3-intent-primary:disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-primary:disabled.bp3-active, .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-primary.bp3-disabled.bp3-active{
			background:rgba(19, 124, 189, 0.3); }
		.bp3-html-select.bp3-minimal select.bp3-intent-primary .bp3-button-spinner .bp3-spinner-head, .bp3-select.bp3-minimal select.bp3-intent-primary .bp3-button-spinner .bp3-spinner-head{
		  stroke:#106ba3; }
		.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary{
		  color:#48aff0; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary:hover, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary:hover,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary:hover, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary:hover{
			background:rgba(19, 124, 189, 0.2);
			color:#48aff0; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary:active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary:active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary:active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-active{
			background:rgba(19, 124, 189, 0.3);
			color:#48aff0; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary:disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary:disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary:disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary:disabled, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary.bp3-disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-disabled{
			background:none;
			color:rgba(72, 175, 240, 0.5); }
			.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary:disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary:disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary:disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary:disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-primary.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-primary.bp3-disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-primary.bp3-disabled.bp3-active{
			  background:rgba(19, 124, 189, 0.3); }
	  .bp3-html-select.bp3-minimal select.bp3-intent-success,
	  .bp3-select.bp3-minimal select.bp3-intent-success{
		color:#0d8050; }
		.bp3-html-select.bp3-minimal select.bp3-intent-success:hover,
		.bp3-select.bp3-minimal select.bp3-intent-success:hover, .bp3-html-select.bp3-minimal select.bp3-intent-success:active,
		.bp3-select.bp3-minimal select.bp3-intent-success:active, .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-success.bp3-active{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#0d8050; }
		.bp3-html-select.bp3-minimal select.bp3-intent-success:hover,
		.bp3-select.bp3-minimal select.bp3-intent-success:hover{
		  background:rgba(15, 153, 96, 0.15);
		  color:#0d8050; }
		.bp3-html-select.bp3-minimal select.bp3-intent-success:active,
		.bp3-select.bp3-minimal select.bp3-intent-success:active, .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-success.bp3-active{
		  background:rgba(15, 153, 96, 0.3);
		  color:#0d8050; }
		.bp3-html-select.bp3-minimal select.bp3-intent-success:disabled,
		.bp3-select.bp3-minimal select.bp3-intent-success:disabled, .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-disabled,
		.bp3-select.bp3-minimal select.bp3-intent-success.bp3-disabled{
		  background:none;
		  color:rgba(13, 128, 80, 0.5); }
		  .bp3-html-select.bp3-minimal select.bp3-intent-success:disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-success:disabled.bp3-active, .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-success.bp3-disabled.bp3-active{
			background:rgba(15, 153, 96, 0.3); }
		.bp3-html-select.bp3-minimal select.bp3-intent-success .bp3-button-spinner .bp3-spinner-head, .bp3-select.bp3-minimal select.bp3-intent-success .bp3-button-spinner .bp3-spinner-head{
		  stroke:#0d8050; }
		.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success{
		  color:#3dcc91; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success:hover, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success:hover,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success:hover, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success:hover{
			background:rgba(15, 153, 96, 0.2);
			color:#3dcc91; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success:active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success:active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success:active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-active{
			background:rgba(15, 153, 96, 0.3);
			color:#3dcc91; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success:disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success:disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success:disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success:disabled, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success.bp3-disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-disabled{
			background:none;
			color:rgba(61, 204, 145, 0.5); }
			.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success:disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success:disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success:disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success:disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-success.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-success.bp3-disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-success.bp3-disabled.bp3-active{
			  background:rgba(15, 153, 96, 0.3); }
	  .bp3-html-select.bp3-minimal select.bp3-intent-warning,
	  .bp3-select.bp3-minimal select.bp3-intent-warning{
		color:#bf7326; }
		.bp3-html-select.bp3-minimal select.bp3-intent-warning:hover,
		.bp3-select.bp3-minimal select.bp3-intent-warning:hover, .bp3-html-select.bp3-minimal select.bp3-intent-warning:active,
		.bp3-select.bp3-minimal select.bp3-intent-warning:active, .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-warning.bp3-active{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#bf7326; }
		.bp3-html-select.bp3-minimal select.bp3-intent-warning:hover,
		.bp3-select.bp3-minimal select.bp3-intent-warning:hover{
		  background:rgba(217, 130, 43, 0.15);
		  color:#bf7326; }
		.bp3-html-select.bp3-minimal select.bp3-intent-warning:active,
		.bp3-select.bp3-minimal select.bp3-intent-warning:active, .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-warning.bp3-active{
		  background:rgba(217, 130, 43, 0.3);
		  color:#bf7326; }
		.bp3-html-select.bp3-minimal select.bp3-intent-warning:disabled,
		.bp3-select.bp3-minimal select.bp3-intent-warning:disabled, .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-disabled,
		.bp3-select.bp3-minimal select.bp3-intent-warning.bp3-disabled{
		  background:none;
		  color:rgba(191, 115, 38, 0.5); }
		  .bp3-html-select.bp3-minimal select.bp3-intent-warning:disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-warning:disabled.bp3-active, .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-warning.bp3-disabled.bp3-active{
			background:rgba(217, 130, 43, 0.3); }
		.bp3-html-select.bp3-minimal select.bp3-intent-warning .bp3-button-spinner .bp3-spinner-head, .bp3-select.bp3-minimal select.bp3-intent-warning .bp3-button-spinner .bp3-spinner-head{
		  stroke:#bf7326; }
		.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning{
		  color:#ffb366; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning:hover, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning:hover,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning:hover, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning:hover{
			background:rgba(217, 130, 43, 0.2);
			color:#ffb366; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning:active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning:active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning:active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-active{
			background:rgba(217, 130, 43, 0.3);
			color:#ffb366; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning:disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning:disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning:disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning:disabled, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning.bp3-disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-disabled{
			background:none;
			color:rgba(255, 179, 102, 0.5); }
			.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning:disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning:disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning:disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning:disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-warning.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-warning.bp3-disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-warning.bp3-disabled.bp3-active{
			  background:rgba(217, 130, 43, 0.3); }
	  .bp3-html-select.bp3-minimal select.bp3-intent-danger,
	  .bp3-select.bp3-minimal select.bp3-intent-danger{
		color:#c23030; }
		.bp3-html-select.bp3-minimal select.bp3-intent-danger:hover,
		.bp3-select.bp3-minimal select.bp3-intent-danger:hover, .bp3-html-select.bp3-minimal select.bp3-intent-danger:active,
		.bp3-select.bp3-minimal select.bp3-intent-danger:active, .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-danger.bp3-active{
		  background:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:#c23030; }
		.bp3-html-select.bp3-minimal select.bp3-intent-danger:hover,
		.bp3-select.bp3-minimal select.bp3-intent-danger:hover{
		  background:rgba(219, 55, 55, 0.15);
		  color:#c23030; }
		.bp3-html-select.bp3-minimal select.bp3-intent-danger:active,
		.bp3-select.bp3-minimal select.bp3-intent-danger:active, .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-active,
		.bp3-select.bp3-minimal select.bp3-intent-danger.bp3-active{
		  background:rgba(219, 55, 55, 0.3);
		  color:#c23030; }
		.bp3-html-select.bp3-minimal select.bp3-intent-danger:disabled,
		.bp3-select.bp3-minimal select.bp3-intent-danger:disabled, .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-disabled,
		.bp3-select.bp3-minimal select.bp3-intent-danger.bp3-disabled{
		  background:none;
		  color:rgba(194, 48, 48, 0.5); }
		  .bp3-html-select.bp3-minimal select.bp3-intent-danger:disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-danger:disabled.bp3-active, .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-disabled.bp3-active,
		  .bp3-select.bp3-minimal select.bp3-intent-danger.bp3-disabled.bp3-active{
			background:rgba(219, 55, 55, 0.3); }
		.bp3-html-select.bp3-minimal select.bp3-intent-danger .bp3-button-spinner .bp3-spinner-head, .bp3-select.bp3-minimal select.bp3-intent-danger .bp3-button-spinner .bp3-spinner-head{
		  stroke:#c23030; }
		.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger,
		.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger{
		  color:#ff7373; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger:hover, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger:hover,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger:hover, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger:hover{
			background:rgba(219, 55, 55, 0.2);
			color:#ff7373; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger:active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger:active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger:active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger:active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-active,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-active{
			background:rgba(219, 55, 55, 0.3);
			color:#ff7373; }
		  .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger:disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger:disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger:disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger:disabled, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-disabled, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-disabled,
		  .bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger.bp3-disabled, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-disabled{
			background:none;
			color:rgba(255, 115, 115, 0.5); }
			.bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger:disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger:disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger:disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger:disabled.bp3-active, .bp3-dark .bp3-html-select.bp3-minimal select.bp3-intent-danger.bp3-disabled.bp3-active, .bp3-html-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-disabled.bp3-active,
			.bp3-dark .bp3-select.bp3-minimal select.bp3-intent-danger.bp3-disabled.bp3-active, .bp3-select.bp3-minimal .bp3-dark select.bp3-intent-danger.bp3-disabled.bp3-active{
			  background:rgba(219, 55, 55, 0.3); }
	
	.bp3-html-select.bp3-large select,
	.bp3-select.bp3-large select{
	  font-size:16px;
	  height:40px;
	  padding-right:35px; }
	
	.bp3-dark .bp3-html-select select, .bp3-dark .bp3-select select{
	  background-color:#394b59;
	  background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.05)), to(rgba(255, 255, 255, 0)));
	  background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0));
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
	  color:#f5f8fa; }
	  .bp3-dark .bp3-html-select select:hover, .bp3-dark .bp3-select select:hover, .bp3-dark .bp3-html-select select:active, .bp3-dark .bp3-select select:active, .bp3-dark .bp3-html-select select.bp3-active, .bp3-dark .bp3-select select.bp3-active{
		color:#f5f8fa; }
	  .bp3-dark .bp3-html-select select:hover, .bp3-dark .bp3-select select:hover{
		background-color:#30404d;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-html-select select:active, .bp3-dark .bp3-select select:active, .bp3-dark .bp3-html-select select.bp3-active, .bp3-dark .bp3-select select.bp3-active{
		background-color:#202b33;
		background-image:none;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-html-select select:disabled, .bp3-dark .bp3-select select:disabled, .bp3-dark .bp3-html-select select.bp3-disabled, .bp3-dark .bp3-select select.bp3-disabled{
		background-color:rgba(57, 75, 89, 0.5);
		background-image:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-html-select select:disabled.bp3-active, .bp3-dark .bp3-select select:disabled.bp3-active, .bp3-dark .bp3-html-select select.bp3-disabled.bp3-active, .bp3-dark .bp3-select select.bp3-disabled.bp3-active{
		  background:rgba(57, 75, 89, 0.7); }
	  .bp3-dark .bp3-html-select select .bp3-button-spinner .bp3-spinner-head, .bp3-dark .bp3-select select .bp3-button-spinner .bp3-spinner-head{
		background:rgba(16, 22, 26, 0.5);
		stroke:#8a9ba8; }
	
	.bp3-html-select select:disabled,
	.bp3-select select:disabled{
	  background-color:rgba(206, 217, 224, 0.5);
	  -webkit-box-shadow:none;
			  box-shadow:none;
	  color:rgba(92, 112, 128, 0.6);
	  cursor:not-allowed; }
	
	.bp3-html-select .bp3-icon,
	.bp3-select .bp3-icon, .bp3-select::after{
	  color:#5c7080;
	  pointer-events:none;
	  position:absolute;
	  right:7px;
	  top:7px; }
	  .bp3-html-select .bp3-disabled.bp3-icon,
	  .bp3-select .bp3-disabled.bp3-icon, .bp3-disabled.bp3-select::after{
		color:rgba(92, 112, 128, 0.6); }
	.bp3-html-select,
	.bp3-select{
	  display:inline-block;
	  letter-spacing:normal;
	  position:relative;
	  vertical-align:middle; }
	  .bp3-html-select select::-ms-expand,
	  .bp3-select select::-ms-expand{
		display:none; }
	  .bp3-html-select .bp3-icon,
	  .bp3-select .bp3-icon{
		color:#5c7080; }
		.bp3-html-select .bp3-icon:hover,
		.bp3-select .bp3-icon:hover{
		  color:#182026; }
		.bp3-dark .bp3-html-select .bp3-icon, .bp3-dark
		.bp3-select .bp3-icon{
		  color:#a7b6c2; }
		  .bp3-dark .bp3-html-select .bp3-icon:hover, .bp3-dark
		  .bp3-select .bp3-icon:hover{
			color:#f5f8fa; }
	  .bp3-html-select.bp3-large::after,
	  .bp3-html-select.bp3-large .bp3-icon,
	  .bp3-select.bp3-large::after,
	  .bp3-select.bp3-large .bp3-icon{
		right:12px;
		top:12px; }
	  .bp3-html-select.bp3-fill,
	  .bp3-html-select.bp3-fill select,
	  .bp3-select.bp3-fill,
	  .bp3-select.bp3-fill select{
		width:100%; }
	  .bp3-dark .bp3-html-select option, .bp3-dark
	  .bp3-select option{
		background-color:#30404d;
		color:#f5f8fa; }
	  .bp3-dark .bp3-html-select option:disabled, .bp3-dark
	  .bp3-select option:disabled{
		color:rgba(167, 182, 194, 0.6); }
	  .bp3-dark .bp3-html-select::after, .bp3-dark
	  .bp3-select::after{
		color:#a7b6c2; }
	
	.bp3-select::after{
	  font-family:"Icons16", sans-serif;
	  font-size:16px;
	  font-style:normal;
	  font-weight:400;
	  line-height:1;
	  -moz-osx-font-smoothing:grayscale;
	  -webkit-font-smoothing:antialiased;
	  content:""; }
	.bp3-running-text table, table.bp3-html-table{
	  border-spacing:0;
	  font-size:14px; }
	  .bp3-running-text table th, table.bp3-html-table th,
	  .bp3-running-text table td,
	  table.bp3-html-table td{
		padding:11px;
		text-align:left;
		vertical-align:top; }
	  .bp3-running-text table th, table.bp3-html-table th{
		color:#182026;
		font-weight:600; }
	  
	  .bp3-running-text table td,
	  table.bp3-html-table td{
		color:#182026; }
	  .bp3-running-text table tbody tr:first-child th, table.bp3-html-table tbody tr:first-child th,
	  .bp3-running-text table tbody tr:first-child td,
	  table.bp3-html-table tbody tr:first-child td{
		-webkit-box-shadow:inset 0 1px 0 0 rgba(16, 22, 26, 0.15);
				box-shadow:inset 0 1px 0 0 rgba(16, 22, 26, 0.15); }
	  .bp3-dark .bp3-running-text table th, .bp3-running-text .bp3-dark table th, .bp3-dark table.bp3-html-table th{
		color:#f5f8fa; }
	  .bp3-dark .bp3-running-text table td, .bp3-running-text .bp3-dark table td, .bp3-dark table.bp3-html-table td{
		color:#f5f8fa; }
	  .bp3-dark .bp3-running-text table tbody tr:first-child th, .bp3-running-text .bp3-dark table tbody tr:first-child th, .bp3-dark table.bp3-html-table tbody tr:first-child th,
	  .bp3-dark .bp3-running-text table tbody tr:first-child td,
	  .bp3-running-text .bp3-dark table tbody tr:first-child td,
	  .bp3-dark table.bp3-html-table tbody tr:first-child td{
		-webkit-box-shadow:inset 0 1px 0 0 rgba(255, 255, 255, 0.15);
				box-shadow:inset 0 1px 0 0 rgba(255, 255, 255, 0.15); }
	
	table.bp3-html-table.bp3-html-table-condensed th,
	table.bp3-html-table.bp3-html-table-condensed td, table.bp3-html-table.bp3-small th,
	table.bp3-html-table.bp3-small td{
	  padding-bottom:6px;
	  padding-top:6px; }
	
	table.bp3-html-table.bp3-html-table-striped tbody tr:nth-child(odd) td{
	  background:rgba(191, 204, 214, 0.15); }
	
	table.bp3-html-table.bp3-html-table-bordered th:not(:first-child){
	  -webkit-box-shadow:inset 1px 0 0 0 rgba(16, 22, 26, 0.15);
			  box-shadow:inset 1px 0 0 0 rgba(16, 22, 26, 0.15); }
	
	table.bp3-html-table.bp3-html-table-bordered tbody tr td{
	  -webkit-box-shadow:inset 0 1px 0 0 rgba(16, 22, 26, 0.15);
			  box-shadow:inset 0 1px 0 0 rgba(16, 22, 26, 0.15); }
	  table.bp3-html-table.bp3-html-table-bordered tbody tr td:not(:first-child){
		-webkit-box-shadow:inset 1px 1px 0 0 rgba(16, 22, 26, 0.15);
				box-shadow:inset 1px 1px 0 0 rgba(16, 22, 26, 0.15); }
	
	table.bp3-html-table.bp3-html-table-bordered.bp3-html-table-striped tbody tr:not(:first-child) td{
	  -webkit-box-shadow:none;
			  box-shadow:none; }
	  table.bp3-html-table.bp3-html-table-bordered.bp3-html-table-striped tbody tr:not(:first-child) td:not(:first-child){
		-webkit-box-shadow:inset 1px 0 0 0 rgba(16, 22, 26, 0.15);
				box-shadow:inset 1px 0 0 0 rgba(16, 22, 26, 0.15); }
	
	table.bp3-html-table.bp3-interactive tbody tr:hover td{
	  background-color:rgba(191, 204, 214, 0.3);
	  cursor:pointer; }
	
	table.bp3-html-table.bp3-interactive tbody tr:active td{
	  background-color:rgba(191, 204, 214, 0.4); }
	
	.bp3-dark table.bp3-html-table{ }
	  .bp3-dark table.bp3-html-table.bp3-html-table-striped tbody tr:nth-child(odd) td{
		background:rgba(92, 112, 128, 0.15); }
	  .bp3-dark table.bp3-html-table.bp3-html-table-bordered th:not(:first-child){
		-webkit-box-shadow:inset 1px 0 0 0 rgba(255, 255, 255, 0.15);
				box-shadow:inset 1px 0 0 0 rgba(255, 255, 255, 0.15); }
	  .bp3-dark table.bp3-html-table.bp3-html-table-bordered tbody tr td{
		-webkit-box-shadow:inset 0 1px 0 0 rgba(255, 255, 255, 0.15);
				box-shadow:inset 0 1px 0 0 rgba(255, 255, 255, 0.15); }
		.bp3-dark table.bp3-html-table.bp3-html-table-bordered tbody tr td:not(:first-child){
		  -webkit-box-shadow:inset 1px 1px 0 0 rgba(255, 255, 255, 0.15);
				  box-shadow:inset 1px 1px 0 0 rgba(255, 255, 255, 0.15); }
	  .bp3-dark table.bp3-html-table.bp3-html-table-bordered.bp3-html-table-striped tbody tr:not(:first-child) td{
		-webkit-box-shadow:inset 1px 0 0 0 rgba(255, 255, 255, 0.15);
				box-shadow:inset 1px 0 0 0 rgba(255, 255, 255, 0.15); }
		.bp3-dark table.bp3-html-table.bp3-html-table-bordered.bp3-html-table-striped tbody tr:not(:first-child) td:first-child{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-dark table.bp3-html-table.bp3-interactive tbody tr:hover td{
		background-color:rgba(92, 112, 128, 0.3);
		cursor:pointer; }
	  .bp3-dark table.bp3-html-table.bp3-interactive tbody tr:active td{
		background-color:rgba(92, 112, 128, 0.4); }
	
	.bp3-key-combo{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center; }
	  .bp3-key-combo > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-key-combo > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-key-combo::before,
	  .bp3-key-combo > *{
		margin-right:5px; }
	  .bp3-key-combo:empty::before,
	  .bp3-key-combo > :last-child{
		margin-right:0; }
	
	.bp3-hotkey-dialog{
	  padding-bottom:0;
	  top:40px; }
	  .bp3-hotkey-dialog .bp3-dialog-body{
		margin:0;
		padding:0; }
	  .bp3-hotkey-dialog .bp3-hotkey-label{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1; }
	
	.bp3-hotkey-column{
	  margin:auto;
	  max-height:80vh;
	  overflow-y:auto;
	  padding:30px; }
	  .bp3-hotkey-column .bp3-heading{
		margin-bottom:20px; }
		.bp3-hotkey-column .bp3-heading:not(:first-child){
		  margin-top:40px; }
	
	.bp3-hotkey{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-pack:justify;
		  -ms-flex-pack:justify;
			  justify-content:space-between;
	  margin-left:0;
	  margin-right:0; }
	  .bp3-hotkey:not(:last-child){
		margin-bottom:10px; }
	.bp3-icon{
	  display:inline-block;
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  vertical-align:text-bottom; }
	  .bp3-icon:not(:empty)::before{
		content:"" !important;
		content:unset !important; }
	  .bp3-icon > svg{
		display:block; }
		.bp3-icon > svg:not([fill]){
		  fill:currentColor; }
	
	.bp3-icon.bp3-intent-primary, .bp3-icon-standard.bp3-intent-primary, .bp3-icon-large.bp3-intent-primary{
	  color:#106ba3; }
	  .bp3-dark .bp3-icon.bp3-intent-primary, .bp3-dark .bp3-icon-standard.bp3-intent-primary, .bp3-dark .bp3-icon-large.bp3-intent-primary{
		color:#48aff0; }
	
	.bp3-icon.bp3-intent-success, .bp3-icon-standard.bp3-intent-success, .bp3-icon-large.bp3-intent-success{
	  color:#0d8050; }
	  .bp3-dark .bp3-icon.bp3-intent-success, .bp3-dark .bp3-icon-standard.bp3-intent-success, .bp3-dark .bp3-icon-large.bp3-intent-success{
		color:#3dcc91; }
	
	.bp3-icon.bp3-intent-warning, .bp3-icon-standard.bp3-intent-warning, .bp3-icon-large.bp3-intent-warning{
	  color:#bf7326; }
	  .bp3-dark .bp3-icon.bp3-intent-warning, .bp3-dark .bp3-icon-standard.bp3-intent-warning, .bp3-dark .bp3-icon-large.bp3-intent-warning{
		color:#ffb366; }
	
	.bp3-icon.bp3-intent-danger, .bp3-icon-standard.bp3-intent-danger, .bp3-icon-large.bp3-intent-danger{
	  color:#c23030; }
	  .bp3-dark .bp3-icon.bp3-intent-danger, .bp3-dark .bp3-icon-standard.bp3-intent-danger, .bp3-dark .bp3-icon-large.bp3-intent-danger{
		color:#ff7373; }
	
	span.bp3-icon-standard{
	  font-family:"Icons16", sans-serif;
	  font-size:16px;
	  font-style:normal;
	  font-weight:400;
	  line-height:1;
	  -moz-osx-font-smoothing:grayscale;
	  -webkit-font-smoothing:antialiased;
	  display:inline-block; }
	
	span.bp3-icon-large{
	  font-family:"Icons20", sans-serif;
	  font-size:20px;
	  font-style:normal;
	  font-weight:400;
	  line-height:1;
	  -moz-osx-font-smoothing:grayscale;
	  -webkit-font-smoothing:antialiased;
	  display:inline-block; }
	
	span.bp3-icon:empty{
	  font-family:"Icons20";
	  font-size:inherit;
	  font-style:normal;
	  font-weight:400;
	  line-height:1; }
	  span.bp3-icon:empty::before{
		-moz-osx-font-smoothing:grayscale;
		-webkit-font-smoothing:antialiased; }
	
	.bp3-icon-add::before{
	  content:""; }
	
	.bp3-icon-add-column-left::before{
	  content:""; }
	
	.bp3-icon-add-column-right::before{
	  content:""; }
	
	.bp3-icon-add-row-bottom::before{
	  content:""; }
	
	.bp3-icon-add-row-top::before{
	  content:""; }
	
	.bp3-icon-add-to-artifact::before{
	  content:""; }
	
	.bp3-icon-add-to-folder::before{
	  content:""; }
	
	.bp3-icon-airplane::before{
	  content:""; }
	
	.bp3-icon-align-center::before{
	  content:""; }
	
	.bp3-icon-align-justify::before{
	  content:""; }
	
	.bp3-icon-align-left::before{
	  content:""; }
	
	.bp3-icon-align-right::before{
	  content:""; }
	
	.bp3-icon-alignment-bottom::before{
	  content:""; }
	
	.bp3-icon-alignment-horizontal-center::before{
	  content:""; }
	
	.bp3-icon-alignment-left::before{
	  content:""; }
	
	.bp3-icon-alignment-right::before{
	  content:""; }
	
	.bp3-icon-alignment-top::before{
	  content:""; }
	
	.bp3-icon-alignment-vertical-center::before{
	  content:""; }
	
	.bp3-icon-annotation::before{
	  content:""; }
	
	.bp3-icon-application::before{
	  content:""; }
	
	.bp3-icon-applications::before{
	  content:""; }
	
	.bp3-icon-archive::before{
	  content:""; }
	
	.bp3-icon-arrow-bottom-left::before{
	  content:"↙"; }
	
	.bp3-icon-arrow-bottom-right::before{
	  content:"↘"; }
	
	.bp3-icon-arrow-down::before{
	  content:"↓"; }
	
	.bp3-icon-arrow-left::before{
	  content:"←"; }
	
	.bp3-icon-arrow-right::before{
	  content:"→"; }
	
	.bp3-icon-arrow-top-left::before{
	  content:"↖"; }
	
	.bp3-icon-arrow-top-right::before{
	  content:"↗"; }
	
	.bp3-icon-arrow-up::before{
	  content:"↑"; }
	
	.bp3-icon-arrows-horizontal::before{
	  content:"↔"; }
	
	.bp3-icon-arrows-vertical::before{
	  content:"↕"; }
	
	.bp3-icon-asterisk::before{
	  content:"*"; }
	
	.bp3-icon-automatic-updates::before{
	  content:""; }
	
	.bp3-icon-badge::before{
	  content:""; }
	
	.bp3-icon-ban-circle::before{
	  content:""; }
	
	.bp3-icon-bank-account::before{
	  content:""; }
	
	.bp3-icon-barcode::before{
	  content:""; }
	
	.bp3-icon-blank::before{
	  content:""; }
	
	.bp3-icon-blocked-person::before{
	  content:""; }
	
	.bp3-icon-bold::before{
	  content:""; }
	
	.bp3-icon-book::before{
	  content:""; }
	
	.bp3-icon-bookmark::before{
	  content:""; }
	
	.bp3-icon-box::before{
	  content:""; }
	
	.bp3-icon-briefcase::before{
	  content:""; }
	
	.bp3-icon-bring-data::before{
	  content:""; }
	
	.bp3-icon-build::before{
	  content:""; }
	
	.bp3-icon-calculator::before{
	  content:""; }
	
	.bp3-icon-calendar::before{
	  content:""; }
	
	.bp3-icon-camera::before{
	  content:""; }
	
	.bp3-icon-caret-down::before{
	  content:"⌄"; }
	
	.bp3-icon-caret-left::before{
	  content:"〈"; }
	
	.bp3-icon-caret-right::before{
	  content:"〉"; }
	
	.bp3-icon-caret-up::before{
	  content:"⌃"; }
	
	.bp3-icon-cell-tower::before{
	  content:""; }
	
	.bp3-icon-changes::before{
	  content:""; }
	
	.bp3-icon-chart::before{
	  content:""; }
	
	.bp3-icon-chat::before{
	  content:""; }
	
	.bp3-icon-chevron-backward::before{
	  content:""; }
	
	.bp3-icon-chevron-down::before{
	  content:""; }
	
	.bp3-icon-chevron-forward::before{
	  content:""; }
	
	.bp3-icon-chevron-left::before{
	  content:""; }
	
	.bp3-icon-chevron-right::before{
	  content:""; }
	
	.bp3-icon-chevron-up::before{
	  content:""; }
	
	.bp3-icon-circle::before{
	  content:""; }
	
	.bp3-icon-circle-arrow-down::before{
	  content:""; }
	
	.bp3-icon-circle-arrow-left::before{
	  content:""; }
	
	.bp3-icon-circle-arrow-right::before{
	  content:""; }
	
	.bp3-icon-circle-arrow-up::before{
	  content:""; }
	
	.bp3-icon-citation::before{
	  content:""; }
	
	.bp3-icon-clean::before{
	  content:""; }
	
	.bp3-icon-clipboard::before{
	  content:""; }
	
	.bp3-icon-cloud::before{
	  content:"☁"; }
	
	.bp3-icon-cloud-download::before{
	  content:""; }
	
	.bp3-icon-cloud-upload::before{
	  content:""; }
	
	.bp3-icon-code::before{
	  content:""; }
	
	.bp3-icon-code-block::before{
	  content:""; }
	
	.bp3-icon-cog::before{
	  content:""; }
	
	.bp3-icon-collapse-all::before{
	  content:""; }
	
	.bp3-icon-column-layout::before{
	  content:""; }
	
	.bp3-icon-comment::before{
	  content:""; }
	
	.bp3-icon-comparison::before{
	  content:""; }
	
	.bp3-icon-compass::before{
	  content:""; }
	
	.bp3-icon-compressed::before{
	  content:""; }
	
	.bp3-icon-confirm::before{
	  content:""; }
	
	.bp3-icon-console::before{
	  content:""; }
	
	.bp3-icon-contrast::before{
	  content:""; }
	
	.bp3-icon-control::before{
	  content:""; }
	
	.bp3-icon-credit-card::before{
	  content:""; }
	
	.bp3-icon-cross::before{
	  content:"✗"; }
	
	.bp3-icon-crown::before{
	  content:""; }
	
	.bp3-icon-cube::before{
	  content:""; }
	
	.bp3-icon-cube-add::before{
	  content:""; }
	
	.bp3-icon-cube-remove::before{
	  content:""; }
	
	.bp3-icon-curved-range-chart::before{
	  content:""; }
	
	.bp3-icon-cut::before{
	  content:""; }
	
	.bp3-icon-dashboard::before{
	  content:""; }
	
	.bp3-icon-data-lineage::before{
	  content:""; }
	
	.bp3-icon-database::before{
	  content:""; }
	
	.bp3-icon-delete::before{
	  content:""; }
	
	.bp3-icon-delta::before{
	  content:"Δ"; }
	
	.bp3-icon-derive-column::before{
	  content:""; }
	
	.bp3-icon-desktop::before{
	  content:""; }
	
	.bp3-icon-diagnosis::before{
	  content:""; }
	
	.bp3-icon-diagram-tree::before{
	  content:""; }
	
	.bp3-icon-direction-left::before{
	  content:""; }
	
	.bp3-icon-direction-right::before{
	  content:""; }
	
	.bp3-icon-disable::before{
	  content:""; }
	
	.bp3-icon-document::before{
	  content:""; }
	
	.bp3-icon-document-open::before{
	  content:""; }
	
	.bp3-icon-document-share::before{
	  content:""; }
	
	.bp3-icon-dollar::before{
	  content:"$"; }
	
	.bp3-icon-dot::before{
	  content:"•"; }
	
	.bp3-icon-double-caret-horizontal::before{
	  content:""; }
	
	.bp3-icon-double-caret-vertical::before{
	  content:""; }
	
	.bp3-icon-double-chevron-down::before{
	  content:""; }
	
	.bp3-icon-double-chevron-left::before{
	  content:""; }
	
	.bp3-icon-double-chevron-right::before{
	  content:""; }
	
	.bp3-icon-double-chevron-up::before{
	  content:""; }
	
	.bp3-icon-doughnut-chart::before{
	  content:""; }
	
	.bp3-icon-download::before{
	  content:""; }
	
	.bp3-icon-drag-handle-horizontal::before{
	  content:""; }
	
	.bp3-icon-drag-handle-vertical::before{
	  content:""; }
	
	.bp3-icon-draw::before{
	  content:""; }
	
	.bp3-icon-drive-time::before{
	  content:""; }
	
	.bp3-icon-duplicate::before{
	  content:""; }
	
	.bp3-icon-edit::before{
	  content:"✎"; }
	
	.bp3-icon-eject::before{
	  content:"⏏"; }
	
	.bp3-icon-endorsed::before{
	  content:""; }
	
	.bp3-icon-envelope::before{
	  content:"✉"; }
	
	.bp3-icon-equals::before{
	  content:""; }
	
	.bp3-icon-eraser::before{
	  content:""; }
	
	.bp3-icon-error::before{
	  content:""; }
	
	.bp3-icon-euro::before{
	  content:"€"; }
	
	.bp3-icon-exchange::before{
	  content:""; }
	
	.bp3-icon-exclude-row::before{
	  content:""; }
	
	.bp3-icon-expand-all::before{
	  content:""; }
	
	.bp3-icon-export::before{
	  content:""; }
	
	.bp3-icon-eye-off::before{
	  content:""; }
	
	.bp3-icon-eye-on::before{
	  content:""; }
	
	.bp3-icon-eye-open::before{
	  content:""; }
	
	.bp3-icon-fast-backward::before{
	  content:""; }
	
	.bp3-icon-fast-forward::before{
	  content:""; }
	
	.bp3-icon-feed::before{
	  content:""; }
	
	.bp3-icon-feed-subscribed::before{
	  content:""; }
	
	.bp3-icon-film::before{
	  content:""; }
	
	.bp3-icon-filter::before{
	  content:""; }
	
	.bp3-icon-filter-keep::before{
	  content:""; }
	
	.bp3-icon-filter-list::before{
	  content:""; }
	
	.bp3-icon-filter-open::before{
	  content:""; }
	
	.bp3-icon-filter-remove::before{
	  content:""; }
	
	.bp3-icon-flag::before{
	  content:"⚑"; }
	
	.bp3-icon-flame::before{
	  content:""; }
	
	.bp3-icon-flash::before{
	  content:""; }
	
	.bp3-icon-floppy-disk::before{
	  content:""; }
	
	.bp3-icon-flow-branch::before{
	  content:""; }
	
	.bp3-icon-flow-end::before{
	  content:""; }
	
	.bp3-icon-flow-linear::before{
	  content:""; }
	
	.bp3-icon-flow-review::before{
	  content:""; }
	
	.bp3-icon-flow-review-branch::before{
	  content:""; }
	
	.bp3-icon-flows::before{
	  content:""; }
	
	.bp3-icon-folder-close::before{
	  content:""; }
	
	.bp3-icon-folder-new::before{
	  content:""; }
	
	.bp3-icon-folder-open::before{
	  content:""; }
	
	.bp3-icon-folder-shared::before{
	  content:""; }
	
	.bp3-icon-folder-shared-open::before{
	  content:""; }
	
	.bp3-icon-follower::before{
	  content:""; }
	
	.bp3-icon-following::before{
	  content:""; }
	
	.bp3-icon-font::before{
	  content:""; }
	
	.bp3-icon-fork::before{
	  content:""; }
	
	.bp3-icon-form::before{
	  content:""; }
	
	.bp3-icon-full-circle::before{
	  content:""; }
	
	.bp3-icon-full-stacked-chart::before{
	  content:""; }
	
	.bp3-icon-fullscreen::before{
	  content:""; }
	
	.bp3-icon-function::before{
	  content:""; }
	
	.bp3-icon-gantt-chart::before{
	  content:""; }
	
	.bp3-icon-geolocation::before{
	  content:""; }
	
	.bp3-icon-geosearch::before{
	  content:""; }
	
	.bp3-icon-git-branch::before{
	  content:""; }
	
	.bp3-icon-git-commit::before{
	  content:""; }
	
	.bp3-icon-git-merge::before{
	  content:""; }
	
	.bp3-icon-git-new-branch::before{
	  content:""; }
	
	.bp3-icon-git-pull::before{
	  content:""; }
	
	.bp3-icon-git-push::before{
	  content:""; }
	
	.bp3-icon-git-repo::before{
	  content:""; }
	
	.bp3-icon-glass::before{
	  content:""; }
	
	.bp3-icon-globe::before{
	  content:""; }
	
	.bp3-icon-globe-network::before{
	  content:""; }
	
	.bp3-icon-graph::before{
	  content:""; }
	
	.bp3-icon-graph-remove::before{
	  content:""; }
	
	.bp3-icon-greater-than::before{
	  content:""; }
	
	.bp3-icon-greater-than-or-equal-to::before{
	  content:""; }
	
	.bp3-icon-grid::before{
	  content:""; }
	
	.bp3-icon-grid-view::before{
	  content:""; }
	
	.bp3-icon-group-objects::before{
	  content:""; }
	
	.bp3-icon-grouped-bar-chart::before{
	  content:""; }
	
	.bp3-icon-hand::before{
	  content:""; }
	
	.bp3-icon-hand-down::before{
	  content:""; }
	
	.bp3-icon-hand-left::before{
	  content:""; }
	
	.bp3-icon-hand-right::before{
	  content:""; }
	
	.bp3-icon-hand-up::before{
	  content:""; }
	
	.bp3-icon-header::before{
	  content:""; }
	
	.bp3-icon-header-one::before{
	  content:""; }
	
	.bp3-icon-header-two::before{
	  content:""; }
	
	.bp3-icon-headset::before{
	  content:""; }
	
	.bp3-icon-heart::before{
	  content:"♥"; }
	
	.bp3-icon-heart-broken::before{
	  content:""; }
	
	.bp3-icon-heat-grid::before{
	  content:""; }
	
	.bp3-icon-heatmap::before{
	  content:""; }
	
	.bp3-icon-help::before{
	  content:"?"; }
	
	.bp3-icon-helper-management::before{
	  content:""; }
	
	.bp3-icon-highlight::before{
	  content:""; }
	
	.bp3-icon-history::before{
	  content:""; }
	
	.bp3-icon-home::before{
	  content:"⌂"; }
	
	.bp3-icon-horizontal-bar-chart::before{
	  content:""; }
	
	.bp3-icon-horizontal-bar-chart-asc::before{
	  content:""; }
	
	.bp3-icon-horizontal-bar-chart-desc::before{
	  content:""; }
	
	.bp3-icon-horizontal-distribution::before{
	  content:""; }
	
	.bp3-icon-id-number::before{
	  content:""; }
	
	.bp3-icon-image-rotate-left::before{
	  content:""; }
	
	.bp3-icon-image-rotate-right::before{
	  content:""; }
	
	.bp3-icon-import::before{
	  content:""; }
	
	.bp3-icon-inbox::before{
	  content:""; }
	
	.bp3-icon-inbox-filtered::before{
	  content:""; }
	
	.bp3-icon-inbox-geo::before{
	  content:""; }
	
	.bp3-icon-inbox-search::before{
	  content:""; }
	
	.bp3-icon-inbox-update::before{
	  content:""; }
	
	.bp3-icon-info-sign::before{
	  content:"ℹ"; }
	
	.bp3-icon-inheritance::before{
	  content:""; }
	
	.bp3-icon-inner-join::before{
	  content:""; }
	
	.bp3-icon-insert::before{
	  content:""; }
	
	.bp3-icon-intersection::before{
	  content:""; }
	
	.bp3-icon-ip-address::before{
	  content:""; }
	
	.bp3-icon-issue::before{
	  content:""; }
	
	.bp3-icon-issue-closed::before{
	  content:""; }
	
	.bp3-icon-issue-new::before{
	  content:""; }
	
	.bp3-icon-italic::before{
	  content:""; }
	
	.bp3-icon-join-table::before{
	  content:""; }
	
	.bp3-icon-key::before{
	  content:""; }
	
	.bp3-icon-key-backspace::before{
	  content:""; }
	
	.bp3-icon-key-command::before{
	  content:""; }
	
	.bp3-icon-key-control::before{
	  content:""; }
	
	.bp3-icon-key-delete::before{
	  content:""; }
	
	.bp3-icon-key-enter::before{
	  content:""; }
	
	.bp3-icon-key-escape::before{
	  content:""; }
	
	.bp3-icon-key-option::before{
	  content:""; }
	
	.bp3-icon-key-shift::before{
	  content:""; }
	
	.bp3-icon-key-tab::before{
	  content:""; }
	
	.bp3-icon-known-vehicle::before{
	  content:""; }
	
	.bp3-icon-lab-test::before{
	  content:""; }
	
	.bp3-icon-label::before{
	  content:""; }
	
	.bp3-icon-layer::before{
	  content:""; }
	
	.bp3-icon-layers::before{
	  content:""; }
	
	.bp3-icon-layout::before{
	  content:""; }
	
	.bp3-icon-layout-auto::before{
	  content:""; }
	
	.bp3-icon-layout-balloon::before{
	  content:""; }
	
	.bp3-icon-layout-circle::before{
	  content:""; }
	
	.bp3-icon-layout-grid::before{
	  content:""; }
	
	.bp3-icon-layout-group-by::before{
	  content:""; }
	
	.bp3-icon-layout-hierarchy::before{
	  content:""; }
	
	.bp3-icon-layout-linear::before{
	  content:""; }
	
	.bp3-icon-layout-skew-grid::before{
	  content:""; }
	
	.bp3-icon-layout-sorted-clusters::before{
	  content:""; }
	
	.bp3-icon-learning::before{
	  content:""; }
	
	.bp3-icon-left-join::before{
	  content:""; }
	
	.bp3-icon-less-than::before{
	  content:""; }
	
	.bp3-icon-less-than-or-equal-to::before{
	  content:""; }
	
	.bp3-icon-lifesaver::before{
	  content:""; }
	
	.bp3-icon-lightbulb::before{
	  content:""; }
	
	.bp3-icon-link::before{
	  content:""; }
	
	.bp3-icon-list::before{
	  content:"☰"; }
	
	.bp3-icon-list-columns::before{
	  content:""; }
	
	.bp3-icon-list-detail-view::before{
	  content:""; }
	
	.bp3-icon-locate::before{
	  content:""; }
	
	.bp3-icon-lock::before{
	  content:""; }
	
	.bp3-icon-log-in::before{
	  content:""; }
	
	.bp3-icon-log-out::before{
	  content:""; }
	
	.bp3-icon-manual::before{
	  content:""; }
	
	.bp3-icon-manually-entered-data::before{
	  content:""; }
	
	.bp3-icon-map::before{
	  content:""; }
	
	.bp3-icon-map-create::before{
	  content:""; }
	
	.bp3-icon-map-marker::before{
	  content:""; }
	
	.bp3-icon-maximize::before{
	  content:""; }
	
	.bp3-icon-media::before{
	  content:""; }
	
	.bp3-icon-menu::before{
	  content:""; }
	
	.bp3-icon-menu-closed::before{
	  content:""; }
	
	.bp3-icon-menu-open::before{
	  content:""; }
	
	.bp3-icon-merge-columns::before{
	  content:""; }
	
	.bp3-icon-merge-links::before{
	  content:""; }
	
	.bp3-icon-minimize::before{
	  content:""; }
	
	.bp3-icon-minus::before{
	  content:"−"; }
	
	.bp3-icon-mobile-phone::before{
	  content:""; }
	
	.bp3-icon-mobile-video::before{
	  content:""; }
	
	.bp3-icon-moon::before{
	  content:""; }
	
	.bp3-icon-more::before{
	  content:""; }
	
	.bp3-icon-mountain::before{
	  content:""; }
	
	.bp3-icon-move::before{
	  content:""; }
	
	.bp3-icon-mugshot::before{
	  content:""; }
	
	.bp3-icon-multi-select::before{
	  content:""; }
	
	.bp3-icon-music::before{
	  content:""; }
	
	.bp3-icon-new-drawing::before{
	  content:""; }
	
	.bp3-icon-new-grid-item::before{
	  content:""; }
	
	.bp3-icon-new-layer::before{
	  content:""; }
	
	.bp3-icon-new-layers::before{
	  content:""; }
	
	.bp3-icon-new-link::before{
	  content:""; }
	
	.bp3-icon-new-object::before{
	  content:""; }
	
	.bp3-icon-new-person::before{
	  content:""; }
	
	.bp3-icon-new-prescription::before{
	  content:""; }
	
	.bp3-icon-new-text-box::before{
	  content:""; }
	
	.bp3-icon-ninja::before{
	  content:""; }
	
	.bp3-icon-not-equal-to::before{
	  content:""; }
	
	.bp3-icon-notifications::before{
	  content:""; }
	
	.bp3-icon-notifications-updated::before{
	  content:""; }
	
	.bp3-icon-numbered-list::before{
	  content:""; }
	
	.bp3-icon-numerical::before{
	  content:""; }
	
	.bp3-icon-office::before{
	  content:""; }
	
	.bp3-icon-offline::before{
	  content:""; }
	
	.bp3-icon-oil-field::before{
	  content:""; }
	
	.bp3-icon-one-column::before{
	  content:""; }
	
	.bp3-icon-outdated::before{
	  content:""; }
	
	.bp3-icon-page-layout::before{
	  content:""; }
	
	.bp3-icon-panel-stats::before{
	  content:""; }
	
	.bp3-icon-panel-table::before{
	  content:""; }
	
	.bp3-icon-paperclip::before{
	  content:""; }
	
	.bp3-icon-paragraph::before{
	  content:""; }
	
	.bp3-icon-path::before{
	  content:""; }
	
	.bp3-icon-path-search::before{
	  content:""; }
	
	.bp3-icon-pause::before{
	  content:""; }
	
	.bp3-icon-people::before{
	  content:""; }
	
	.bp3-icon-percentage::before{
	  content:""; }
	
	.bp3-icon-person::before{
	  content:""; }
	
	.bp3-icon-phone::before{
	  content:"☎"; }
	
	.bp3-icon-pie-chart::before{
	  content:""; }
	
	.bp3-icon-pin::before{
	  content:""; }
	
	.bp3-icon-pivot::before{
	  content:""; }
	
	.bp3-icon-pivot-table::before{
	  content:""; }
	
	.bp3-icon-play::before{
	  content:""; }
	
	.bp3-icon-plus::before{
	  content:"+"; }
	
	.bp3-icon-polygon-filter::before{
	  content:""; }
	
	.bp3-icon-power::before{
	  content:""; }
	
	.bp3-icon-predictive-analysis::before{
	  content:""; }
	
	.bp3-icon-prescription::before{
	  content:""; }
	
	.bp3-icon-presentation::before{
	  content:""; }
	
	.bp3-icon-print::before{
	  content:"⎙"; }
	
	.bp3-icon-projects::before{
	  content:""; }
	
	.bp3-icon-properties::before{
	  content:""; }
	
	.bp3-icon-property::before{
	  content:""; }
	
	.bp3-icon-publish-function::before{
	  content:""; }
	
	.bp3-icon-pulse::before{
	  content:""; }
	
	.bp3-icon-random::before{
	  content:""; }
	
	.bp3-icon-record::before{
	  content:""; }
	
	.bp3-icon-redo::before{
	  content:""; }
	
	.bp3-icon-refresh::before{
	  content:""; }
	
	.bp3-icon-regression-chart::before{
	  content:""; }
	
	.bp3-icon-remove::before{
	  content:""; }
	
	.bp3-icon-remove-column::before{
	  content:""; }
	
	.bp3-icon-remove-column-left::before{
	  content:""; }
	
	.bp3-icon-remove-column-right::before{
	  content:""; }
	
	.bp3-icon-remove-row-bottom::before{
	  content:""; }
	
	.bp3-icon-remove-row-top::before{
	  content:""; }
	
	.bp3-icon-repeat::before{
	  content:""; }
	
	.bp3-icon-reset::before{
	  content:""; }
	
	.bp3-icon-resolve::before{
	  content:""; }
	
	.bp3-icon-rig::before{
	  content:""; }
	
	.bp3-icon-right-join::before{
	  content:""; }
	
	.bp3-icon-ring::before{
	  content:""; }
	
	.bp3-icon-rotate-document::before{
	  content:""; }
	
	.bp3-icon-rotate-page::before{
	  content:""; }
	
	.bp3-icon-satellite::before{
	  content:""; }
	
	.bp3-icon-saved::before{
	  content:""; }
	
	.bp3-icon-scatter-plot::before{
	  content:""; }
	
	.bp3-icon-search::before{
	  content:""; }
	
	.bp3-icon-search-around::before{
	  content:""; }
	
	.bp3-icon-search-template::before{
	  content:""; }
	
	.bp3-icon-search-text::before{
	  content:""; }
	
	.bp3-icon-segmented-control::before{
	  content:""; }
	
	.bp3-icon-select::before{
	  content:""; }
	
	.bp3-icon-selection::before{
	  content:"⦿"; }
	
	.bp3-icon-send-to::before{
	  content:""; }
	
	.bp3-icon-send-to-graph::before{
	  content:""; }
	
	.bp3-icon-send-to-map::before{
	  content:""; }
	
	.bp3-icon-series-add::before{
	  content:""; }
	
	.bp3-icon-series-configuration::before{
	  content:""; }
	
	.bp3-icon-series-derived::before{
	  content:""; }
	
	.bp3-icon-series-filtered::before{
	  content:""; }
	
	.bp3-icon-series-search::before{
	  content:""; }
	
	.bp3-icon-settings::before{
	  content:""; }
	
	.bp3-icon-share::before{
	  content:""; }
	
	.bp3-icon-shield::before{
	  content:""; }
	
	.bp3-icon-shop::before{
	  content:""; }
	
	.bp3-icon-shopping-cart::before{
	  content:""; }
	
	.bp3-icon-signal-search::before{
	  content:""; }
	
	.bp3-icon-sim-card::before{
	  content:""; }
	
	.bp3-icon-slash::before{
	  content:""; }
	
	.bp3-icon-small-cross::before{
	  content:""; }
	
	.bp3-icon-small-minus::before{
	  content:""; }
	
	.bp3-icon-small-plus::before{
	  content:""; }
	
	.bp3-icon-small-tick::before{
	  content:""; }
	
	.bp3-icon-snowflake::before{
	  content:""; }
	
	.bp3-icon-social-media::before{
	  content:""; }
	
	.bp3-icon-sort::before{
	  content:""; }
	
	.bp3-icon-sort-alphabetical::before{
	  content:""; }
	
	.bp3-icon-sort-alphabetical-desc::before{
	  content:""; }
	
	.bp3-icon-sort-asc::before{
	  content:""; }
	
	.bp3-icon-sort-desc::before{
	  content:""; }
	
	.bp3-icon-sort-numerical::before{
	  content:""; }
	
	.bp3-icon-sort-numerical-desc::before{
	  content:""; }
	
	.bp3-icon-split-columns::before{
	  content:""; }
	
	.bp3-icon-square::before{
	  content:""; }
	
	.bp3-icon-stacked-chart::before{
	  content:""; }
	
	.bp3-icon-star::before{
	  content:"★"; }
	
	.bp3-icon-star-empty::before{
	  content:"☆"; }
	
	.bp3-icon-step-backward::before{
	  content:""; }
	
	.bp3-icon-step-chart::before{
	  content:""; }
	
	.bp3-icon-step-forward::before{
	  content:""; }
	
	.bp3-icon-stop::before{
	  content:""; }
	
	.bp3-icon-stopwatch::before{
	  content:""; }
	
	.bp3-icon-strikethrough::before{
	  content:""; }
	
	.bp3-icon-style::before{
	  content:""; }
	
	.bp3-icon-swap-horizontal::before{
	  content:""; }
	
	.bp3-icon-swap-vertical::before{
	  content:""; }
	
	.bp3-icon-symbol-circle::before{
	  content:""; }
	
	.bp3-icon-symbol-cross::before{
	  content:""; }
	
	.bp3-icon-symbol-diamond::before{
	  content:""; }
	
	.bp3-icon-symbol-square::before{
	  content:""; }
	
	.bp3-icon-symbol-triangle-down::before{
	  content:""; }
	
	.bp3-icon-symbol-triangle-up::before{
	  content:""; }
	
	.bp3-icon-tag::before{
	  content:""; }
	
	.bp3-icon-take-action::before{
	  content:""; }
	
	.bp3-icon-taxi::before{
	  content:""; }
	
	.bp3-icon-text-highlight::before{
	  content:""; }
	
	.bp3-icon-th::before{
	  content:""; }
	
	.bp3-icon-th-derived::before{
	  content:""; }
	
	.bp3-icon-th-disconnect::before{
	  content:""; }
	
	.bp3-icon-th-filtered::before{
	  content:""; }
	
	.bp3-icon-th-list::before{
	  content:""; }
	
	.bp3-icon-thumbs-down::before{
	  content:""; }
	
	.bp3-icon-thumbs-up::before{
	  content:""; }
	
	.bp3-icon-tick::before{
	  content:"✓"; }
	
	.bp3-icon-tick-circle::before{
	  content:""; }
	
	.bp3-icon-time::before{
	  content:"⏲"; }
	
	.bp3-icon-timeline-area-chart::before{
	  content:""; }
	
	.bp3-icon-timeline-bar-chart::before{
	  content:""; }
	
	.bp3-icon-timeline-events::before{
	  content:""; }
	
	.bp3-icon-timeline-line-chart::before{
	  content:""; }
	
	.bp3-icon-tint::before{
	  content:""; }
	
	.bp3-icon-torch::before{
	  content:""; }
	
	.bp3-icon-tractor::before{
	  content:""; }
	
	.bp3-icon-train::before{
	  content:""; }
	
	.bp3-icon-translate::before{
	  content:""; }
	
	.bp3-icon-trash::before{
	  content:""; }
	
	.bp3-icon-tree::before{
	  content:""; }
	
	.bp3-icon-trending-down::before{
	  content:""; }
	
	.bp3-icon-trending-up::before{
	  content:""; }
	
	.bp3-icon-truck::before{
	  content:""; }
	
	.bp3-icon-two-columns::before{
	  content:""; }
	
	.bp3-icon-unarchive::before{
	  content:""; }
	
	.bp3-icon-underline::before{
	  content:"⎁"; }
	
	.bp3-icon-undo::before{
	  content:"⎌"; }
	
	.bp3-icon-ungroup-objects::before{
	  content:""; }
	
	.bp3-icon-unknown-vehicle::before{
	  content:""; }
	
	.bp3-icon-unlock::before{
	  content:""; }
	
	.bp3-icon-unpin::before{
	  content:""; }
	
	.bp3-icon-unresolve::before{
	  content:""; }
	
	.bp3-icon-updated::before{
	  content:""; }
	
	.bp3-icon-upload::before{
	  content:""; }
	
	.bp3-icon-user::before{
	  content:""; }
	
	.bp3-icon-variable::before{
	  content:""; }
	
	.bp3-icon-vertical-bar-chart-asc::before{
	  content:""; }
	
	.bp3-icon-vertical-bar-chart-desc::before{
	  content:""; }
	
	.bp3-icon-vertical-distribution::before{
	  content:""; }
	
	.bp3-icon-video::before{
	  content:""; }
	
	.bp3-icon-volume-down::before{
	  content:""; }
	
	.bp3-icon-volume-off::before{
	  content:""; }
	
	.bp3-icon-volume-up::before{
	  content:""; }
	
	.bp3-icon-walk::before{
	  content:""; }
	
	.bp3-icon-warning-sign::before{
	  content:""; }
	
	.bp3-icon-waterfall-chart::before{
	  content:""; }
	
	.bp3-icon-widget::before{
	  content:""; }
	
	.bp3-icon-widget-button::before{
	  content:""; }
	
	.bp3-icon-widget-footer::before{
	  content:""; }
	
	.bp3-icon-widget-header::before{
	  content:""; }
	
	.bp3-icon-wrench::before{
	  content:""; }
	
	.bp3-icon-zoom-in::before{
	  content:""; }
	
	.bp3-icon-zoom-out::before{
	  content:""; }
	
	.bp3-icon-zoom-to-fit::before{
	  content:""; }
	.bp3-submenu > .bp3-popover-wrapper{
	  display:block; }
	
	.bp3-submenu .bp3-popover-target{
	  display:block; }
	  .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-menu-item{ }
	
	.bp3-submenu.bp3-popover{
	  -webkit-box-shadow:none;
			  box-shadow:none;
	  padding:0 5px; }
	  .bp3-submenu.bp3-popover > .bp3-popover-content{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-submenu.bp3-popover, .bp3-submenu.bp3-popover.bp3-dark{
		-webkit-box-shadow:none;
				box-shadow:none; }
		.bp3-dark .bp3-submenu.bp3-popover > .bp3-popover-content, .bp3-submenu.bp3-popover.bp3-dark > .bp3-popover-content{
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
	.bp3-menu{
	  background:#ffffff;
	  border-radius:3px;
	  color:#182026;
	  list-style:none;
	  margin:0;
	  min-width:180px;
	  padding:5px;
	  text-align:left; }
	
	.bp3-menu-divider{
	  border-top:1px solid rgba(16, 22, 26, 0.15);
	  display:block;
	  margin:5px; }
	  .bp3-dark .bp3-menu-divider{
		border-top-color:rgba(255, 255, 255, 0.15); }
	
	.bp3-menu-item{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:start;
		  -ms-flex-align:start;
			  align-items:flex-start;
	  border-radius:2px;
	  color:inherit;
	  line-height:20px;
	  padding:5px 7px;
	  text-decoration:none;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-menu-item > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-menu-item > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-menu-item::before,
	  .bp3-menu-item > *{
		margin-right:7px; }
	  .bp3-menu-item:empty::before,
	  .bp3-menu-item > :last-child{
		margin-right:0; }
	  .bp3-menu-item > .bp3-fill{
		word-break:break-word; }
	  .bp3-menu-item:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-menu-item{
		background-color:rgba(167, 182, 194, 0.3);
		cursor:pointer;
		text-decoration:none; }
	  .bp3-menu-item.bp3-disabled{
		background-color:inherit;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed; }
	  .bp3-dark .bp3-menu-item{
		color:inherit; }
		.bp3-dark .bp3-menu-item:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-menu-item{
		  background-color:rgba(138, 155, 168, 0.15);
		  color:inherit; }
		.bp3-dark .bp3-menu-item.bp3-disabled{
		  background-color:inherit;
		  color:rgba(167, 182, 194, 0.6); }
	  .bp3-menu-item.bp3-intent-primary{
		color:#106ba3; }
		.bp3-menu-item.bp3-intent-primary .bp3-icon{
		  color:inherit; }
		.bp3-menu-item.bp3-intent-primary::before, .bp3-menu-item.bp3-intent-primary::after,
		.bp3-menu-item.bp3-intent-primary .bp3-menu-item-label{
		  color:#106ba3; }
		.bp3-menu-item.bp3-intent-primary:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-menu-item.bp3-intent-primary.bp3-active{
		  background-color:#137cbd; }
		.bp3-menu-item.bp3-intent-primary:active{
		  background-color:#106ba3; }
		.bp3-menu-item.bp3-intent-primary:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-menu-item.bp3-intent-primary:hover::before, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::before, .bp3-menu-item.bp3-intent-primary:hover::after, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::after,
		.bp3-menu-item.bp3-intent-primary:hover .bp3-menu-item-label,
		.bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item .bp3-menu-item-label, .bp3-menu-item.bp3-intent-primary:active, .bp3-menu-item.bp3-intent-primary:active::before, .bp3-menu-item.bp3-intent-primary:active::after,
		.bp3-menu-item.bp3-intent-primary:active .bp3-menu-item-label, .bp3-menu-item.bp3-intent-primary.bp3-active, .bp3-menu-item.bp3-intent-primary.bp3-active::before, .bp3-menu-item.bp3-intent-primary.bp3-active::after,
		.bp3-menu-item.bp3-intent-primary.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-menu-item.bp3-intent-success{
		color:#0d8050; }
		.bp3-menu-item.bp3-intent-success .bp3-icon{
		  color:inherit; }
		.bp3-menu-item.bp3-intent-success::before, .bp3-menu-item.bp3-intent-success::after,
		.bp3-menu-item.bp3-intent-success .bp3-menu-item-label{
		  color:#0d8050; }
		.bp3-menu-item.bp3-intent-success:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-menu-item.bp3-intent-success.bp3-active{
		  background-color:#0f9960; }
		.bp3-menu-item.bp3-intent-success:active{
		  background-color:#0d8050; }
		.bp3-menu-item.bp3-intent-success:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-menu-item.bp3-intent-success:hover::before, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::before, .bp3-menu-item.bp3-intent-success:hover::after, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::after,
		.bp3-menu-item.bp3-intent-success:hover .bp3-menu-item-label,
		.bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item .bp3-menu-item-label, .bp3-menu-item.bp3-intent-success:active, .bp3-menu-item.bp3-intent-success:active::before, .bp3-menu-item.bp3-intent-success:active::after,
		.bp3-menu-item.bp3-intent-success:active .bp3-menu-item-label, .bp3-menu-item.bp3-intent-success.bp3-active, .bp3-menu-item.bp3-intent-success.bp3-active::before, .bp3-menu-item.bp3-intent-success.bp3-active::after,
		.bp3-menu-item.bp3-intent-success.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-menu-item.bp3-intent-warning{
		color:#bf7326; }
		.bp3-menu-item.bp3-intent-warning .bp3-icon{
		  color:inherit; }
		.bp3-menu-item.bp3-intent-warning::before, .bp3-menu-item.bp3-intent-warning::after,
		.bp3-menu-item.bp3-intent-warning .bp3-menu-item-label{
		  color:#bf7326; }
		.bp3-menu-item.bp3-intent-warning:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-menu-item.bp3-intent-warning.bp3-active{
		  background-color:#d9822b; }
		.bp3-menu-item.bp3-intent-warning:active{
		  background-color:#bf7326; }
		.bp3-menu-item.bp3-intent-warning:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-menu-item.bp3-intent-warning:hover::before, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::before, .bp3-menu-item.bp3-intent-warning:hover::after, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::after,
		.bp3-menu-item.bp3-intent-warning:hover .bp3-menu-item-label,
		.bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item .bp3-menu-item-label, .bp3-menu-item.bp3-intent-warning:active, .bp3-menu-item.bp3-intent-warning:active::before, .bp3-menu-item.bp3-intent-warning:active::after,
		.bp3-menu-item.bp3-intent-warning:active .bp3-menu-item-label, .bp3-menu-item.bp3-intent-warning.bp3-active, .bp3-menu-item.bp3-intent-warning.bp3-active::before, .bp3-menu-item.bp3-intent-warning.bp3-active::after,
		.bp3-menu-item.bp3-intent-warning.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-menu-item.bp3-intent-danger{
		color:#c23030; }
		.bp3-menu-item.bp3-intent-danger .bp3-icon{
		  color:inherit; }
		.bp3-menu-item.bp3-intent-danger::before, .bp3-menu-item.bp3-intent-danger::after,
		.bp3-menu-item.bp3-intent-danger .bp3-menu-item-label{
		  color:#c23030; }
		.bp3-menu-item.bp3-intent-danger:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-menu-item.bp3-intent-danger.bp3-active{
		  background-color:#db3737; }
		.bp3-menu-item.bp3-intent-danger:active{
		  background-color:#c23030; }
		.bp3-menu-item.bp3-intent-danger:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-menu-item.bp3-intent-danger:hover::before, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::before, .bp3-menu-item.bp3-intent-danger:hover::after, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::after,
		.bp3-menu-item.bp3-intent-danger:hover .bp3-menu-item-label,
		.bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item .bp3-menu-item-label, .bp3-menu-item.bp3-intent-danger:active, .bp3-menu-item.bp3-intent-danger:active::before, .bp3-menu-item.bp3-intent-danger:active::after,
		.bp3-menu-item.bp3-intent-danger:active .bp3-menu-item-label, .bp3-menu-item.bp3-intent-danger.bp3-active, .bp3-menu-item.bp3-intent-danger.bp3-active::before, .bp3-menu-item.bp3-intent-danger.bp3-active::after,
		.bp3-menu-item.bp3-intent-danger.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-menu-item::before{
		font-family:"Icons16", sans-serif;
		font-size:16px;
		font-style:normal;
		font-weight:400;
		line-height:1;
		-moz-osx-font-smoothing:grayscale;
		-webkit-font-smoothing:antialiased;
		margin-right:7px; }
	  .bp3-menu-item::before,
	  .bp3-menu-item > .bp3-icon{
		color:#5c7080;
		margin-top:2px; }
	  .bp3-menu-item .bp3-menu-item-label{
		color:#5c7080; }
	  .bp3-menu-item:hover, .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-menu-item{
		color:inherit; }
	  .bp3-menu-item.bp3-active, .bp3-menu-item:active{
		background-color:rgba(115, 134, 148, 0.3); }
	  .bp3-menu-item.bp3-disabled{
		background-color:inherit !important;
		color:rgba(92, 112, 128, 0.6) !important;
		cursor:not-allowed !important;
		outline:none !important; }
		.bp3-menu-item.bp3-disabled::before,
		.bp3-menu-item.bp3-disabled > .bp3-icon,
		.bp3-menu-item.bp3-disabled .bp3-menu-item-label{
		  color:rgba(92, 112, 128, 0.6) !important; }
	  .bp3-large .bp3-menu-item{
		font-size:16px;
		line-height:22px;
		padding:9px 7px; }
		.bp3-large .bp3-menu-item .bp3-icon{
		  margin-top:3px; }
		.bp3-large .bp3-menu-item::before{
		  font-family:"Icons20", sans-serif;
		  font-size:20px;
		  font-style:normal;
		  font-weight:400;
		  line-height:1;
		  -moz-osx-font-smoothing:grayscale;
		  -webkit-font-smoothing:antialiased;
		  margin-right:10px;
		  margin-top:1px; }
	
	button.bp3-menu-item{
	  background:none;
	  border:none;
	  text-align:left;
	  width:100%; }
	.bp3-menu-header{
	  border-top:1px solid rgba(16, 22, 26, 0.15);
	  display:block;
	  margin:5px;
	  cursor:default;
	  padding-left:2px; }
	  .bp3-dark .bp3-menu-header{
		border-top-color:rgba(255, 255, 255, 0.15); }
	  .bp3-menu-header:first-of-type{
		border-top:none; }
	  .bp3-menu-header > h6{
		color:#182026;
		font-weight:600;
		overflow:hidden;
		text-overflow:ellipsis;
		white-space:nowrap;
		word-wrap:normal;
		line-height:17px;
		margin:0;
		padding:10px 7px 0 1px; }
		.bp3-dark .bp3-menu-header > h6{
		  color:#f5f8fa; }
	  .bp3-menu-header:first-of-type > h6{
		padding-top:0; }
	  .bp3-large .bp3-menu-header > h6{
		font-size:18px;
		padding-bottom:5px;
		padding-top:15px; }
	  .bp3-large .bp3-menu-header:first-of-type > h6{
		padding-top:0; }
	
	.bp3-dark .bp3-menu{
	  background:#30404d;
	  color:#f5f8fa; }
	
	.bp3-dark .bp3-menu-item{ }
	  .bp3-dark .bp3-menu-item.bp3-intent-primary{
		color:#48aff0; }
		.bp3-dark .bp3-menu-item.bp3-intent-primary .bp3-icon{
		  color:inherit; }
		.bp3-dark .bp3-menu-item.bp3-intent-primary::before, .bp3-dark .bp3-menu-item.bp3-intent-primary::after,
		.bp3-dark .bp3-menu-item.bp3-intent-primary .bp3-menu-item-label{
		  color:#48aff0; }
		.bp3-dark .bp3-menu-item.bp3-intent-primary:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-primary.bp3-active{
		  background-color:#137cbd; }
		.bp3-dark .bp3-menu-item.bp3-intent-primary:active{
		  background-color:#106ba3; }
		.bp3-dark .bp3-menu-item.bp3-intent-primary:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-primary:hover::before, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::before, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::before, .bp3-dark .bp3-menu-item.bp3-intent-primary:hover::after, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::after, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item::after,
		.bp3-dark .bp3-menu-item.bp3-intent-primary:hover .bp3-menu-item-label,
		.bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item .bp3-menu-item-label,
		.bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-primary.bp3-menu-item .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-primary:active, .bp3-dark .bp3-menu-item.bp3-intent-primary:active::before, .bp3-dark .bp3-menu-item.bp3-intent-primary:active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-primary:active .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-primary.bp3-active, .bp3-dark .bp3-menu-item.bp3-intent-primary.bp3-active::before, .bp3-dark .bp3-menu-item.bp3-intent-primary.bp3-active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-primary.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-dark .bp3-menu-item.bp3-intent-success{
		color:#3dcc91; }
		.bp3-dark .bp3-menu-item.bp3-intent-success .bp3-icon{
		  color:inherit; }
		.bp3-dark .bp3-menu-item.bp3-intent-success::before, .bp3-dark .bp3-menu-item.bp3-intent-success::after,
		.bp3-dark .bp3-menu-item.bp3-intent-success .bp3-menu-item-label{
		  color:#3dcc91; }
		.bp3-dark .bp3-menu-item.bp3-intent-success:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-success.bp3-active{
		  background-color:#0f9960; }
		.bp3-dark .bp3-menu-item.bp3-intent-success:active{
		  background-color:#0d8050; }
		.bp3-dark .bp3-menu-item.bp3-intent-success:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-success:hover::before, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::before, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::before, .bp3-dark .bp3-menu-item.bp3-intent-success:hover::after, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::after, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item::after,
		.bp3-dark .bp3-menu-item.bp3-intent-success:hover .bp3-menu-item-label,
		.bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item .bp3-menu-item-label,
		.bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-success.bp3-menu-item .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-success:active, .bp3-dark .bp3-menu-item.bp3-intent-success:active::before, .bp3-dark .bp3-menu-item.bp3-intent-success:active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-success:active .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-success.bp3-active, .bp3-dark .bp3-menu-item.bp3-intent-success.bp3-active::before, .bp3-dark .bp3-menu-item.bp3-intent-success.bp3-active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-success.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-dark .bp3-menu-item.bp3-intent-warning{
		color:#ffb366; }
		.bp3-dark .bp3-menu-item.bp3-intent-warning .bp3-icon{
		  color:inherit; }
		.bp3-dark .bp3-menu-item.bp3-intent-warning::before, .bp3-dark .bp3-menu-item.bp3-intent-warning::after,
		.bp3-dark .bp3-menu-item.bp3-intent-warning .bp3-menu-item-label{
		  color:#ffb366; }
		.bp3-dark .bp3-menu-item.bp3-intent-warning:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-warning.bp3-active{
		  background-color:#d9822b; }
		.bp3-dark .bp3-menu-item.bp3-intent-warning:active{
		  background-color:#bf7326; }
		.bp3-dark .bp3-menu-item.bp3-intent-warning:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-warning:hover::before, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::before, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::before, .bp3-dark .bp3-menu-item.bp3-intent-warning:hover::after, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::after, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item::after,
		.bp3-dark .bp3-menu-item.bp3-intent-warning:hover .bp3-menu-item-label,
		.bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item .bp3-menu-item-label,
		.bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-warning.bp3-menu-item .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-warning:active, .bp3-dark .bp3-menu-item.bp3-intent-warning:active::before, .bp3-dark .bp3-menu-item.bp3-intent-warning:active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-warning:active .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-warning.bp3-active, .bp3-dark .bp3-menu-item.bp3-intent-warning.bp3-active::before, .bp3-dark .bp3-menu-item.bp3-intent-warning.bp3-active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-warning.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-dark .bp3-menu-item.bp3-intent-danger{
		color:#ff7373; }
		.bp3-dark .bp3-menu-item.bp3-intent-danger .bp3-icon{
		  color:inherit; }
		.bp3-dark .bp3-menu-item.bp3-intent-danger::before, .bp3-dark .bp3-menu-item.bp3-intent-danger::after,
		.bp3-dark .bp3-menu-item.bp3-intent-danger .bp3-menu-item-label{
		  color:#ff7373; }
		.bp3-dark .bp3-menu-item.bp3-intent-danger:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-danger.bp3-active{
		  background-color:#db3737; }
		.bp3-dark .bp3-menu-item.bp3-intent-danger:active{
		  background-color:#c23030; }
		.bp3-dark .bp3-menu-item.bp3-intent-danger:hover, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item, .bp3-dark .bp3-menu-item.bp3-intent-danger:hover::before, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::before, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::before, .bp3-dark .bp3-menu-item.bp3-intent-danger:hover::after, .bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::after, .bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item::after,
		.bp3-dark .bp3-menu-item.bp3-intent-danger:hover .bp3-menu-item-label,
		.bp3-dark .bp3-submenu .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item .bp3-menu-item-label,
		.bp3-submenu .bp3-dark .bp3-popover-target.bp3-popover-open > .bp3-intent-danger.bp3-menu-item .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-danger:active, .bp3-dark .bp3-menu-item.bp3-intent-danger:active::before, .bp3-dark .bp3-menu-item.bp3-intent-danger:active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-danger:active .bp3-menu-item-label, .bp3-dark .bp3-menu-item.bp3-intent-danger.bp3-active, .bp3-dark .bp3-menu-item.bp3-intent-danger.bp3-active::before, .bp3-dark .bp3-menu-item.bp3-intent-danger.bp3-active::after,
		.bp3-dark .bp3-menu-item.bp3-intent-danger.bp3-active .bp3-menu-item-label{
		  color:#ffffff; }
	  .bp3-dark .bp3-menu-item::before,
	  .bp3-dark .bp3-menu-item > .bp3-icon{
		color:#a7b6c2; }
	  .bp3-dark .bp3-menu-item .bp3-menu-item-label{
		color:#a7b6c2; }
	  .bp3-dark .bp3-menu-item.bp3-active, .bp3-dark .bp3-menu-item:active{
		background-color:rgba(138, 155, 168, 0.3); }
	  .bp3-dark .bp3-menu-item.bp3-disabled{
		color:rgba(167, 182, 194, 0.6) !important; }
		.bp3-dark .bp3-menu-item.bp3-disabled::before,
		.bp3-dark .bp3-menu-item.bp3-disabled > .bp3-icon,
		.bp3-dark .bp3-menu-item.bp3-disabled .bp3-menu-item-label{
		  color:rgba(167, 182, 194, 0.6) !important; }
	
	.bp3-dark .bp3-menu-divider,
	.bp3-dark .bp3-menu-header{
	  border-color:rgba(255, 255, 255, 0.15); }
	
	.bp3-dark .bp3-menu-header > h6{
	  color:#f5f8fa; }
	
	.bp3-label .bp3-menu{
	  margin-top:5px; }
	.bp3-navbar{
	  background-color:#ffffff;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.2);
	  height:50px;
	  padding:0 15px;
	  position:relative;
	  width:100%;
	  z-index:10; }
	  .bp3-navbar.bp3-dark,
	  .bp3-dark .bp3-navbar{
		background-color:#394b59; }
	  .bp3-navbar.bp3-dark{
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-dark .bp3-navbar{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 0 0 rgba(16, 22, 26, 0), 0 1px 1px rgba(16, 22, 26, 0.4); }
	  .bp3-navbar.bp3-fixed-top{
		left:0;
		position:fixed;
		right:0;
		top:0; }
	
	.bp3-navbar-heading{
	  font-size:16px;
	  margin-right:15px; }
	
	.bp3-navbar-group{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  height:50px; }
	  .bp3-navbar-group.bp3-align-left{
		float:left; }
	  .bp3-navbar-group.bp3-align-right{
		float:right; }
	
	.bp3-navbar-divider{
	  border-left:1px solid rgba(16, 22, 26, 0.15);
	  height:20px;
	  margin:0 10px; }
	  .bp3-dark .bp3-navbar-divider{
		border-left-color:rgba(255, 255, 255, 0.15); }
	.bp3-non-ideal-state{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  height:100%;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  text-align:center;
	  width:100%; }
	  .bp3-non-ideal-state > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-non-ideal-state > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-non-ideal-state::before,
	  .bp3-non-ideal-state > *{
		margin-bottom:20px; }
	  .bp3-non-ideal-state:empty::before,
	  .bp3-non-ideal-state > :last-child{
		margin-bottom:0; }
	  .bp3-non-ideal-state > *{
		max-width:400px; }
	
	.bp3-non-ideal-state-visual{
	  color:rgba(92, 112, 128, 0.6);
	  font-size:60px; }
	  .bp3-dark .bp3-non-ideal-state-visual{
		color:rgba(167, 182, 194, 0.6); }
	
	.bp3-overflow-list{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -ms-flex-wrap:nowrap;
		  flex-wrap:nowrap;
	  min-width:0; }
	
	.bp3-overflow-list-spacer{
	  -ms-flex-negative:1;
		  flex-shrink:1;
	  width:1px; }
	
	body.bp3-overlay-open{
	  overflow:hidden; }
	
	.bp3-overlay{
	  bottom:0;
	  left:0;
	  position:static;
	  right:0;
	  top:0;
	  z-index:20; }
	  .bp3-overlay:not(.bp3-overlay-open){
		pointer-events:none; }
	  .bp3-overlay.bp3-overlay-container{
		overflow:hidden;
		position:fixed; }
		.bp3-overlay.bp3-overlay-container.bp3-overlay-inline{
		  position:absolute; }
	  .bp3-overlay.bp3-overlay-scroll-container{
		overflow:auto;
		position:fixed; }
		.bp3-overlay.bp3-overlay-scroll-container.bp3-overlay-inline{
		  position:absolute; }
	  .bp3-overlay.bp3-overlay-inline{
		display:inline;
		overflow:visible; }
	
	.bp3-overlay-content{
	  position:fixed;
	  z-index:20; }
	  .bp3-overlay-inline .bp3-overlay-content,
	  .bp3-overlay-scroll-container .bp3-overlay-content{
		position:absolute; }
	
	.bp3-overlay-backdrop{
	  bottom:0;
	  left:0;
	  position:fixed;
	  right:0;
	  top:0;
	  opacity:1;
	  background-color:rgba(16, 22, 26, 0.7);
	  overflow:auto;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none;
	  z-index:20; }
	  .bp3-overlay-backdrop.bp3-overlay-enter, .bp3-overlay-backdrop.bp3-overlay-appear{
		opacity:0; }
	  .bp3-overlay-backdrop.bp3-overlay-enter-active, .bp3-overlay-backdrop.bp3-overlay-appear-active{
		opacity:1;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:200ms;
				transition-duration:200ms;
		-webkit-transition-property:opacity;
		transition-property:opacity;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-overlay-backdrop.bp3-overlay-exit{
		opacity:1; }
	  .bp3-overlay-backdrop.bp3-overlay-exit-active{
		opacity:0;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:200ms;
				transition-duration:200ms;
		-webkit-transition-property:opacity;
		transition-property:opacity;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-overlay-backdrop:focus{
		outline:none; }
	  .bp3-overlay-inline .bp3-overlay-backdrop{
		position:absolute; }
	.bp3-panel-stack{
	  overflow:hidden;
	  position:relative; }
	
	.bp3-panel-stack-header{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  -webkit-box-shadow:0 1px rgba(16, 22, 26, 0.15);
			  box-shadow:0 1px rgba(16, 22, 26, 0.15);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -ms-flex-negative:0;
		  flex-shrink:0;
	  height:30px;
	  z-index:1; }
	  .bp3-dark .bp3-panel-stack-header{
		-webkit-box-shadow:0 1px rgba(255, 255, 255, 0.15);
				box-shadow:0 1px rgba(255, 255, 255, 0.15); }
	  .bp3-panel-stack-header > span{
		-webkit-box-align:stretch;
			-ms-flex-align:stretch;
				align-items:stretch;
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		-webkit-box-flex:1;
			-ms-flex:1;
				flex:1; }
	  .bp3-panel-stack-header .bp3-heading{
		margin:0 5px; }
	
	.bp3-button.bp3-panel-stack-header-back{
	  margin-left:5px;
	  padding-left:0;
	  white-space:nowrap; }
	  .bp3-button.bp3-panel-stack-header-back .bp3-icon{
		margin:0 2px; }
	
	.bp3-panel-stack-view{
	  bottom:0;
	  left:0;
	  position:absolute;
	  right:0;
	  top:0;
	  background-color:#ffffff;
	  border-right:1px solid rgba(16, 22, 26, 0.15);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  margin-right:-1px;
	  overflow-y:auto;
	  z-index:1; }
	  .bp3-dark .bp3-panel-stack-view{
		background-color:#30404d; }
	  .bp3-panel-stack-view:nth-last-child(n + 4){
		display:none; }
	
	.bp3-panel-stack-push .bp3-panel-stack-enter, .bp3-panel-stack-push .bp3-panel-stack-appear{
	  -webkit-transform:translateX(100%);
			  transform:translateX(100%);
	  opacity:0; }
	
	.bp3-panel-stack-push .bp3-panel-stack-enter-active, .bp3-panel-stack-push .bp3-panel-stack-appear-active{
	  -webkit-transform:translate(0%);
			  transform:translate(0%);
	  opacity:1;
	  -webkit-transition-delay:0;
			  transition-delay:0;
	  -webkit-transition-duration:400ms;
			  transition-duration:400ms;
	  -webkit-transition-property:opacity, -webkit-transform;
	  transition-property:opacity, -webkit-transform;
	  transition-property:transform, opacity;
	  transition-property:transform, opacity, -webkit-transform;
	  -webkit-transition-timing-function:ease;
			  transition-timing-function:ease; }
	
	.bp3-panel-stack-push .bp3-panel-stack-exit{
	  -webkit-transform:translate(0%);
			  transform:translate(0%);
	  opacity:1; }
	
	.bp3-panel-stack-push .bp3-panel-stack-exit-active{
	  -webkit-transform:translateX(-50%);
			  transform:translateX(-50%);
	  opacity:0;
	  -webkit-transition-delay:0;
			  transition-delay:0;
	  -webkit-transition-duration:400ms;
			  transition-duration:400ms;
	  -webkit-transition-property:opacity, -webkit-transform;
	  transition-property:opacity, -webkit-transform;
	  transition-property:transform, opacity;
	  transition-property:transform, opacity, -webkit-transform;
	  -webkit-transition-timing-function:ease;
			  transition-timing-function:ease; }
	
	.bp3-panel-stack-pop .bp3-panel-stack-enter, .bp3-panel-stack-pop .bp3-panel-stack-appear{
	  -webkit-transform:translateX(-50%);
			  transform:translateX(-50%);
	  opacity:0; }
	
	.bp3-panel-stack-pop .bp3-panel-stack-enter-active, .bp3-panel-stack-pop .bp3-panel-stack-appear-active{
	  -webkit-transform:translate(0%);
			  transform:translate(0%);
	  opacity:1;
	  -webkit-transition-delay:0;
			  transition-delay:0;
	  -webkit-transition-duration:400ms;
			  transition-duration:400ms;
	  -webkit-transition-property:opacity, -webkit-transform;
	  transition-property:opacity, -webkit-transform;
	  transition-property:transform, opacity;
	  transition-property:transform, opacity, -webkit-transform;
	  -webkit-transition-timing-function:ease;
			  transition-timing-function:ease; }
	
	.bp3-panel-stack-pop .bp3-panel-stack-exit{
	  -webkit-transform:translate(0%);
			  transform:translate(0%);
	  opacity:1; }
	
	.bp3-panel-stack-pop .bp3-panel-stack-exit-active{
	  -webkit-transform:translateX(100%);
			  transform:translateX(100%);
	  opacity:0;
	  -webkit-transition-delay:0;
			  transition-delay:0;
	  -webkit-transition-duration:400ms;
			  transition-duration:400ms;
	  -webkit-transition-property:opacity, -webkit-transform;
	  transition-property:opacity, -webkit-transform;
	  transition-property:transform, opacity;
	  transition-property:transform, opacity, -webkit-transform;
	  -webkit-transition-timing-function:ease;
			  transition-timing-function:ease; }
	.bp3-popover{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
	  -webkit-transform:scale(1);
			  transform:scale(1);
	  border-radius:3px;
	  display:inline-block;
	  z-index:20; }
	  .bp3-popover .bp3-popover-arrow{
		height:30px;
		position:absolute;
		width:30px; }
		.bp3-popover .bp3-popover-arrow::before{
		  height:20px;
		  margin:5px;
		  width:20px; }
	  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-popover{
		margin-bottom:17px;
		margin-top:-17px; }
		.bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-popover > .bp3-popover-arrow{
		  bottom:-11px; }
		  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-popover > .bp3-popover-arrow svg{
			-webkit-transform:rotate(-90deg);
					transform:rotate(-90deg); }
	  .bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-popover{
		margin-left:17px; }
		.bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-popover > .bp3-popover-arrow{
		  left:-11px; }
		  .bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-popover > .bp3-popover-arrow svg{
			-webkit-transform:rotate(0);
					transform:rotate(0); }
	  .bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-popover{
		margin-top:17px; }
		.bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-popover > .bp3-popover-arrow{
		  top:-11px; }
		  .bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-popover > .bp3-popover-arrow svg{
			-webkit-transform:rotate(90deg);
					transform:rotate(90deg); }
	  .bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-popover{
		margin-left:-17px;
		margin-right:17px; }
		.bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-popover > .bp3-popover-arrow{
		  right:-11px; }
		  .bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-popover > .bp3-popover-arrow svg{
			-webkit-transform:rotate(180deg);
					transform:rotate(180deg); }
	  .bp3-tether-element-attached-middle > .bp3-popover > .bp3-popover-arrow{
		top:50%;
		-webkit-transform:translateY(-50%);
				transform:translateY(-50%); }
	  .bp3-tether-element-attached-center > .bp3-popover > .bp3-popover-arrow{
		right:50%;
		-webkit-transform:translateX(50%);
				transform:translateX(50%); }
	  .bp3-tether-element-attached-top.bp3-tether-target-attached-top > .bp3-popover > .bp3-popover-arrow{
		top:-0.3934px; }
	  .bp3-tether-element-attached-right.bp3-tether-target-attached-right > .bp3-popover > .bp3-popover-arrow{
		right:-0.3934px; }
	  .bp3-tether-element-attached-left.bp3-tether-target-attached-left > .bp3-popover > .bp3-popover-arrow{
		left:-0.3934px; }
	  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-bottom > .bp3-popover > .bp3-popover-arrow{
		bottom:-0.3934px; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-left > .bp3-popover{
		-webkit-transform-origin:top left;
				transform-origin:top left; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-center > .bp3-popover{
		-webkit-transform-origin:top center;
				transform-origin:top center; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-right > .bp3-popover{
		-webkit-transform-origin:top right;
				transform-origin:top right; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-left > .bp3-popover{
		-webkit-transform-origin:center left;
				transform-origin:center left; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-center > .bp3-popover{
		-webkit-transform-origin:center center;
				transform-origin:center center; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-right > .bp3-popover{
		-webkit-transform-origin:center right;
				transform-origin:center right; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-left > .bp3-popover{
		-webkit-transform-origin:bottom left;
				transform-origin:bottom left; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-center > .bp3-popover{
		-webkit-transform-origin:bottom center;
				transform-origin:bottom center; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-right > .bp3-popover{
		-webkit-transform-origin:bottom right;
				transform-origin:bottom right; }
	  .bp3-popover .bp3-popover-content{
		background:#ffffff;
		color:inherit; }
	  .bp3-popover .bp3-popover-arrow::before{
		-webkit-box-shadow:1px 1px 6px rgba(16, 22, 26, 0.2);
				box-shadow:1px 1px 6px rgba(16, 22, 26, 0.2); }
	  .bp3-popover .bp3-popover-arrow-border{
		fill:#10161a;
		fill-opacity:0.1; }
	  .bp3-popover .bp3-popover-arrow-fill{
		fill:#ffffff; }
	  .bp3-popover-enter > .bp3-popover, .bp3-popover-appear > .bp3-popover{
		-webkit-transform:scale(0.3);
				transform:scale(0.3); }
	  .bp3-popover-enter-active > .bp3-popover, .bp3-popover-appear-active > .bp3-popover{
		-webkit-transform:scale(1);
				transform:scale(1);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	  .bp3-popover-exit > .bp3-popover{
		-webkit-transform:scale(1);
				transform:scale(1); }
	  .bp3-popover-exit-active > .bp3-popover{
		-webkit-transform:scale(0.3);
				transform:scale(0.3);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	  .bp3-popover .bp3-popover-content{
		border-radius:3px;
		position:relative; }
	  .bp3-popover.bp3-popover-content-sizing .bp3-popover-content{
		max-width:350px;
		padding:20px; }
	  .bp3-popover-target + .bp3-overlay .bp3-popover.bp3-popover-content-sizing{
		width:350px; }
	  .bp3-popover.bp3-minimal{
		margin:0 !important; }
		.bp3-popover.bp3-minimal .bp3-popover-arrow{
		  display:none; }
		.bp3-popover.bp3-minimal.bp3-popover{
		  -webkit-transform:scale(1);
				  transform:scale(1); }
		  .bp3-popover-enter > .bp3-popover.bp3-minimal.bp3-popover, .bp3-popover-appear > .bp3-popover.bp3-minimal.bp3-popover{
			-webkit-transform:scale(1);
					transform:scale(1); }
		  .bp3-popover-enter-active > .bp3-popover.bp3-minimal.bp3-popover, .bp3-popover-appear-active > .bp3-popover.bp3-minimal.bp3-popover{
			-webkit-transform:scale(1);
					transform:scale(1);
			-webkit-transition-delay:0;
					transition-delay:0;
			-webkit-transition-duration:100ms;
					transition-duration:100ms;
			-webkit-transition-property:-webkit-transform;
			transition-property:-webkit-transform;
			transition-property:transform;
			transition-property:transform, -webkit-transform;
			-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
					transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
		  .bp3-popover-exit > .bp3-popover.bp3-minimal.bp3-popover{
			-webkit-transform:scale(1);
					transform:scale(1); }
		  .bp3-popover-exit-active > .bp3-popover.bp3-minimal.bp3-popover{
			-webkit-transform:scale(1);
					transform:scale(1);
			-webkit-transition-delay:0;
					transition-delay:0;
			-webkit-transition-duration:100ms;
					transition-duration:100ms;
			-webkit-transition-property:-webkit-transform;
			transition-property:-webkit-transform;
			transition-property:transform;
			transition-property:transform, -webkit-transform;
			-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
					transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-popover.bp3-dark,
	  .bp3-dark .bp3-popover{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
		.bp3-popover.bp3-dark .bp3-popover-content,
		.bp3-dark .bp3-popover .bp3-popover-content{
		  background:#30404d;
		  color:inherit; }
		.bp3-popover.bp3-dark .bp3-popover-arrow::before,
		.bp3-dark .bp3-popover .bp3-popover-arrow::before{
		  -webkit-box-shadow:1px 1px 6px rgba(16, 22, 26, 0.4);
				  box-shadow:1px 1px 6px rgba(16, 22, 26, 0.4); }
		.bp3-popover.bp3-dark .bp3-popover-arrow-border,
		.bp3-dark .bp3-popover .bp3-popover-arrow-border{
		  fill:#10161a;
		  fill-opacity:0.2; }
		.bp3-popover.bp3-dark .bp3-popover-arrow-fill,
		.bp3-dark .bp3-popover .bp3-popover-arrow-fill{
		  fill:#30404d; }
	
	.bp3-popover-arrow::before{
	  border-radius:2px;
	  content:"";
	  display:block;
	  position:absolute;
	  -webkit-transform:rotate(45deg);
			  transform:rotate(45deg); }
	
	.bp3-tether-pinned .bp3-popover-arrow{
	  display:none; }
	
	.bp3-popover-backdrop{
	  background:rgba(255, 255, 255, 0); }
	
	.bp3-transition-container{
	  opacity:1;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  z-index:20; }
	  .bp3-transition-container.bp3-popover-enter, .bp3-transition-container.bp3-popover-appear{
		opacity:0; }
	  .bp3-transition-container.bp3-popover-enter-active, .bp3-transition-container.bp3-popover-appear-active{
		opacity:1;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:100ms;
				transition-duration:100ms;
		-webkit-transition-property:opacity;
		transition-property:opacity;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-transition-container.bp3-popover-exit{
		opacity:1; }
	  .bp3-transition-container.bp3-popover-exit-active{
		opacity:0;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:100ms;
				transition-duration:100ms;
		-webkit-transition-property:opacity;
		transition-property:opacity;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-transition-container:focus{
		outline:none; }
	  .bp3-transition-container.bp3-popover-leave .bp3-popover-content{
		pointer-events:none; }
	  .bp3-transition-container[data-x-out-of-boundaries]{
		display:none; }
	
	span.bp3-popover-target{
	  display:inline-block; }
	
	.bp3-popover-wrapper.bp3-fill{
	  width:100%; }
	
	.bp3-portal{
	  left:0;
	  position:absolute;
	  right:0;
	  top:0; }
	@-webkit-keyframes linear-progress-bar-stripes{
	  from{
		background-position:0 0; }
	  to{
		background-position:30px 0; } }
	@keyframes linear-progress-bar-stripes{
	  from{
		background-position:0 0; }
	  to{
		background-position:30px 0; } }
	
	.bp3-progress-bar{
	  background:rgba(92, 112, 128, 0.2);
	  border-radius:40px;
	  display:block;
	  height:8px;
	  overflow:hidden;
	  position:relative;
	  width:100%; }
	  .bp3-progress-bar .bp3-progress-meter{
		background:linear-gradient(-45deg, rgba(255, 255, 255, 0.2) 25%, transparent 25%, transparent 50%, rgba(255, 255, 255, 0.2) 50%, rgba(255, 255, 255, 0.2) 75%, transparent 75%);
		background-color:rgba(92, 112, 128, 0.8);
		background-size:30px 30px;
		border-radius:40px;
		height:100%;
		position:absolute;
		-webkit-transition:width 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:width 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		width:100%; }
	  .bp3-progress-bar:not(.bp3-no-animation):not(.bp3-no-stripes) .bp3-progress-meter{
		animation:linear-progress-bar-stripes 300ms linear infinite reverse; }
	  .bp3-progress-bar.bp3-no-stripes .bp3-progress-meter{
		background-image:none; }
	
	.bp3-dark .bp3-progress-bar{
	  background:rgba(16, 22, 26, 0.5); }
	  .bp3-dark .bp3-progress-bar .bp3-progress-meter{
		background-color:#8a9ba8; }
	
	.bp3-progress-bar.bp3-intent-primary .bp3-progress-meter{
	  background-color:#137cbd; }
	
	.bp3-progress-bar.bp3-intent-success .bp3-progress-meter{
	  background-color:#0f9960; }
	
	.bp3-progress-bar.bp3-intent-warning .bp3-progress-meter{
	  background-color:#d9822b; }
	
	.bp3-progress-bar.bp3-intent-danger .bp3-progress-meter{
	  background-color:#db3737; }
	@-webkit-keyframes skeleton-glow{
	  from{
		background:rgba(206, 217, 224, 0.2);
		border-color:rgba(206, 217, 224, 0.2); }
	  to{
		background:rgba(92, 112, 128, 0.2);
		border-color:rgba(92, 112, 128, 0.2); } }
	@keyframes skeleton-glow{
	  from{
		background:rgba(206, 217, 224, 0.2);
		border-color:rgba(206, 217, 224, 0.2); }
	  to{
		background:rgba(92, 112, 128, 0.2);
		border-color:rgba(92, 112, 128, 0.2); } }
	.bp3-skeleton{
	  -webkit-animation:1000ms linear infinite alternate skeleton-glow;
			  animation:1000ms linear infinite alternate skeleton-glow;
	  background:rgba(206, 217, 224, 0.2);
	  background-clip:padding-box !important;
	  border-color:rgba(206, 217, 224, 0.2) !important;
	  border-radius:2px;
	  -webkit-box-shadow:none !important;
			  box-shadow:none !important;
	  color:transparent !important;
	  cursor:default;
	  pointer-events:none;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-skeleton::before, .bp3-skeleton::after,
	  .bp3-skeleton *{
		visibility:hidden !important; }
	.bp3-slider{
	  height:40px;
	  min-width:150px;
	  width:100%;
	  cursor:default;
	  outline:none;
	  position:relative;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-slider:hover{
		cursor:pointer; }
	  .bp3-slider:active{
		cursor:-webkit-grabbing;
		cursor:grabbing; }
	  .bp3-slider.bp3-disabled{
		cursor:not-allowed;
		opacity:0.5; }
	  .bp3-slider.bp3-slider-unlabeled{
		height:16px; }
	
	.bp3-slider-track,
	.bp3-slider-progress{
	  height:6px;
	  left:0;
	  right:0;
	  top:5px;
	  position:absolute; }
	
	.bp3-slider-track{
	  border-radius:3px;
	  overflow:hidden; }
	
	.bp3-slider-progress{
	  background:rgba(92, 112, 128, 0.2); }
	  .bp3-dark .bp3-slider-progress{
		background:rgba(16, 22, 26, 0.5); }
	  .bp3-slider-progress.bp3-intent-primary{
		background-color:#137cbd; }
	  .bp3-slider-progress.bp3-intent-success{
		background-color:#0f9960; }
	  .bp3-slider-progress.bp3-intent-warning{
		background-color:#d9822b; }
	  .bp3-slider-progress.bp3-intent-danger{
		background-color:#db3737; }
	
	.bp3-slider-handle{
	  background-color:#f5f8fa;
	  background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.8)), to(rgba(255, 255, 255, 0)));
	  background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0));
	  -webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
			  box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
	  color:#182026;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
	  cursor:pointer;
	  height:16px;
	  left:0;
	  position:absolute;
	  top:0;
	  width:16px; }
	  .bp3-slider-handle:hover{
		background-clip:padding-box;
		background-color:#ebf1f5;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1); }
	  .bp3-slider-handle:active, .bp3-slider-handle.bp3-active{
		background-color:#d8e1e8;
		background-image:none;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
	  .bp3-slider-handle:disabled, .bp3-slider-handle.bp3-disabled{
		background-color:rgba(206, 217, 224, 0.5);
		background-image:none;
		-webkit-box-shadow:none;
				box-shadow:none;
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed;
		outline:none; }
		.bp3-slider-handle:disabled.bp3-active, .bp3-slider-handle:disabled.bp3-active:hover, .bp3-slider-handle.bp3-disabled.bp3-active, .bp3-slider-handle.bp3-disabled.bp3-active:hover{
		  background:rgba(206, 217, 224, 0.7); }
	  .bp3-slider-handle:focus{
		z-index:1; }
	  .bp3-slider-handle:hover{
		background-clip:padding-box;
		background-color:#ebf1f5;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 -1px 0 rgba(16, 22, 26, 0.1);
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 1px 1px rgba(16, 22, 26, 0.2);
		cursor:-webkit-grab;
		cursor:grab;
		z-index:2; }
	  .bp3-slider-handle.bp3-active{
		background-color:#d8e1e8;
		background-image:none;
		-webkit-box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				box-shadow:inset 0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 2px rgba(16, 22, 26, 0.2);
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 1px rgba(16, 22, 26, 0.1);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), inset 0 1px 1px rgba(16, 22, 26, 0.1);
		cursor:-webkit-grabbing;
		cursor:grabbing; }
	  .bp3-disabled .bp3-slider-handle{
		background:#bfccd6;
		-webkit-box-shadow:none;
				box-shadow:none;
		pointer-events:none; }
	  .bp3-dark .bp3-slider-handle{
		background-color:#394b59;
		background-image:-webkit-gradient(linear, left top, left bottom, from(rgba(255, 255, 255, 0.05)), to(rgba(255, 255, 255, 0)));
		background-image:linear-gradient(to bottom, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0));
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
		color:#f5f8fa; }
		.bp3-dark .bp3-slider-handle:hover, .bp3-dark .bp3-slider-handle:active, .bp3-dark .bp3-slider-handle.bp3-active{
		  color:#f5f8fa; }
		.bp3-dark .bp3-slider-handle:hover{
		  background-color:#30404d;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-slider-handle:active, .bp3-dark .bp3-slider-handle.bp3-active{
		  background-color:#202b33;
		  background-image:none;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.6), inset 0 1px 2px rgba(16, 22, 26, 0.2); }
		.bp3-dark .bp3-slider-handle:disabled, .bp3-dark .bp3-slider-handle.bp3-disabled{
		  background-color:rgba(57, 75, 89, 0.5);
		  background-image:none;
		  -webkit-box-shadow:none;
				  box-shadow:none;
		  color:rgba(167, 182, 194, 0.6); }
		  .bp3-dark .bp3-slider-handle:disabled.bp3-active, .bp3-dark .bp3-slider-handle.bp3-disabled.bp3-active{
			background:rgba(57, 75, 89, 0.7); }
		.bp3-dark .bp3-slider-handle .bp3-button-spinner .bp3-spinner-head{
		  background:rgba(16, 22, 26, 0.5);
		  stroke:#8a9ba8; }
		.bp3-dark .bp3-slider-handle, .bp3-dark .bp3-slider-handle:hover{
		  background-color:#394b59; }
		.bp3-dark .bp3-slider-handle.bp3-active{
		  background-color:#293742; }
	  .bp3-dark .bp3-disabled .bp3-slider-handle{
		background:#5c7080;
		border-color:#5c7080;
		-webkit-box-shadow:none;
				box-shadow:none; }
	  .bp3-slider-handle .bp3-slider-label{
		background:#394b59;
		border-radius:3px;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
		color:#f5f8fa;
		margin-left:8px; }
		.bp3-dark .bp3-slider-handle .bp3-slider-label{
		  background:#e1e8ed;
		  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
		  color:#394b59; }
		.bp3-disabled .bp3-slider-handle .bp3-slider-label{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-slider-handle.bp3-start, .bp3-slider-handle.bp3-end{
		width:8px; }
	  .bp3-slider-handle.bp3-start{
		border-bottom-right-radius:0;
		border-top-right-radius:0; }
	  .bp3-slider-handle.bp3-end{
		border-bottom-left-radius:0;
		border-top-left-radius:0;
		margin-left:8px; }
		.bp3-slider-handle.bp3-end .bp3-slider-label{
		  margin-left:0; }
	
	.bp3-slider-label{
	  -webkit-transform:translate(-50%, 20px);
			  transform:translate(-50%, 20px);
	  display:inline-block;
	  font-size:12px;
	  line-height:1;
	  padding:2px 5px;
	  position:absolute;
	  vertical-align:top; }
	
	.bp3-slider.bp3-vertical{
	  height:150px;
	  min-width:40px;
	  width:40px; }
	  .bp3-slider.bp3-vertical .bp3-slider-track,
	  .bp3-slider.bp3-vertical .bp3-slider-progress{
		bottom:0;
		height:auto;
		left:5px;
		top:0;
		width:6px; }
	  .bp3-slider.bp3-vertical .bp3-slider-progress{
		top:auto; }
	  .bp3-slider.bp3-vertical .bp3-slider-label{
		-webkit-transform:translate(20px, 50%);
				transform:translate(20px, 50%); }
	  .bp3-slider.bp3-vertical .bp3-slider-handle{
		top:auto; }
		.bp3-slider.bp3-vertical .bp3-slider-handle .bp3-slider-label{
		  margin-left:0;
		  margin-top:-8px; }
		.bp3-slider.bp3-vertical .bp3-slider-handle.bp3-end, .bp3-slider.bp3-vertical .bp3-slider-handle.bp3-start{
		  height:8px;
		  margin-left:0;
		  width:16px; }
		.bp3-slider.bp3-vertical .bp3-slider-handle.bp3-start{
		  border-bottom-right-radius:3px;
		  border-top-left-radius:0; }
		  .bp3-slider.bp3-vertical .bp3-slider-handle.bp3-start .bp3-slider-label{
			-webkit-transform:translate(20px);
					transform:translate(20px); }
		.bp3-slider.bp3-vertical .bp3-slider-handle.bp3-end{
		  border-bottom-left-radius:0;
		  border-bottom-right-radius:0;
		  border-top-left-radius:3px;
		  margin-bottom:8px; }
	
	@-webkit-keyframes pt-spinner-animation{
	  from{
		-webkit-transform:rotate(0deg);
				transform:rotate(0deg); }
	  to{
		-webkit-transform:rotate(360deg);
				transform:rotate(360deg); } }
	
	@keyframes pt-spinner-animation{
	  from{
		-webkit-transform:rotate(0deg);
				transform:rotate(0deg); }
	  to{
		-webkit-transform:rotate(360deg);
				transform:rotate(360deg); } }
	
	.bp3-spinner{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-pack:center;
		  -ms-flex-pack:center;
			  justify-content:center;
	  overflow:visible;
	  vertical-align:middle; }
	  .bp3-spinner svg{
		display:block; }
	  .bp3-spinner path{
		fill-opacity:0; }
	  .bp3-spinner .bp3-spinner-head{
		stroke:rgba(92, 112, 128, 0.8);
		stroke-linecap:round;
		-webkit-transform-origin:center;
				transform-origin:center;
		-webkit-transition:stroke-dashoffset 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
		transition:stroke-dashoffset 200ms cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-spinner .bp3-spinner-track{
		stroke:rgba(92, 112, 128, 0.2); }
	
	.bp3-spinner-animation{
	  -webkit-animation:pt-spinner-animation 500ms linear infinite;
			  animation:pt-spinner-animation 500ms linear infinite; }
	  .bp3-no-spin > .bp3-spinner-animation{
		-webkit-animation:none;
				animation:none; }
	
	.bp3-dark .bp3-spinner .bp3-spinner-head{
	  stroke:#8a9ba8; }
	
	.bp3-dark .bp3-spinner .bp3-spinner-track{
	  stroke:rgba(16, 22, 26, 0.5); }
	
	.bp3-spinner.bp3-intent-primary .bp3-spinner-head{
	  stroke:#137cbd; }
	
	.bp3-spinner.bp3-intent-success .bp3-spinner-head{
	  stroke:#0f9960; }
	
	.bp3-spinner.bp3-intent-warning .bp3-spinner-head{
	  stroke:#d9822b; }
	
	.bp3-spinner.bp3-intent-danger .bp3-spinner-head{
	  stroke:#db3737; }
	.bp3-tabs.bp3-vertical{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex; }
	  .bp3-tabs.bp3-vertical > .bp3-tab-list{
		-webkit-box-align:start;
			-ms-flex-align:start;
				align-items:flex-start;
		-webkit-box-orient:vertical;
		-webkit-box-direction:normal;
			-ms-flex-direction:column;
				flex-direction:column; }
		.bp3-tabs.bp3-vertical > .bp3-tab-list .bp3-tab{
		  border-radius:3px;
		  padding:0 10px;
		  width:100%; }
		  .bp3-tabs.bp3-vertical > .bp3-tab-list .bp3-tab[aria-selected="true"]{
			background-color:rgba(19, 124, 189, 0.2);
			-webkit-box-shadow:none;
					box-shadow:none; }
		.bp3-tabs.bp3-vertical > .bp3-tab-list .bp3-tab-indicator-wrapper .bp3-tab-indicator{
		  background-color:rgba(19, 124, 189, 0.2);
		  border-radius:3px;
		  bottom:0;
		  height:auto;
		  left:0;
		  right:0;
		  top:0; }
	  .bp3-tabs.bp3-vertical > .bp3-tab-panel{
		margin-top:0;
		padding-left:20px; }
	
	.bp3-tab-list{
	  -webkit-box-align:end;
		  -ms-flex-align:end;
			  align-items:flex-end;
	  border:none;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  list-style:none;
	  margin:0;
	  padding:0;
	  position:relative; }
	  .bp3-tab-list > *:not(:last-child){
		margin-right:20px; }
	
	.bp3-tab{
	  overflow:hidden;
	  text-overflow:ellipsis;
	  white-space:nowrap;
	  word-wrap:normal;
	  color:#182026;
	  cursor:pointer;
	  -webkit-box-flex:0;
		  -ms-flex:0 0 auto;
			  flex:0 0 auto;
	  font-size:14px;
	  line-height:30px;
	  max-width:100%;
	  position:relative;
	  vertical-align:top; }
	  .bp3-tab a{
		color:inherit;
		display:block;
		text-decoration:none; }
	  .bp3-tab-indicator-wrapper ~ .bp3-tab{
		background-color:transparent !important;
		-webkit-box-shadow:none !important;
				box-shadow:none !important; }
	  .bp3-tab[aria-disabled="true"]{
		color:rgba(92, 112, 128, 0.6);
		cursor:not-allowed; }
	  .bp3-tab[aria-selected="true"]{
		border-radius:0;
		-webkit-box-shadow:inset 0 -3px 0 #106ba3;
				box-shadow:inset 0 -3px 0 #106ba3; }
	  .bp3-tab[aria-selected="true"], .bp3-tab:not([aria-disabled="true"]):hover{
		color:#106ba3; }
	  .bp3-tab:focus{
		-moz-outline-radius:0; }
	  .bp3-large > .bp3-tab{
		font-size:16px;
		line-height:40px; }
	
	.bp3-tab-panel{
	  margin-top:20px; }
	  .bp3-tab-panel[aria-hidden="true"]{
		display:none; }
	
	.bp3-tab-indicator-wrapper{
	  left:0;
	  pointer-events:none;
	  position:absolute;
	  top:0;
	  -webkit-transform:translateX(0), translateY(0);
			  transform:translateX(0), translateY(0);
	  -webkit-transition:height, width, -webkit-transform;
	  transition:height, width, -webkit-transform;
	  transition:height, transform, width;
	  transition:height, transform, width, -webkit-transform;
	  -webkit-transition-duration:200ms;
			  transition-duration:200ms;
	  -webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
			  transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-tab-indicator-wrapper .bp3-tab-indicator{
		background-color:#106ba3;
		bottom:0;
		height:3px;
		left:0;
		position:absolute;
		right:0; }
	  .bp3-tab-indicator-wrapper.bp3-no-animation{
		-webkit-transition:none;
		transition:none; }
	
	.bp3-dark .bp3-tab{
	  color:#f5f8fa; }
	  .bp3-dark .bp3-tab[aria-disabled="true"]{
		color:rgba(167, 182, 194, 0.6); }
	  .bp3-dark .bp3-tab[aria-selected="true"]{
		-webkit-box-shadow:inset 0 -3px 0 #48aff0;
				box-shadow:inset 0 -3px 0 #48aff0; }
	  .bp3-dark .bp3-tab[aria-selected="true"], .bp3-dark .bp3-tab:not([aria-disabled="true"]):hover{
		color:#48aff0; }
	
	.bp3-dark .bp3-tab-indicator{
	  background-color:#48aff0; }
	
	.bp3-flex-expander{
	  -webkit-box-flex:1;
		  -ms-flex:1 1;
			  flex:1 1; }
	.bp3-tag{
	  display:-webkit-inline-box;
	  display:-ms-inline-flexbox;
	  display:inline-flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  background-color:#5c7080;
	  border:none;
	  border-radius:3px;
	  -webkit-box-shadow:none;
			  box-shadow:none;
	  color:#f5f8fa;
	  font-size:12px;
	  line-height:16px;
	  max-width:100%;
	  min-height:20px;
	  min-width:20px;
	  padding:2px 6px;
	  position:relative; }
	  .bp3-tag.bp3-interactive{
		cursor:pointer; }
		.bp3-tag.bp3-interactive:hover{
		  background-color:rgba(92, 112, 128, 0.85); }
		.bp3-tag.bp3-interactive.bp3-active, .bp3-tag.bp3-interactive:active{
		  background-color:rgba(92, 112, 128, 0.7); }
	  .bp3-tag > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-tag > .bp3-fill{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-tag::before,
	  .bp3-tag > *{
		margin-right:4px; }
	  .bp3-tag:empty::before,
	  .bp3-tag > :last-child{
		margin-right:0; }
	  .bp3-tag:focus{
		outline:rgba(19, 124, 189, 0.6) auto 2px;
		outline-offset:0;
		-moz-outline-radius:6px; }
	  .bp3-tag.bp3-round{
		border-radius:30px;
		padding-left:8px;
		padding-right:8px; }
	  .bp3-dark .bp3-tag{
		background-color:#bfccd6;
		color:#182026; }
		.bp3-dark .bp3-tag.bp3-interactive{
		  cursor:pointer; }
		  .bp3-dark .bp3-tag.bp3-interactive:hover{
			background-color:rgba(191, 204, 214, 0.85); }
		  .bp3-dark .bp3-tag.bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-interactive:active{
			background-color:rgba(191, 204, 214, 0.7); }
		.bp3-dark .bp3-tag > .bp3-icon, .bp3-dark .bp3-tag .bp3-icon-standard, .bp3-dark .bp3-tag .bp3-icon-large{
		  fill:currentColor; }
	  .bp3-tag > .bp3-icon, .bp3-tag .bp3-icon-standard, .bp3-tag .bp3-icon-large{
		fill:#ffffff; }
	  .bp3-tag.bp3-large,
	  .bp3-large .bp3-tag{
		font-size:14px;
		line-height:20px;
		min-height:30px;
		min-width:30px;
		padding:5px 10px; }
		.bp3-tag.bp3-large::before,
		.bp3-tag.bp3-large > *,
		.bp3-large .bp3-tag::before,
		.bp3-large .bp3-tag > *{
		  margin-right:7px; }
		.bp3-tag.bp3-large:empty::before,
		.bp3-tag.bp3-large > :last-child,
		.bp3-large .bp3-tag:empty::before,
		.bp3-large .bp3-tag > :last-child{
		  margin-right:0; }
		.bp3-tag.bp3-large.bp3-round,
		.bp3-large .bp3-tag.bp3-round{
		  padding-left:12px;
		  padding-right:12px; }
	  .bp3-tag.bp3-intent-primary{
		background:#137cbd;
		color:#ffffff; }
		.bp3-tag.bp3-intent-primary.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-intent-primary.bp3-interactive:hover{
			background-color:rgba(19, 124, 189, 0.85); }
		  .bp3-tag.bp3-intent-primary.bp3-interactive.bp3-active, .bp3-tag.bp3-intent-primary.bp3-interactive:active{
			background-color:rgba(19, 124, 189, 0.7); }
	  .bp3-tag.bp3-intent-success{
		background:#0f9960;
		color:#ffffff; }
		.bp3-tag.bp3-intent-success.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-intent-success.bp3-interactive:hover{
			background-color:rgba(15, 153, 96, 0.85); }
		  .bp3-tag.bp3-intent-success.bp3-interactive.bp3-active, .bp3-tag.bp3-intent-success.bp3-interactive:active{
			background-color:rgba(15, 153, 96, 0.7); }
	  .bp3-tag.bp3-intent-warning{
		background:#d9822b;
		color:#ffffff; }
		.bp3-tag.bp3-intent-warning.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-intent-warning.bp3-interactive:hover{
			background-color:rgba(217, 130, 43, 0.85); }
		  .bp3-tag.bp3-intent-warning.bp3-interactive.bp3-active, .bp3-tag.bp3-intent-warning.bp3-interactive:active{
			background-color:rgba(217, 130, 43, 0.7); }
	  .bp3-tag.bp3-intent-danger{
		background:#db3737;
		color:#ffffff; }
		.bp3-tag.bp3-intent-danger.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-intent-danger.bp3-interactive:hover{
			background-color:rgba(219, 55, 55, 0.85); }
		  .bp3-tag.bp3-intent-danger.bp3-interactive.bp3-active, .bp3-tag.bp3-intent-danger.bp3-interactive:active{
			background-color:rgba(219, 55, 55, 0.7); }
	  .bp3-tag.bp3-fill{
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		width:100%; }
	  .bp3-tag.bp3-minimal > .bp3-icon, .bp3-tag.bp3-minimal .bp3-icon-standard, .bp3-tag.bp3-minimal .bp3-icon-large{
		fill:#5c7080; }
	  .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]){
		background-color:rgba(138, 155, 168, 0.2);
		color:#182026; }
		.bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive:hover{
			background-color:rgba(92, 112, 128, 0.3); }
		  .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive.bp3-active, .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive:active{
			background-color:rgba(92, 112, 128, 0.4); }
		.bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]){
		  color:#f5f8fa; }
		  .bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive{
			cursor:pointer; }
			.bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive:hover{
			  background-color:rgba(191, 204, 214, 0.3); }
			.bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]).bp3-interactive:active{
			  background-color:rgba(191, 204, 214, 0.4); }
		  .bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]) > .bp3-icon, .bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]) .bp3-icon-standard, .bp3-dark .bp3-tag.bp3-minimal:not([class*="bp3-intent-"]) .bp3-icon-large{
			fill:#a7b6c2; }
	  .bp3-tag.bp3-minimal.bp3-intent-primary{
		background-color:rgba(19, 124, 189, 0.15);
		color:#106ba3; }
		.bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive:hover{
			background-color:rgba(19, 124, 189, 0.25); }
		  .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive.bp3-active, .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive:active{
			background-color:rgba(19, 124, 189, 0.35); }
		.bp3-tag.bp3-minimal.bp3-intent-primary > .bp3-icon, .bp3-tag.bp3-minimal.bp3-intent-primary .bp3-icon-standard, .bp3-tag.bp3-minimal.bp3-intent-primary .bp3-icon-large{
		  fill:#137cbd; }
		.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-primary{
		  background-color:rgba(19, 124, 189, 0.25);
		  color:#48aff0; }
		  .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive{
			cursor:pointer; }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive:hover{
			  background-color:rgba(19, 124, 189, 0.35); }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-primary.bp3-interactive:active{
			  background-color:rgba(19, 124, 189, 0.45); }
	  .bp3-tag.bp3-minimal.bp3-intent-success{
		background-color:rgba(15, 153, 96, 0.15);
		color:#0d8050; }
		.bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive:hover{
			background-color:rgba(15, 153, 96, 0.25); }
		  .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive.bp3-active, .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive:active{
			background-color:rgba(15, 153, 96, 0.35); }
		.bp3-tag.bp3-minimal.bp3-intent-success > .bp3-icon, .bp3-tag.bp3-minimal.bp3-intent-success .bp3-icon-standard, .bp3-tag.bp3-minimal.bp3-intent-success .bp3-icon-large{
		  fill:#0f9960; }
		.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-success{
		  background-color:rgba(15, 153, 96, 0.25);
		  color:#3dcc91; }
		  .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive{
			cursor:pointer; }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive:hover{
			  background-color:rgba(15, 153, 96, 0.35); }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-success.bp3-interactive:active{
			  background-color:rgba(15, 153, 96, 0.45); }
	  .bp3-tag.bp3-minimal.bp3-intent-warning{
		background-color:rgba(217, 130, 43, 0.15);
		color:#bf7326; }
		.bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive:hover{
			background-color:rgba(217, 130, 43, 0.25); }
		  .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive.bp3-active, .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive:active{
			background-color:rgba(217, 130, 43, 0.35); }
		.bp3-tag.bp3-minimal.bp3-intent-warning > .bp3-icon, .bp3-tag.bp3-minimal.bp3-intent-warning .bp3-icon-standard, .bp3-tag.bp3-minimal.bp3-intent-warning .bp3-icon-large{
		  fill:#d9822b; }
		.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-warning{
		  background-color:rgba(217, 130, 43, 0.25);
		  color:#ffb366; }
		  .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive{
			cursor:pointer; }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive:hover{
			  background-color:rgba(217, 130, 43, 0.35); }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-warning.bp3-interactive:active{
			  background-color:rgba(217, 130, 43, 0.45); }
	  .bp3-tag.bp3-minimal.bp3-intent-danger{
		background-color:rgba(219, 55, 55, 0.15);
		color:#c23030; }
		.bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive{
		  cursor:pointer; }
		  .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive:hover{
			background-color:rgba(219, 55, 55, 0.25); }
		  .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive.bp3-active, .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive:active{
			background-color:rgba(219, 55, 55, 0.35); }
		.bp3-tag.bp3-minimal.bp3-intent-danger > .bp3-icon, .bp3-tag.bp3-minimal.bp3-intent-danger .bp3-icon-standard, .bp3-tag.bp3-minimal.bp3-intent-danger .bp3-icon-large{
		  fill:#db3737; }
		.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-danger{
		  background-color:rgba(219, 55, 55, 0.25);
		  color:#ff7373; }
		  .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive{
			cursor:pointer; }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive:hover{
			  background-color:rgba(219, 55, 55, 0.35); }
			.bp3-dark .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive.bp3-active, .bp3-dark .bp3-tag.bp3-minimal.bp3-intent-danger.bp3-interactive:active{
			  background-color:rgba(219, 55, 55, 0.45); }
	
	.bp3-tag-remove{
	  background:none;
	  border:none;
	  color:inherit;
	  cursor:pointer;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  margin-bottom:-2px;
	  margin-right:-6px !important;
	  margin-top:-2px;
	  opacity:0.5;
	  padding:2px;
	  padding-left:0; }
	  .bp3-tag-remove:hover{
		background:none;
		opacity:0.8;
		text-decoration:none; }
	  .bp3-tag-remove:active{
		opacity:1; }
	  .bp3-tag-remove:empty::before{
		font-family:"Icons16", sans-serif;
		font-size:16px;
		font-style:normal;
		font-weight:400;
		line-height:1;
		-moz-osx-font-smoothing:grayscale;
		-webkit-font-smoothing:antialiased;
		content:""; }
	  .bp3-large .bp3-tag-remove{
		margin-right:-10px !important;
		padding:0 5px 0 0; }
		.bp3-large .bp3-tag-remove:empty::before{
		  font-family:"Icons20", sans-serif;
		  font-size:20px;
		  font-style:normal;
		  font-weight:400;
		  line-height:1; }
	.bp3-tag-input{
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  -webkit-box-orient:horizontal;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:row;
			  flex-direction:row;
	  -webkit-box-align:start;
		  -ms-flex-align:start;
			  align-items:flex-start;
	  cursor:text;
	  height:auto;
	  line-height:inherit;
	  min-height:30px;
	  padding-left:5px;
	  padding-right:0; }
	  .bp3-tag-input > *{
		-webkit-box-flex:0;
			-ms-flex-positive:0;
				flex-grow:0;
		-ms-flex-negative:0;
			flex-shrink:0; }
	  .bp3-tag-input > .bp3-tag-input-values{
		-webkit-box-flex:1;
			-ms-flex-positive:1;
				flex-grow:1;
		-ms-flex-negative:1;
			flex-shrink:1; }
	  .bp3-tag-input .bp3-tag-input-icon{
		color:#5c7080;
		margin-left:2px;
		margin-right:7px;
		margin-top:7px; }
	  .bp3-tag-input .bp3-tag-input-values{
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex;
		-webkit-box-orient:horizontal;
		-webkit-box-direction:normal;
			-ms-flex-direction:row;
				flex-direction:row;
		-webkit-box-align:center;
			-ms-flex-align:center;
				align-items:center;
		-ms-flex-item-align:stretch;
			align-self:stretch;
		-ms-flex-wrap:wrap;
			flex-wrap:wrap;
		margin-right:7px;
		margin-top:5px;
		min-width:0; }
		.bp3-tag-input .bp3-tag-input-values > *{
		  -webkit-box-flex:0;
			  -ms-flex-positive:0;
				  flex-grow:0;
		  -ms-flex-negative:0;
			  flex-shrink:0; }
		.bp3-tag-input .bp3-tag-input-values > .bp3-fill{
		  -webkit-box-flex:1;
			  -ms-flex-positive:1;
				  flex-grow:1;
		  -ms-flex-negative:1;
			  flex-shrink:1; }
		.bp3-tag-input .bp3-tag-input-values::before,
		.bp3-tag-input .bp3-tag-input-values > *{
		  margin-right:5px; }
		.bp3-tag-input .bp3-tag-input-values:empty::before,
		.bp3-tag-input .bp3-tag-input-values > :last-child{
		  margin-right:0; }
		.bp3-tag-input .bp3-tag-input-values:first-child .bp3-input-ghost:first-child{
		  padding-left:5px; }
		.bp3-tag-input .bp3-tag-input-values > *{
		  margin-bottom:5px; }
	  .bp3-tag-input .bp3-tag{
		overflow-wrap:break-word; }
		.bp3-tag-input .bp3-tag.bp3-active{
		  outline:rgba(19, 124, 189, 0.6) auto 2px;
		  outline-offset:0;
		  -moz-outline-radius:6px; }
	  .bp3-tag-input .bp3-input-ghost{
		-webkit-box-flex:1;
			-ms-flex:1 1 auto;
				flex:1 1 auto;
		line-height:20px;
		width:80px; }
		.bp3-tag-input .bp3-input-ghost:disabled, .bp3-tag-input .bp3-input-ghost.bp3-disabled{
		  cursor:not-allowed; }
	  .bp3-tag-input .bp3-button,
	  .bp3-tag-input .bp3-spinner{
		margin:3px;
		margin-left:0; }
	  .bp3-tag-input .bp3-button{
		min-height:24px;
		min-width:24px;
		padding:0 7px; }
	  .bp3-tag-input.bp3-large{
		height:auto;
		min-height:40px; }
		.bp3-tag-input.bp3-large::before,
		.bp3-tag-input.bp3-large > *{
		  margin-right:10px; }
		.bp3-tag-input.bp3-large:empty::before,
		.bp3-tag-input.bp3-large > :last-child{
		  margin-right:0; }
		.bp3-tag-input.bp3-large .bp3-tag-input-icon{
		  margin-left:5px;
		  margin-top:10px; }
		.bp3-tag-input.bp3-large .bp3-input-ghost{
		  line-height:30px; }
		.bp3-tag-input.bp3-large .bp3-button{
		  min-height:30px;
		  min-width:30px;
		  padding:5px 10px;
		  margin:5px;
		  margin-left:0; }
		.bp3-tag-input.bp3-large .bp3-spinner{
		  margin:8px;
		  margin-left:0; }
	  .bp3-tag-input.bp3-active{
		background-color:#ffffff;
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-tag-input.bp3-active.bp3-intent-primary{
		  -webkit-box-shadow:0 0 0 1px #106ba3, 0 0 0 3px rgba(16, 107, 163, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #106ba3, 0 0 0 3px rgba(16, 107, 163, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-tag-input.bp3-active.bp3-intent-success{
		  -webkit-box-shadow:0 0 0 1px #0d8050, 0 0 0 3px rgba(13, 128, 80, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #0d8050, 0 0 0 3px rgba(13, 128, 80, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-tag-input.bp3-active.bp3-intent-warning{
		  -webkit-box-shadow:0 0 0 1px #bf7326, 0 0 0 3px rgba(191, 115, 38, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #bf7326, 0 0 0 3px rgba(191, 115, 38, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
		.bp3-tag-input.bp3-active.bp3-intent-danger{
		  -webkit-box-shadow:0 0 0 1px #c23030, 0 0 0 3px rgba(194, 48, 48, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2);
				  box-shadow:0 0 0 1px #c23030, 0 0 0 3px rgba(194, 48, 48, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.2); }
	  .bp3-dark .bp3-tag-input .bp3-tag-input-icon, .bp3-tag-input.bp3-dark .bp3-tag-input-icon{
		color:#a7b6c2; }
	  .bp3-dark .bp3-tag-input .bp3-input-ghost, .bp3-tag-input.bp3-dark .bp3-input-ghost{
		color:#f5f8fa; }
		.bp3-dark .bp3-tag-input .bp3-input-ghost::-webkit-input-placeholder, .bp3-tag-input.bp3-dark .bp3-input-ghost::-webkit-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-tag-input .bp3-input-ghost::-moz-placeholder, .bp3-tag-input.bp3-dark .bp3-input-ghost::-moz-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-tag-input .bp3-input-ghost:-ms-input-placeholder, .bp3-tag-input.bp3-dark .bp3-input-ghost:-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-tag-input .bp3-input-ghost::-ms-input-placeholder, .bp3-tag-input.bp3-dark .bp3-input-ghost::-ms-input-placeholder{
		  color:rgba(167, 182, 194, 0.6); }
		.bp3-dark .bp3-tag-input .bp3-input-ghost::placeholder, .bp3-tag-input.bp3-dark .bp3-input-ghost::placeholder{
		  color:rgba(167, 182, 194, 0.6); }
	  .bp3-dark .bp3-tag-input.bp3-active, .bp3-tag-input.bp3-dark.bp3-active{
		background-color:rgba(16, 22, 26, 0.3);
		-webkit-box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px #137cbd, 0 0 0 1px #137cbd, 0 0 0 3px rgba(19, 124, 189, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-tag-input.bp3-active.bp3-intent-primary, .bp3-tag-input.bp3-dark.bp3-active.bp3-intent-primary{
		  -webkit-box-shadow:0 0 0 1px #106ba3, 0 0 0 3px rgba(16, 107, 163, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #106ba3, 0 0 0 3px rgba(16, 107, 163, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-tag-input.bp3-active.bp3-intent-success, .bp3-tag-input.bp3-dark.bp3-active.bp3-intent-success{
		  -webkit-box-shadow:0 0 0 1px #0d8050, 0 0 0 3px rgba(13, 128, 80, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #0d8050, 0 0 0 3px rgba(13, 128, 80, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-tag-input.bp3-active.bp3-intent-warning, .bp3-tag-input.bp3-dark.bp3-active.bp3-intent-warning{
		  -webkit-box-shadow:0 0 0 1px #bf7326, 0 0 0 3px rgba(191, 115, 38, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #bf7326, 0 0 0 3px rgba(191, 115, 38, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
		.bp3-dark .bp3-tag-input.bp3-active.bp3-intent-danger, .bp3-tag-input.bp3-dark.bp3-active.bp3-intent-danger{
		  -webkit-box-shadow:0 0 0 1px #c23030, 0 0 0 3px rgba(194, 48, 48, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4);
				  box-shadow:0 0 0 1px #c23030, 0 0 0 3px rgba(194, 48, 48, 0.3), inset 0 0 0 1px rgba(16, 22, 26, 0.3), inset 0 1px 1px rgba(16, 22, 26, 0.4); }
	
	.bp3-input-ghost{
	  background:none;
	  border:none;
	  -webkit-box-shadow:none;
			  box-shadow:none;
	  padding:0; }
	  .bp3-input-ghost::-webkit-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input-ghost::-moz-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input-ghost:-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input-ghost::-ms-input-placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input-ghost::placeholder{
		color:rgba(92, 112, 128, 0.6);
		opacity:1; }
	  .bp3-input-ghost:focus{
		outline:none !important; }
	.bp3-toast{
	  -webkit-box-align:start;
		  -ms-flex-align:start;
			  align-items:flex-start;
	  background-color:#ffffff;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  margin:20px 0 0;
	  max-width:500px;
	  min-width:300px;
	  pointer-events:all;
	  position:relative !important; }
	  .bp3-toast.bp3-toast-enter, .bp3-toast.bp3-toast-appear{
		-webkit-transform:translateY(-40px);
				transform:translateY(-40px); }
	  .bp3-toast.bp3-toast-enter-active, .bp3-toast.bp3-toast-appear-active{
		-webkit-transform:translateY(0);
				transform:translateY(0);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	  .bp3-toast.bp3-toast-enter ~ .bp3-toast, .bp3-toast.bp3-toast-appear ~ .bp3-toast{
		-webkit-transform:translateY(-40px);
				transform:translateY(-40px); }
	  .bp3-toast.bp3-toast-enter-active ~ .bp3-toast, .bp3-toast.bp3-toast-appear-active ~ .bp3-toast{
		-webkit-transform:translateY(0);
				transform:translateY(0);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11);
				transition-timing-function:cubic-bezier(0.54, 1.12, 0.38, 1.11); }
	  .bp3-toast.bp3-toast-exit{
		opacity:1;
		-webkit-filter:blur(0);
				filter:blur(0); }
	  .bp3-toast.bp3-toast-exit-active{
		opacity:0;
		-webkit-filter:blur(10px);
				filter:blur(10px);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:300ms;
				transition-duration:300ms;
		-webkit-transition-property:opacity, -webkit-filter;
		transition-property:opacity, -webkit-filter;
		transition-property:opacity, filter;
		transition-property:opacity, filter, -webkit-filter;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-toast.bp3-toast-exit ~ .bp3-toast{
		-webkit-transform:translateY(0);
				transform:translateY(0); }
	  .bp3-toast.bp3-toast-exit-active ~ .bp3-toast{
		-webkit-transform:translateY(-40px);
				transform:translateY(-40px);
		-webkit-transition-delay:50ms;
				transition-delay:50ms;
		-webkit-transition-duration:100ms;
				transition-duration:100ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-toast .bp3-button-group{
		-webkit-box-flex:0;
			-ms-flex:0 0 auto;
				flex:0 0 auto;
		padding:5px;
		padding-left:0; }
	  .bp3-toast > .bp3-icon{
		color:#5c7080;
		margin:12px;
		margin-right:0; }
	  .bp3-toast.bp3-dark,
	  .bp3-dark .bp3-toast{
		background-color:#394b59;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
		.bp3-toast.bp3-dark > .bp3-icon,
		.bp3-dark .bp3-toast > .bp3-icon{
		  color:#a7b6c2; }
	  .bp3-toast[class*="bp3-intent-"] a{
		color:rgba(255, 255, 255, 0.7); }
		.bp3-toast[class*="bp3-intent-"] a:hover{
		  color:#ffffff; }
	  .bp3-toast[class*="bp3-intent-"] > .bp3-icon{
		color:#ffffff; }
	  .bp3-toast[class*="bp3-intent-"] .bp3-button, .bp3-toast[class*="bp3-intent-"] .bp3-button::before,
	  .bp3-toast[class*="bp3-intent-"] .bp3-button .bp3-icon, .bp3-toast[class*="bp3-intent-"] .bp3-button:active{
		color:rgba(255, 255, 255, 0.7) !important; }
	  .bp3-toast[class*="bp3-intent-"] .bp3-button:focus{
		outline-color:rgba(255, 255, 255, 0.5); }
	  .bp3-toast[class*="bp3-intent-"] .bp3-button:hover{
		background-color:rgba(255, 255, 255, 0.15) !important;
		color:#ffffff !important; }
	  .bp3-toast[class*="bp3-intent-"] .bp3-button:active{
		background-color:rgba(255, 255, 255, 0.3) !important;
		color:#ffffff !important; }
	  .bp3-toast[class*="bp3-intent-"] .bp3-button::after{
		background:rgba(255, 255, 255, 0.3) !important; }
	  .bp3-toast.bp3-intent-primary{
		background-color:#137cbd;
		color:#ffffff; }
	  .bp3-toast.bp3-intent-success{
		background-color:#0f9960;
		color:#ffffff; }
	  .bp3-toast.bp3-intent-warning{
		background-color:#d9822b;
		color:#ffffff; }
	  .bp3-toast.bp3-intent-danger{
		background-color:#db3737;
		color:#ffffff; }
	
	.bp3-toast-message{
	  -webkit-box-flex:1;
		  -ms-flex:1 1 auto;
			  flex:1 1 auto;
	  padding:11px;
	  word-break:break-word; }
	
	.bp3-toast-container{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box !important;
	  display:-ms-flexbox !important;
	  display:flex !important;
	  -webkit-box-orient:vertical;
	  -webkit-box-direction:normal;
		  -ms-flex-direction:column;
			  flex-direction:column;
	  left:0;
	  overflow:hidden;
	  padding:0 20px 20px;
	  pointer-events:none;
	  position:fixed;
	  right:0;
	  z-index:40; }
	  .bp3-toast-container.bp3-toast-container-top{
		top:0; }
	  .bp3-toast-container.bp3-toast-container-bottom{
		bottom:0;
		-webkit-box-orient:vertical;
		-webkit-box-direction:reverse;
			-ms-flex-direction:column-reverse;
				flex-direction:column-reverse;
		top:auto; }
	  .bp3-toast-container.bp3-toast-container-left{
		-webkit-box-align:start;
			-ms-flex-align:start;
				align-items:flex-start; }
	  .bp3-toast-container.bp3-toast-container-right{
		-webkit-box-align:end;
			-ms-flex-align:end;
				align-items:flex-end; }
	
	.bp3-toast-container-bottom .bp3-toast.bp3-toast-enter:not(.bp3-toast-enter-active),
	.bp3-toast-container-bottom .bp3-toast.bp3-toast-enter:not(.bp3-toast-enter-active) ~ .bp3-toast, .bp3-toast-container-bottom .bp3-toast.bp3-toast-appear:not(.bp3-toast-appear-active),
	.bp3-toast-container-bottom .bp3-toast.bp3-toast-appear:not(.bp3-toast-appear-active) ~ .bp3-toast,
	.bp3-toast-container-bottom .bp3-toast.bp3-toast-exit-active ~ .bp3-toast,
	.bp3-toast-container-bottom .bp3-toast.bp3-toast-leave-active ~ .bp3-toast{
	  -webkit-transform:translateY(60px);
			  transform:translateY(60px); }
	.bp3-tooltip{
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 2px 4px rgba(16, 22, 26, 0.2), 0 8px 24px rgba(16, 22, 26, 0.2);
	  -webkit-transform:scale(1);
			  transform:scale(1); }
	  .bp3-tooltip .bp3-popover-arrow{
		height:22px;
		position:absolute;
		width:22px; }
		.bp3-tooltip .bp3-popover-arrow::before{
		  height:14px;
		  margin:4px;
		  width:14px; }
	  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-tooltip{
		margin-bottom:11px;
		margin-top:-11px; }
		.bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-tooltip > .bp3-popover-arrow{
		  bottom:-8px; }
		  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-top > .bp3-tooltip > .bp3-popover-arrow svg{
			-webkit-transform:rotate(-90deg);
					transform:rotate(-90deg); }
	  .bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-tooltip{
		margin-left:11px; }
		.bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-tooltip > .bp3-popover-arrow{
		  left:-8px; }
		  .bp3-tether-element-attached-left.bp3-tether-target-attached-right > .bp3-tooltip > .bp3-popover-arrow svg{
			-webkit-transform:rotate(0);
					transform:rotate(0); }
	  .bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-tooltip{
		margin-top:11px; }
		.bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-tooltip > .bp3-popover-arrow{
		  top:-8px; }
		  .bp3-tether-element-attached-top.bp3-tether-target-attached-bottom > .bp3-tooltip > .bp3-popover-arrow svg{
			-webkit-transform:rotate(90deg);
					transform:rotate(90deg); }
	  .bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-tooltip{
		margin-left:-11px;
		margin-right:11px; }
		.bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-tooltip > .bp3-popover-arrow{
		  right:-8px; }
		  .bp3-tether-element-attached-right.bp3-tether-target-attached-left > .bp3-tooltip > .bp3-popover-arrow svg{
			-webkit-transform:rotate(180deg);
					transform:rotate(180deg); }
	  .bp3-tether-element-attached-middle > .bp3-tooltip > .bp3-popover-arrow{
		top:50%;
		-webkit-transform:translateY(-50%);
				transform:translateY(-50%); }
	  .bp3-tether-element-attached-center > .bp3-tooltip > .bp3-popover-arrow{
		right:50%;
		-webkit-transform:translateX(50%);
				transform:translateX(50%); }
	  .bp3-tether-element-attached-top.bp3-tether-target-attached-top > .bp3-tooltip > .bp3-popover-arrow{
		top:-0.22183px; }
	  .bp3-tether-element-attached-right.bp3-tether-target-attached-right > .bp3-tooltip > .bp3-popover-arrow{
		right:-0.22183px; }
	  .bp3-tether-element-attached-left.bp3-tether-target-attached-left > .bp3-tooltip > .bp3-popover-arrow{
		left:-0.22183px; }
	  .bp3-tether-element-attached-bottom.bp3-tether-target-attached-bottom > .bp3-tooltip > .bp3-popover-arrow{
		bottom:-0.22183px; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-left > .bp3-tooltip{
		-webkit-transform-origin:top left;
				transform-origin:top left; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-center > .bp3-tooltip{
		-webkit-transform-origin:top center;
				transform-origin:top center; }
	  .bp3-tether-element-attached-top.bp3-tether-element-attached-right > .bp3-tooltip{
		-webkit-transform-origin:top right;
				transform-origin:top right; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-left > .bp3-tooltip{
		-webkit-transform-origin:center left;
				transform-origin:center left; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-center > .bp3-tooltip{
		-webkit-transform-origin:center center;
				transform-origin:center center; }
	  .bp3-tether-element-attached-middle.bp3-tether-element-attached-right > .bp3-tooltip{
		-webkit-transform-origin:center right;
				transform-origin:center right; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-left > .bp3-tooltip{
		-webkit-transform-origin:bottom left;
				transform-origin:bottom left; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-center > .bp3-tooltip{
		-webkit-transform-origin:bottom center;
				transform-origin:bottom center; }
	  .bp3-tether-element-attached-bottom.bp3-tether-element-attached-right > .bp3-tooltip{
		-webkit-transform-origin:bottom right;
				transform-origin:bottom right; }
	  .bp3-tooltip .bp3-popover-content{
		background:#394b59;
		color:#f5f8fa; }
	  .bp3-tooltip .bp3-popover-arrow::before{
		-webkit-box-shadow:1px 1px 6px rgba(16, 22, 26, 0.2);
				box-shadow:1px 1px 6px rgba(16, 22, 26, 0.2); }
	  .bp3-tooltip .bp3-popover-arrow-border{
		fill:#10161a;
		fill-opacity:0.1; }
	  .bp3-tooltip .bp3-popover-arrow-fill{
		fill:#394b59; }
	  .bp3-popover-enter > .bp3-tooltip, .bp3-popover-appear > .bp3-tooltip{
		-webkit-transform:scale(0.8);
				transform:scale(0.8); }
	  .bp3-popover-enter-active > .bp3-tooltip, .bp3-popover-appear-active > .bp3-tooltip{
		-webkit-transform:scale(1);
				transform:scale(1);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:100ms;
				transition-duration:100ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-popover-exit > .bp3-tooltip{
		-webkit-transform:scale(1);
				transform:scale(1); }
	  .bp3-popover-exit-active > .bp3-tooltip{
		-webkit-transform:scale(0.8);
				transform:scale(0.8);
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:100ms;
				transition-duration:100ms;
		-webkit-transition-property:-webkit-transform;
		transition-property:-webkit-transform;
		transition-property:transform;
		transition-property:transform, -webkit-transform;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-tooltip .bp3-popover-content{
		padding:10px 12px; }
	  .bp3-tooltip.bp3-dark,
	  .bp3-dark .bp3-tooltip{
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 2px 4px rgba(16, 22, 26, 0.4), 0 8px 24px rgba(16, 22, 26, 0.4); }
		.bp3-tooltip.bp3-dark .bp3-popover-content,
		.bp3-dark .bp3-tooltip .bp3-popover-content{
		  background:#e1e8ed;
		  color:#394b59; }
		.bp3-tooltip.bp3-dark .bp3-popover-arrow::before,
		.bp3-dark .bp3-tooltip .bp3-popover-arrow::before{
		  -webkit-box-shadow:1px 1px 6px rgba(16, 22, 26, 0.4);
				  box-shadow:1px 1px 6px rgba(16, 22, 26, 0.4); }
		.bp3-tooltip.bp3-dark .bp3-popover-arrow-border,
		.bp3-dark .bp3-tooltip .bp3-popover-arrow-border{
		  fill:#10161a;
		  fill-opacity:0.2; }
		.bp3-tooltip.bp3-dark .bp3-popover-arrow-fill,
		.bp3-dark .bp3-tooltip .bp3-popover-arrow-fill{
		  fill:#e1e8ed; }
	  .bp3-tooltip.bp3-intent-primary .bp3-popover-content{
		background:#137cbd;
		color:#ffffff; }
	  .bp3-tooltip.bp3-intent-primary .bp3-popover-arrow-fill{
		fill:#137cbd; }
	  .bp3-tooltip.bp3-intent-success .bp3-popover-content{
		background:#0f9960;
		color:#ffffff; }
	  .bp3-tooltip.bp3-intent-success .bp3-popover-arrow-fill{
		fill:#0f9960; }
	  .bp3-tooltip.bp3-intent-warning .bp3-popover-content{
		background:#d9822b;
		color:#ffffff; }
	  .bp3-tooltip.bp3-intent-warning .bp3-popover-arrow-fill{
		fill:#d9822b; }
	  .bp3-tooltip.bp3-intent-danger .bp3-popover-content{
		background:#db3737;
		color:#ffffff; }
	  .bp3-tooltip.bp3-intent-danger .bp3-popover-arrow-fill{
		fill:#db3737; }
	
	.bp3-tooltip-indicator{
	  border-bottom:dotted 1px;
	  cursor:help; }
	.bp3-tree .bp3-icon, .bp3-tree .bp3-icon-standard, .bp3-tree .bp3-icon-large{
	  color:#5c7080; }
	  .bp3-tree .bp3-icon.bp3-intent-primary, .bp3-tree .bp3-icon-standard.bp3-intent-primary, .bp3-tree .bp3-icon-large.bp3-intent-primary{
		color:#137cbd; }
	  .bp3-tree .bp3-icon.bp3-intent-success, .bp3-tree .bp3-icon-standard.bp3-intent-success, .bp3-tree .bp3-icon-large.bp3-intent-success{
		color:#0f9960; }
	  .bp3-tree .bp3-icon.bp3-intent-warning, .bp3-tree .bp3-icon-standard.bp3-intent-warning, .bp3-tree .bp3-icon-large.bp3-intent-warning{
		color:#d9822b; }
	  .bp3-tree .bp3-icon.bp3-intent-danger, .bp3-tree .bp3-icon-standard.bp3-intent-danger, .bp3-tree .bp3-icon-large.bp3-intent-danger{
		color:#db3737; }
	
	.bp3-tree-node-list{
	  list-style:none;
	  margin:0;
	  padding-left:0; }
	
	.bp3-tree-root{
	  background-color:transparent;
	  cursor:default;
	  padding-left:0;
	  position:relative; }
	
	.bp3-tree-node-content-0{
	  padding-left:0px; }
	
	.bp3-tree-node-content-1{
	  padding-left:23px; }
	
	.bp3-tree-node-content-2{
	  padding-left:46px; }
	
	.bp3-tree-node-content-3{
	  padding-left:69px; }
	
	.bp3-tree-node-content-4{
	  padding-left:92px; }
	
	.bp3-tree-node-content-5{
	  padding-left:115px; }
	
	.bp3-tree-node-content-6{
	  padding-left:138px; }
	
	.bp3-tree-node-content-7{
	  padding-left:161px; }
	
	.bp3-tree-node-content-8{
	  padding-left:184px; }
	
	.bp3-tree-node-content-9{
	  padding-left:207px; }
	
	.bp3-tree-node-content-10{
	  padding-left:230px; }
	
	.bp3-tree-node-content-11{
	  padding-left:253px; }
	
	.bp3-tree-node-content-12{
	  padding-left:276px; }
	
	.bp3-tree-node-content-13{
	  padding-left:299px; }
	
	.bp3-tree-node-content-14{
	  padding-left:322px; }
	
	.bp3-tree-node-content-15{
	  padding-left:345px; }
	
	.bp3-tree-node-content-16{
	  padding-left:368px; }
	
	.bp3-tree-node-content-17{
	  padding-left:391px; }
	
	.bp3-tree-node-content-18{
	  padding-left:414px; }
	
	.bp3-tree-node-content-19{
	  padding-left:437px; }
	
	.bp3-tree-node-content-20{
	  padding-left:460px; }
	
	.bp3-tree-node-content{
	  -webkit-box-align:center;
		  -ms-flex-align:center;
			  align-items:center;
	  display:-webkit-box;
	  display:-ms-flexbox;
	  display:flex;
	  height:30px;
	  padding-right:5px;
	  width:100%; }
	  .bp3-tree-node-content:hover{
		background-color:rgba(191, 204, 214, 0.4); }
	
	.bp3-tree-node-caret,
	.bp3-tree-node-caret-none{
	  min-width:30px; }
	
	.bp3-tree-node-caret{
	  color:#5c7080;
	  cursor:pointer;
	  padding:7px;
	  -webkit-transform:rotate(0deg);
			  transform:rotate(0deg);
	  -webkit-transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:-webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9);
	  transition:transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9), -webkit-transform 200ms cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-tree-node-caret:hover{
		color:#182026; }
	  .bp3-dark .bp3-tree-node-caret{
		color:#a7b6c2; }
		.bp3-dark .bp3-tree-node-caret:hover{
		  color:#f5f8fa; }
	  .bp3-tree-node-caret.bp3-tree-node-caret-open{
		-webkit-transform:rotate(90deg);
				transform:rotate(90deg); }
	  .bp3-tree-node-caret.bp3-icon-standard::before{
		content:""; }
	
	.bp3-tree-node-icon{
	  margin-right:7px;
	  position:relative; }
	
	.bp3-tree-node-label{
	  overflow:hidden;
	  text-overflow:ellipsis;
	  white-space:nowrap;
	  word-wrap:normal;
	  -webkit-box-flex:1;
		  -ms-flex:1 1 auto;
			  flex:1 1 auto;
	  position:relative;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-tree-node-label span{
		display:inline; }
	
	.bp3-tree-node-secondary-label{
	  padding:0 5px;
	  -webkit-user-select:none;
		 -moz-user-select:none;
		  -ms-user-select:none;
			  user-select:none; }
	  .bp3-tree-node-secondary-label .bp3-popover-wrapper,
	  .bp3-tree-node-secondary-label .bp3-popover-target{
		-webkit-box-align:center;
			-ms-flex-align:center;
				align-items:center;
		display:-webkit-box;
		display:-ms-flexbox;
		display:flex; }
	
	.bp3-tree-node.bp3-disabled .bp3-tree-node-content{
	  background-color:inherit;
	  color:rgba(92, 112, 128, 0.6);
	  cursor:not-allowed; }
	
	.bp3-tree-node.bp3-disabled .bp3-tree-node-caret,
	.bp3-tree-node.bp3-disabled .bp3-tree-node-icon{
	  color:rgba(92, 112, 128, 0.6);
	  cursor:not-allowed; }
	
	.bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content{
	  background-color:#137cbd; }
	  .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content,
	  .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content .bp3-icon, .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content .bp3-icon-standard, .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content .bp3-icon-large{
		color:#ffffff; }
	  .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content .bp3-tree-node-caret::before{
		color:rgba(255, 255, 255, 0.7); }
	  .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content .bp3-tree-node-caret:hover::before{
		color:#ffffff; }
	
	.bp3-dark .bp3-tree-node-content:hover{
	  background-color:rgba(92, 112, 128, 0.3); }
	
	.bp3-dark .bp3-tree .bp3-icon, .bp3-dark .bp3-tree .bp3-icon-standard, .bp3-dark .bp3-tree .bp3-icon-large{
	  color:#a7b6c2; }
	  .bp3-dark .bp3-tree .bp3-icon.bp3-intent-primary, .bp3-dark .bp3-tree .bp3-icon-standard.bp3-intent-primary, .bp3-dark .bp3-tree .bp3-icon-large.bp3-intent-primary{
		color:#137cbd; }
	  .bp3-dark .bp3-tree .bp3-icon.bp3-intent-success, .bp3-dark .bp3-tree .bp3-icon-standard.bp3-intent-success, .bp3-dark .bp3-tree .bp3-icon-large.bp3-intent-success{
		color:#0f9960; }
	  .bp3-dark .bp3-tree .bp3-icon.bp3-intent-warning, .bp3-dark .bp3-tree .bp3-icon-standard.bp3-intent-warning, .bp3-dark .bp3-tree .bp3-icon-large.bp3-intent-warning{
		color:#d9822b; }
	  .bp3-dark .bp3-tree .bp3-icon.bp3-intent-danger, .bp3-dark .bp3-tree .bp3-icon-standard.bp3-intent-danger, .bp3-dark .bp3-tree .bp3-icon-large.bp3-intent-danger{
		color:#db3737; }
	
	.bp3-dark .bp3-tree-node.bp3-tree-node-selected > .bp3-tree-node-content{
	  background-color:#137cbd; }
	.bp3-omnibar{
	  -webkit-filter:blur(0);
			  filter:blur(0);
	  opacity:1;
	  background-color:#ffffff;
	  border-radius:3px;
	  -webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
			  box-shadow:0 0 0 1px rgba(16, 22, 26, 0.1), 0 4px 8px rgba(16, 22, 26, 0.2), 0 18px 46px 6px rgba(16, 22, 26, 0.2);
	  left:calc(50% - 250px);
	  top:20vh;
	  width:500px;
	  z-index:21; }
	  .bp3-omnibar.bp3-overlay-enter, .bp3-omnibar.bp3-overlay-appear{
		-webkit-filter:blur(20px);
				filter:blur(20px);
		opacity:0.2; }
	  .bp3-omnibar.bp3-overlay-enter-active, .bp3-omnibar.bp3-overlay-appear-active{
		-webkit-filter:blur(0);
				filter:blur(0);
		opacity:1;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:200ms;
				transition-duration:200ms;
		-webkit-transition-property:opacity, -webkit-filter;
		transition-property:opacity, -webkit-filter;
		transition-property:filter, opacity;
		transition-property:filter, opacity, -webkit-filter;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-omnibar.bp3-overlay-exit{
		-webkit-filter:blur(0);
				filter:blur(0);
		opacity:1; }
	  .bp3-omnibar.bp3-overlay-exit-active{
		-webkit-filter:blur(20px);
				filter:blur(20px);
		opacity:0.2;
		-webkit-transition-delay:0;
				transition-delay:0;
		-webkit-transition-duration:200ms;
				transition-duration:200ms;
		-webkit-transition-property:opacity, -webkit-filter;
		transition-property:opacity, -webkit-filter;
		transition-property:filter, opacity;
		transition-property:filter, opacity, -webkit-filter;
		-webkit-transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9);
				transition-timing-function:cubic-bezier(0.4, 1, 0.75, 0.9); }
	  .bp3-omnibar .bp3-input{
		background-color:transparent;
		border-radius:0; }
		.bp3-omnibar .bp3-input, .bp3-omnibar .bp3-input:focus{
		  -webkit-box-shadow:none;
				  box-shadow:none; }
	  .bp3-omnibar .bp3-menu{
		background-color:transparent;
		border-radius:0;
		-webkit-box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.15);
				box-shadow:inset 0 1px 0 rgba(16, 22, 26, 0.15);
		max-height:calc(60vh - 40px);
		overflow:auto; }
		.bp3-omnibar .bp3-menu:empty{
		  display:none; }
	  .bp3-dark .bp3-omnibar, .bp3-omnibar.bp3-dark{
		background-color:#30404d;
		-webkit-box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4);
				box-shadow:0 0 0 1px rgba(16, 22, 26, 0.2), 0 4px 8px rgba(16, 22, 26, 0.4), 0 18px 46px 6px rgba(16, 22, 26, 0.4); }
	
	.bp3-omnibar-overlay .bp3-overlay-backdrop{
	  background-color:rgba(16, 22, 26, 0.2); }
	
	.bp3-select-popover .bp3-popover-content{
	  padding:5px; }
	
	.bp3-select-popover .bp3-input-group{
	  margin-bottom:0; }
	
	.bp3-select-popover .bp3-menu{
	  max-height:300px;
	  max-width:400px;
	  overflow:auto;
	  padding:0; }
	  .bp3-select-popover .bp3-menu:not(:first-child){
		padding-top:5px; }
	
	.bp3-multi-select{
	  min-width:150px; }
	
	.bp3-multi-select-popover .bp3-menu{
	  max-height:300px;
	  max-width:400px;
	  overflow:auto; }
	
	.bp3-select-popover .bp3-popover-content{
	  padding:5px; }
	
	.bp3-select-popover .bp3-input-group{
	  margin-bottom:0; }
	
	.bp3-select-popover .bp3-menu{
	  max-height:300px;
	  max-width:400px;
	  overflow:auto;
	  padding:0; }
	  .bp3-select-popover .bp3-menu:not(:first-child){
		padding-top:5px; }
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/* This file was auto-generated by ensureUiComponents() in @jupyterlab/buildutils */
	
	/**
	 * (DEPRECATED) Support for consuming icons as CSS background images
	 */
	
	/* Icons urls */
	
	:root {
	  --jp-icon-add: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTE5IDEzaC02djZoLTJ2LTZINXYtMmg2VjVoMnY2aDZ2MnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-bug: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTIwIDhoLTIuODFjLS40NS0uNzgtMS4wNy0xLjQ1LTEuODItMS45NkwxNyA0LjQxIDE1LjU5IDNsLTIuMTcgMi4xN0MxMi45NiA1LjA2IDEyLjQ5IDUgMTIgNWMtLjQ5IDAtLjk2LjA2LTEuNDEuMTdMOC40MSAzIDcgNC40MWwxLjYyIDEuNjNDNy44OCA2LjU1IDcuMjYgNy4yMiA2LjgxIDhINHYyaDIuMDljLS4wNS4zMy0uMDkuNjYtLjA5IDF2MUg0djJoMnYxYzAgLjM0LjA0LjY3LjA5IDFINHYyaDIuODFjMS4wNCAxLjc5IDIuOTcgMyA1LjE5IDNzNC4xNS0xLjIxIDUuMTktM0gyMHYtMmgtMi4wOWMuMDUtLjMzLjA5LS42Ni4wOS0xdi0xaDJ2LTJoLTJ2LTFjMC0uMzQtLjA0LS42Ny0uMDktMUgyMFY4em0tNiA4aC00di0yaDR2MnptMC00aC00di0yaDR2MnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-build: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTYiIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTE0LjkgMTcuNDVDMTYuMjUgMTcuNDUgMTcuMzUgMTYuMzUgMTcuMzUgMTVDMTcuMzUgMTMuNjUgMTYuMjUgMTIuNTUgMTQuOSAxMi41NUMxMy41NCAxMi41NSAxMi40NSAxMy42NSAxMi40NSAxNUMxMi40NSAxNi4zNSAxMy41NCAxNy40NSAxNC45IDE3LjQ1Wk0yMC4xIDE1LjY4TDIxLjU4IDE2Ljg0QzIxLjcxIDE2Ljk1IDIxLjc1IDE3LjEzIDIxLjY2IDE3LjI5TDIwLjI2IDE5LjcxQzIwLjE3IDE5Ljg2IDIwIDE5LjkyIDE5LjgzIDE5Ljg2TDE4LjA5IDE5LjE2QzE3LjczIDE5LjQ0IDE3LjMzIDE5LjY3IDE2LjkxIDE5Ljg1TDE2LjY0IDIxLjdDMTYuNjIgMjEuODcgMTYuNDcgMjIgMTYuMyAyMkgxMy41QzEzLjMyIDIyIDEzLjE4IDIxLjg3IDEzLjE1IDIxLjdMMTIuODkgMTkuODVDMTIuNDYgMTkuNjcgMTIuMDcgMTkuNDQgMTEuNzEgMTkuMTZMOS45NjAwMiAxOS44NkM5LjgxMDAyIDE5LjkyIDkuNjIwMDIgMTkuODYgOS41NDAwMiAxOS43MUw4LjE0MDAyIDE3LjI5QzguMDUwMDIgMTcuMTMgOC4wOTAwMiAxNi45NSA4LjIyMDAyIDE2Ljg0TDkuNzAwMDIgMTUuNjhMOS42NTAwMSAxNUw5LjcwMDAyIDE0LjMxTDguMjIwMDIgMTMuMTZDOC4wOTAwMiAxMy4wNSA4LjA1MDAyIDEyLjg2IDguMTQwMDIgMTIuNzFMOS41NDAwMiAxMC4yOUM5LjYyMDAyIDEwLjEzIDkuODEwMDIgMTAuMDcgOS45NjAwMiAxMC4xM0wxMS43MSAxMC44NEMxMi4wNyAxMC41NiAxMi40NiAxMC4zMiAxMi44OSAxMC4xNUwxMy4xNSA4LjI4OTk4QzEzLjE4IDguMTI5OTggMTMuMzIgNy45OTk5OCAxMy41IDcuOTk5OThIMTYuM0MxNi40NyA3Ljk5OTk4IDE2LjYyIDguMTI5OTggMTYuNjQgOC4yODk5OEwxNi45MSAxMC4xNUMxNy4zMyAxMC4zMiAxNy43MyAxMC41NiAxOC4wOSAxMC44NEwxOS44MyAxMC4xM0MyMCAxMC4wNyAyMC4xNyAxMC4xMyAyMC4yNiAxMC4yOUwyMS42NiAxMi43MUMyMS43NSAxMi44NiAyMS43MSAxMy4wNSAyMS41OCAxMy4xNkwyMC4xIDE0LjMxTDIwLjE1IDE1TDIwLjEgMTUuNjhaIi8+CiAgICA8cGF0aCBkPSJNNy4zMjk2NiA3LjQ0NDU0QzguMDgzMSA3LjAwOTU0IDguMzM5MzIgNi4wNTMzMiA3LjkwNDMyIDUuMjk5ODhDNy40NjkzMiA0LjU0NjQzIDYuNTA4MSA0LjI4MTU2IDUuNzU0NjYgNC43MTY1NkM1LjM5MTc2IDQuOTI2MDggNS4xMjY5NSA1LjI3MTE4IDUuMDE4NDkgNS42NzU5NEM0LjkxMDA0IDYuMDgwNzEgNC45NjY4MiA2LjUxMTk4IDUuMTc2MzQgNi44NzQ4OEM1LjYxMTM0IDcuNjI4MzIgNi41NzYyMiA3Ljg3OTU0IDcuMzI5NjYgNy40NDQ1NFpNOS42NTcxOCA0Ljc5NTkzTDEwLjg2NzIgNC45NTE3OUMxMC45NjI4IDQuOTc3NDEgMTEuMDQwMiA1LjA3MTMzIDExLjAzODIgNS4xODc5M0wxMS4wMzg4IDYuOTg4OTNDMTEuMDQ1NSA3LjEwMDU0IDEwLjk2MTYgNy4xOTUxOCAxMC44NTUgNy4yMTA1NEw5LjY2MDAxIDcuMzgwODNMOS4yMzkxNSA4LjEzMTg4TDkuNjY5NjEgOS4yNTc0NUM5LjcwNzI5IDkuMzYyNzEgOS42NjkzNCA5LjQ3Njk5IDkuNTc0MDggOS41MzE5OUw4LjAxNTIzIDEwLjQzMkM3LjkxMTMxIDEwLjQ5MiA3Ljc5MzM3IDEwLjQ2NzcgNy43MjEwNSAxMC4zODI0TDYuOTg3NDggOS40MzE4OEw2LjEwOTMxIDkuNDMwODNMNS4zNDcwNCAxMC4zOTA1QzUuMjg5MDkgMTAuNDcwMiA1LjE3MzgzIDEwLjQ5MDUgNS4wNzE4NyAxMC40MzM5TDMuNTEyNDUgOS41MzI5M0MzLjQxMDQ5IDkuNDc2MzMgMy4zNzY0NyA5LjM1NzQxIDMuNDEwNzUgOS4yNTY3OUwzLjg2MzQ3IDguMTQwOTNMMy42MTc0OSA3Ljc3NDg4TDMuNDIzNDcgNy4zNzg4M0wyLjIzMDc1IDcuMjEyOTdDMi4xMjY0NyA3LjE5MjM1IDIuMDQwNDkgNy4xMDM0MiAyLjA0MjQ1IDYuOTg2ODJMMi4wNDE4NyA1LjE4NTgyQzIuMDQzODMgNS4wNjkyMiAyLjExOTA5IDQuOTc5NTggMi4yMTcwNCA0Ljk2OTIyTDMuNDIwNjUgNC43OTM5M0wzLjg2NzQ5IDQuMDI3ODhMMy40MTEwNSAyLjkxNzMxQzMuMzczMzcgMi44MTIwNCAzLjQxMTMxIDIuNjk3NzYgMy41MTUyMyAyLjYzNzc2TDUuMDc0MDggMS43Mzc3NkM1LjE2OTM0IDEuNjgyNzYgNS4yODcyOSAxLjcwNzA0IDUuMzU5NjEgMS43OTIzMUw2LjExOTE1IDIuNzI3ODhMNi45ODAwMSAyLjczODkzTDcuNzI0OTYgMS43ODkyMkM3Ljc5MTU2IDEuNzA0NTggNy45MTU0OCAxLjY3OTIyIDguMDA4NzkgMS43NDA4Mkw5LjU2ODIxIDIuNjQxODJDOS42NzAxNyAyLjY5ODQyIDkuNzEyODUgMi44MTIzNCA5LjY4NzIzIDIuOTA3OTdMOS4yMTcxOCA0LjAzMzgzTDkuNDYzMTYgNC4zOTk4OEw5LjY1NzE4IDQuNzk1OTNaIi8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-caret-down-empty-thin: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIwIDIwIj4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSIgc2hhcGUtcmVuZGVyaW5nPSJnZW9tZXRyaWNQcmVjaXNpb24iPgoJCTxwb2x5Z29uIGNsYXNzPSJzdDEiIHBvaW50cz0iOS45LDEzLjYgMy42LDcuNCA0LjQsNi42IDkuOSwxMi4yIDE1LjQsNi43IDE2LjEsNy40ICIvPgoJPC9nPgo8L3N2Zz4K);
	  --jp-icon-caret-down-empty: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiIHNoYXBlLXJlbmRlcmluZz0iZ2VvbWV0cmljUHJlY2lzaW9uIj4KICAgIDxwYXRoIGQ9Ik01LjIsNS45TDksOS43bDMuOC0zLjhsMS4yLDEuMmwtNC45LDVsLTQuOS01TDUuMiw1Ljl6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-caret-down: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiIHNoYXBlLXJlbmRlcmluZz0iZ2VvbWV0cmljUHJlY2lzaW9uIj4KICAgIDxwYXRoIGQ9Ik01LjIsNy41TDksMTEuMmwzLjgtMy44SDUuMnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-caret-left: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSIgc2hhcGUtcmVuZGVyaW5nPSJnZW9tZXRyaWNQcmVjaXNpb24iPgoJCTxwYXRoIGQ9Ik0xMC44LDEyLjhMNy4xLDlsMy44LTMuOGwwLDcuNkgxMC44eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-caret-right: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiIHNoYXBlLXJlbmRlcmluZz0iZ2VvbWV0cmljUHJlY2lzaW9uIj4KICAgIDxwYXRoIGQ9Ik03LjIsNS4yTDEwLjksOWwtMy44LDMuOFY1LjJINy4yeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-caret-up-empty-thin: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIwIDIwIj4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSIgc2hhcGUtcmVuZGVyaW5nPSJnZW9tZXRyaWNQcmVjaXNpb24iPgoJCTxwb2x5Z29uIGNsYXNzPSJzdDEiIHBvaW50cz0iMTUuNCwxMy4zIDkuOSw3LjcgNC40LDEzLjIgMy42LDEyLjUgOS45LDYuMyAxNi4xLDEyLjYgIi8+Cgk8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-caret-up: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSIgc2hhcGUtcmVuZGVyaW5nPSJnZW9tZXRyaWNQcmVjaXNpb24iPgoJCTxwYXRoIGQ9Ik01LjIsMTAuNUw5LDYuOGwzLjgsMy44SDUuMnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-case-sensitive: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIwIDIwIj4KICA8ZyBjbGFzcz0ianAtaWNvbjIiIGZpbGw9IiM0MTQxNDEiPgogICAgPHJlY3QgeD0iMiIgeT0iMiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ii8+CiAgPC9nPgogIDxnIGNsYXNzPSJqcC1pY29uLWFjY2VudDIiIGZpbGw9IiNGRkYiPgogICAgPHBhdGggZD0iTTcuNiw4aDAuOWwzLjUsOGgtMS4xTDEwLDE0SDZsLTAuOSwySDRMNy42LDh6IE04LDkuMUw2LjQsMTNoMy4yTDgsOS4xeiIvPgogICAgPHBhdGggZD0iTTE2LjYsOS44Yy0wLjIsMC4xLTAuNCwwLjEtMC43LDAuMWMtMC4yLDAtMC40LTAuMS0wLjYtMC4yYy0wLjEtMC4xLTAuMi0wLjQtMC4yLTAuNyBjLTAuMywwLjMtMC42LDAuNS0wLjksMC43Yy0wLjMsMC4xLTAuNywwLjItMS4xLDAuMmMtMC4zLDAtMC41LDAtMC43LTAuMWMtMC4yLTAuMS0wLjQtMC4yLTAuNi0wLjNjLTAuMi0wLjEtMC4zLTAuMy0wLjQtMC41IGMtMC4xLTAuMi0wLjEtMC40LTAuMS0wLjdjMC0wLjMsMC4xLTAuNiwwLjItMC44YzAuMS0wLjIsMC4zLTAuNCwwLjQtMC41QzEyLDcsMTIuMiw2LjksMTIuNSw2LjhjMC4yLTAuMSwwLjUtMC4xLDAuNy0wLjIgYzAuMy0wLjEsMC41LTAuMSwwLjctMC4xYzAuMiwwLDAuNC0wLjEsMC42LTAuMWMwLjIsMCwwLjMtMC4xLDAuNC0wLjJjMC4xLTAuMSwwLjItMC4yLDAuMi0wLjRjMC0xLTEuMS0xLTEuMy0xIGMtMC40LDAtMS40LDAtMS40LDEuMmgtMC45YzAtMC40LDAuMS0wLjcsMC4yLTFjMC4xLTAuMiwwLjMtMC40LDAuNS0wLjZjMC4yLTAuMiwwLjUtMC4zLDAuOC0wLjNDMTMuMyw0LDEzLjYsNCwxMy45LDQgYzAuMywwLDAuNSwwLDAuOCwwLjFjMC4zLDAsMC41LDAuMSwwLjcsMC4yYzAuMiwwLjEsMC40LDAuMywwLjUsMC41QzE2LDUsMTYsNS4yLDE2LDUuNnYyLjljMCwwLjIsMCwwLjQsMCwwLjUgYzAsMC4xLDAuMSwwLjIsMC4zLDAuMmMwLjEsMCwwLjIsMCwwLjMsMFY5Ljh6IE0xNS4yLDYuOWMtMS4yLDAuNi0zLjEsMC4yLTMuMSwxLjRjMCwxLjQsMy4xLDEsMy4xLTAuNVY2Ljl6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-check: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTkgMTYuMTdMNC44MyAxMmwtMS40MiAxLjQxTDkgMTkgMjEgN2wtMS40MS0xLjQxeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-circle-empty: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEyIDJDNi40NyAyIDIgNi40NyAyIDEyczQuNDcgMTAgMTAgMTAgMTAtNC40NyAxMC0xMFMxNy41MyAyIDEyIDJ6bTAgMThjLTQuNDEgMC04LTMuNTktOC04czMuNTktOCA4LTggOCAzLjU5IDggOC0zLjU5IDgtOCA4eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-circle: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMTggMTgiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPGNpcmNsZSBjeD0iOSIgY3k9IjkiIHI9IjgiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-clear: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8bWFzayBpZD0iZG9udXRIb2xlIj4KICAgIDxyZWN0IHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgZmlsbD0id2hpdGUiIC8+CiAgICA8Y2lyY2xlIGN4PSIxMiIgY3k9IjEyIiByPSI4IiBmaWxsPSJibGFjayIvPgogIDwvbWFzaz4KCiAgPGcgY2xhc3M9ImpwLWljb24zIiBmaWxsPSIjNjE2MTYxIj4KICAgIDxyZWN0IGhlaWdodD0iMTgiIHdpZHRoPSIyIiB4PSIxMSIgeT0iMyIgdHJhbnNmb3JtPSJyb3RhdGUoMzE1LCAxMiwgMTIpIi8+CiAgICA8Y2lyY2xlIGN4PSIxMiIgY3k9IjEyIiByPSIxMCIgbWFzaz0idXJsKCNkb251dEhvbGUpIi8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-close: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbi1ub25lIGpwLWljb24tc2VsZWN0YWJsZS1pbnZlcnNlIGpwLWljb24zLWhvdmVyIiBmaWxsPSJub25lIj4KICAgIDxjaXJjbGUgY3g9IjEyIiBjeT0iMTIiIHI9IjExIi8+CiAgPC9nPgoKICA8ZyBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIGpwLWljb24tYWNjZW50Mi1ob3ZlciIgZmlsbD0iIzYxNjE2MSI+CiAgICA8cGF0aCBkPSJNMTkgNi40MUwxNy41OSA1IDEyIDEwLjU5IDYuNDEgNSA1IDYuNDEgMTAuNTkgMTIgNSAxNy41OSA2LjQxIDE5IDEyIDEzLjQxIDE3LjU5IDE5IDE5IDE3LjU5IDEzLjQxIDEyeiIvPgogIDwvZz4KCiAgPGcgY2xhc3M9ImpwLWljb24tbm9uZSBqcC1pY29uLWJ1c3kiIGZpbGw9Im5vbmUiPgogICAgPGNpcmNsZSBjeD0iMTIiIGN5PSIxMiIgcj0iNyIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-code: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjIiIGhlaWdodD0iMjIiIHZpZXdCb3g9IjAgMCAyOCAyOCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CgkJPHBhdGggZD0iTTExLjQgMTguNkw2LjggMTRMMTEuNCA5LjRMMTAgOEw0IDE0TDEwIDIwTDExLjQgMTguNlpNMTYuNiAxOC42TDIxLjIgMTRMMTYuNiA5LjRMMTggOEwyNCAxNEwxOCAyMEwxNi42IDE4LjZWMTguNloiLz4KCTwvZz4KPC9zdmc+Cg==);
	  --jp-icon-console: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIwMCAyMDAiPgogIDxnIGNsYXNzPSJqcC1pY29uLWJyYW5kMSBqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiMwMjg4RDEiPgogICAgPHBhdGggZD0iTTIwIDE5LjhoMTYwdjE1OS45SDIweiIvPgogIDwvZz4KICA8ZyBjbGFzcz0ianAtaWNvbi1zZWxlY3RhYmxlLWludmVyc2UiIGZpbGw9IiNmZmYiPgogICAgPHBhdGggZD0iTTEwNSAxMjcuM2g0MHYxMi44aC00MHpNNTEuMSA3N0w3NCA5OS45bC0yMy4zIDIzLjMgMTAuNSAxMC41IDIzLjMtMjMuM0w5NSA5OS45IDg0LjUgODkuNCA2MS42IDY2LjV6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-copy: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMTggMTgiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTExLjksMUgzLjJDMi40LDEsMS43LDEuNywxLjcsMi41djEwLjJoMS41VjIuNWg4LjdWMXogTTE0LjEsMy45aC04Yy0wLjgsMC0xLjUsMC43LTEuNSwxLjV2MTAuMmMwLDAuOCwwLjcsMS41LDEuNSwxLjVoOCBjMC44LDAsMS41LTAuNywxLjUtMS41VjUuNEMxNS41LDQuNiwxNC45LDMuOSwxNC4xLDMuOXogTTE0LjEsMTUuNWgtOFY1LjRoOFYxNS41eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-cut: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTkuNjQgNy42NGMuMjMtLjUuMzYtMS4wNS4zNi0xLjY0IDAtMi4yMS0xLjc5LTQtNC00UzIgMy43OSAyIDZzMS43OSA0IDQgNGMuNTkgMCAxLjE0LS4xMyAxLjY0LS4zNkwxMCAxMmwtMi4zNiAyLjM2QzcuMTQgMTQuMTMgNi41OSAxNCA2IDE0Yy0yLjIxIDAtNCAxLjc5LTQgNHMxLjc5IDQgNCA0IDQtMS43OSA0LTRjMC0uNTktLjEzLTEuMTQtLjM2LTEuNjRMMTIgMTRsNyA3aDN2LTFMOS42NCA3LjY0ek02IDhjLTEuMSAwLTItLjg5LTItMnMuOS0yIDItMiAyIC44OSAyIDItLjkgMi0yIDJ6bTAgMTJjLTEuMSAwLTItLjg5LTItMnMuOS0yIDItMiAyIC44OSAyIDItLjkgMi0yIDJ6bTYtNy41Yy0uMjggMC0uNS0uMjItLjUtLjVzLjIyLS41LjUtLjUuNS4yMi41LjUtLjIyLjUtLjUuNXpNMTkgM2wtNiA2IDIgMiA3LTdWM3oiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-download: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTE5IDloLTRWM0g5djZINWw3IDcgNy03ek01IDE4djJoMTR2LTJINXoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-edit: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTMgMTcuMjVWMjFoMy43NUwxNy44MSA5Ljk0bC0zLjc1LTMuNzVMMyAxNy4yNXpNMjAuNzEgNy4wNGMuMzktLjM5LjM5LTEuMDIgMC0xLjQxbC0yLjM0LTIuMzRjLS4zOS0uMzktMS4wMi0uMzktMS40MSAwbC0xLjgzIDEuODMgMy43NSAzLjc1IDEuODMtMS44M3oiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-ellipses: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPGNpcmNsZSBjeD0iNSIgY3k9IjEyIiByPSIyIi8+CiAgICA8Y2lyY2xlIGN4PSIxMiIgY3k9IjEyIiByPSIyIi8+CiAgICA8Y2lyY2xlIGN4PSIxOSIgY3k9IjEyIiByPSIyIi8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-extension: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTIwLjUgMTFIMTlWN2MwLTEuMS0uOS0yLTItMmgtNFYzLjVDMTMgMi4xMiAxMS44OCAxIDEwLjUgMVM4IDIuMTIgOCAzLjVWNUg0Yy0xLjEgMC0xLjk5LjktMS45OSAydjMuOEgzLjVjMS40OSAwIDIuNyAxLjIxIDIuNyAyLjdzLTEuMjEgMi43LTIuNyAyLjdIMlYyMGMwIDEuMS45IDIgMiAyaDMuOHYtMS41YzAtMS40OSAxLjIxLTIuNyAyLjctMi43IDEuNDkgMCAyLjcgMS4yMSAyLjcgMi43VjIySDE3YzEuMSAwIDItLjkgMi0ydi00aDEuNWMxLjM4IDAgMi41LTEuMTIgMi41LTIuNVMyMS44OCAxMSAyMC41IDExeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-fast-forward: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTQgMThsOC41LTZMNCA2djEyem05LTEydjEybDguNS02TDEzIDZ6Ii8+CiAgICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-file-upload: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTkgMTZoNnYtNmg0bC03LTctNyA3aDR6bS00IDJoMTR2Mkg1eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-file: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMTkuMyA4LjJsLTUuNS01LjVjLS4zLS4zLS43LS41LTEuMi0uNUgzLjljLS44LjEtMS42LjktMS42IDEuOHYxNC4xYzAgLjkuNyAxLjYgMS42IDEuNmgxNC4yYy45IDAgMS42LS43IDEuNi0xLjZWOS40Yy4xLS41LS4xLS45LS40LTEuMnptLTUuOC0zLjNsMy40IDMuNmgtMy40VjQuOXptMy45IDEyLjdINC43Yy0uMSAwLS4yIDAtLjItLjJWNC43YzAtLjIuMS0uMy4yLS4zaDcuMnY0LjRzMCAuOC4zIDEuMWMuMy4zIDEuMS4zIDEuMS4zaDQuM3Y3LjJzLS4xLjItLjIuMnoiLz4KPC9zdmc+Cg==);
	  --jp-icon-filter-list: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEwIDE4aDR2LTJoLTR2MnpNMyA2djJoMThWNkgzem0zIDdoMTJ2LTJINnYyeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-folder: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMTAgNEg0Yy0xLjEgMC0xLjk5LjktMS45OSAyTDIgMThjMCAxLjEuOSAyIDIgMmgxNmMxLjEgMCAyLS45IDItMlY4YzAtMS4xLS45LTItMi0yaC04bC0yLTJ6Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-html5: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDUxMiA1MTIiPgogIDxwYXRoIGNsYXNzPSJqcC1pY29uMCBqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiMwMDAiIGQ9Ik0xMDguNCAwaDIzdjIyLjhoMjEuMlYwaDIzdjY5aC0yM1Y0NmgtMjF2MjNoLTIzLjJNMjA2IDIzaC0yMC4zVjBoNjMuN3YyM0gyMjl2NDZoLTIzbTUzLjUtNjloMjQuMWwxNC44IDI0LjNMMzEzLjIgMGgyNC4xdjY5aC0yM1YzNC44bC0xNi4xIDI0LjgtMTYuMS0yNC44VjY5aC0yMi42bTg5LjItNjloMjN2NDYuMmgzMi42VjY5aC01NS42Ii8+CiAgPHBhdGggY2xhc3M9ImpwLWljb24tc2VsZWN0YWJsZSIgZmlsbD0iI2U0NGQyNiIgZD0iTTEwNy42IDQ3MWwtMzMtMzcwLjRoMzYyLjhsLTMzIDM3MC4yTDI1NS43IDUxMiIvPgogIDxwYXRoIGNsYXNzPSJqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiNmMTY1MjkiIGQ9Ik0yNTYgNDgwLjVWMTMxaDE0OC4zTDM3NiA0NDciLz4KICA8cGF0aCBjbGFzcz0ianAtaWNvbi1zZWxlY3RhYmxlLWludmVyc2UiIGZpbGw9IiNlYmViZWIiIGQ9Ik0xNDIgMTc2LjNoMTE0djQ1LjRoLTY0LjJsNC4yIDQ2LjVoNjB2NDUuM0gxNTQuNG0yIDIyLjhIMjAybDMuMiAzNi4zIDUwLjggMTMuNnY0Ny40bC05My4yLTI2Ii8+CiAgPHBhdGggY2xhc3M9ImpwLWljb24tc2VsZWN0YWJsZS1pbnZlcnNlIiBmaWxsPSIjZmZmIiBkPSJNMzY5LjYgMTc2LjNIMjU1Ljh2NDUuNGgxMDkuNm0tNC4xIDQ2LjVIMjU1Ljh2NDUuNGg1NmwtNS4zIDU5LTUwLjcgMTMuNnY0Ny4ybDkzLTI1LjgiLz4KPC9zdmc+Cg==);
	  --jp-icon-image: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8cGF0aCBjbGFzcz0ianAtaWNvbi1icmFuZDQganAtaWNvbi1zZWxlY3RhYmxlLWludmVyc2UiIGZpbGw9IiNGRkYiIGQ9Ik0yLjIgMi4yaDE3LjV2MTcuNUgyLjJ6Ii8+CiAgPHBhdGggY2xhc3M9ImpwLWljb24tYnJhbmQwIGpwLWljb24tc2VsZWN0YWJsZSIgZmlsbD0iIzNGNTFCNSIgZD0iTTIuMiAyLjJ2MTcuNWgxNy41bC4xLTE3LjVIMi4yem0xMi4xIDIuMmMxLjIgMCAyLjIgMSAyLjIgMi4ycy0xIDIuMi0yLjIgMi4yLTIuMi0xLTIuMi0yLjIgMS0yLjIgMi4yLTIuMnpNNC40IDE3LjZsMy4zLTguOCAzLjMgNi42IDIuMi0zLjIgNC40IDUuNEg0LjR6Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-inspector: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMjAgNEg0Yy0xLjEgMC0xLjk5LjktMS45OSAyTDIgMThjMCAxLjEuOSAyIDIgMmgxNmMxLjEgMCAyLS45IDItMlY2YzAtMS4xLS45LTItMi0yem0tNSAxNEg0di00aDExdjR6bTAtNUg0VjloMTF2NHptNSA1aC00VjloNHY5eiIvPgo8L3N2Zz4K);
	  --jp-icon-json: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8ZyBjbGFzcz0ianAtaWNvbi13YXJuMSBqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiNGOUE4MjUiPgogICAgPHBhdGggZD0iTTIwLjIgMTEuOGMtMS42IDAtMS43LjUtMS43IDEgMCAuNC4xLjkuMSAxLjMuMS41LjEuOS4xIDEuMyAwIDEuNy0xLjQgMi4zLTMuNSAyLjNoLS45di0xLjloLjVjMS4xIDAgMS40IDAgMS40LS44IDAtLjMgMC0uNi0uMS0xIDAtLjQtLjEtLjgtLjEtMS4yIDAtMS4zIDAtMS44IDEuMy0yLTEuMy0uMi0xLjMtLjctMS4zLTIgMC0uNC4xLS44LjEtMS4yLjEtLjQuMS0uNy4xLTEgMC0uOC0uNC0uNy0xLjQtLjhoLS41VjQuMWguOWMyLjIgMCAzLjUuNyAzLjUgMi4zIDAgLjQtLjEuOS0uMSAxLjMtLjEuNS0uMS45LS4xIDEuMyAwIC41LjIgMSAxLjcgMXYxLjh6TTEuOCAxMC4xYzEuNiAwIDEuNy0uNSAxLjctMSAwLS40LS4xLS45LS4xLTEuMy0uMS0uNS0uMS0uOS0uMS0xLjMgMC0xLjYgMS40LTIuMyAzLjUtMi4zaC45djEuOWgtLjVjLTEgMC0xLjQgMC0xLjQuOCAwIC4zIDAgLjYuMSAxIDAgLjIuMS42LjEgMSAwIDEuMyAwIDEuOC0xLjMgMkM2IDExLjIgNiAxMS43IDYgMTNjMCAuNC0uMS44LS4xIDEuMi0uMS4zLS4xLjctLjEgMSAwIC44LjMuOCAxLjQuOGguNXYxLjloLS45Yy0yLjEgMC0zLjUtLjYtMy41LTIuMyAwLS40LjEtLjkuMS0xLjMuMS0uNS4xLS45LjEtMS4zIDAtLjUtLjItMS0xLjctMXYtMS45eiIvPgogICAgPGNpcmNsZSBjeD0iMTEiIGN5PSIxMy44IiByPSIyLjEiLz4KICAgIDxjaXJjbGUgY3g9IjExIiBjeT0iOC4yIiByPSIyLjEiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-jupyter-favicon: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTUyIiBoZWlnaHQ9IjE2NSIgdmlld0JveD0iMCAwIDE1MiAxNjUiIHZlcnNpb249IjEuMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbi13YXJuMCIgZmlsbD0iI0YzNzcyNiI+CiAgICA8cGF0aCB0cmFuc2Zvcm09InRyYW5zbGF0ZSgwLjA3ODk0NywgMTEwLjU4MjkyNykiIGQ9Ik03NS45NDIyODQyLDI5LjU4MDQ1NjEgQzQzLjMwMjM5NDcsMjkuNTgwNDU2MSAxNC43OTY3ODMyLDE3LjY1MzQ2MzQgMCwwIEM1LjUxMDgzMjExLDE1Ljg0MDY4MjkgMTUuNzgxNTM4OSwyOS41NjY3NzMyIDI5LjM5MDQ5NDcsMzkuMjc4NDE3MSBDNDIuOTk5Nyw0OC45ODk4NTM3IDU5LjI3MzcsNTQuMjA2NzgwNSA3NS45NjA1Nzg5LDU0LjIwNjc4MDUgQzkyLjY0NzQ1NzksNTQuMjA2NzgwNSAxMDguOTIxNDU4LDQ4Ljk4OTg1MzcgMTIyLjUzMDY2MywzOS4yNzg0MTcxIEMxMzYuMTM5NDUzLDI5LjU2Njc3MzIgMTQ2LjQxMDI4NCwxNS44NDA2ODI5IDE1MS45MjExNTgsMCBDMTM3LjA4Nzg2OCwxNy42NTM0NjM0IDEwOC41ODI1ODksMjkuNTgwNDU2MSA3NS45NDIyODQyLDI5LjU4MDQ1NjEgTDc1Ljk0MjI4NDIsMjkuNTgwNDU2MSBaIiAvPgogICAgPHBhdGggdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMC4wMzczNjgsIDAuNzA0ODc4KSIgZD0iTTc1Ljk3ODQ1NzksMjQuNjI2NDA3MyBDMTA4LjYxODc2MywyNC42MjY0MDczIDEzNy4xMjQ0NTgsMzYuNTUzNDQxNSAxNTEuOTIxMTU4LDU0LjIwNjc4MDUgQzE0Ni40MTAyODQsMzguMzY2MjIyIDEzNi4xMzk0NTMsMjQuNjQwMTMxNyAxMjIuNTMwNjYzLDE0LjkyODQ4NzggQzEwOC45MjE0NTgsNS4yMTY4NDM5IDkyLjY0NzQ1NzksMCA3NS45NjA1Nzg5LDAgQzU5LjI3MzcsMCA0Mi45OTk3LDUuMjE2ODQzOSAyOS4zOTA0OTQ3LDE0LjkyODQ4NzggQzE1Ljc4MTUzODksMjQuNjQwMTMxNyA1LjUxMDgzMjExLDM4LjM2NjIyMiAwLDU0LjIwNjc4MDUgQzE0LjgzMzA4MTYsMzYuNTg5OTI5MyA0My4zMzg1Njg0LDI0LjYyNjQwNzMgNzUuOTc4NDU3OSwyNC42MjY0MDczIEw3NS45Nzg0NTc5LDI0LjYyNjQwNzMgWiIgLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-jupyter: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMzkiIGhlaWdodD0iNTEiIHZpZXdCb3g9IjAgMCAzOSA1MSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtMTYzOCAtMjI4MSkiPgogICAgPGcgY2xhc3M9ImpwLWljb24td2FybjAiIGZpbGw9IiNGMzc3MjYiPgogICAgICA8cGF0aCB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxNjM5Ljc0IDIzMTEuOTgpIiBkPSJNIDE4LjI2NDYgNy4xMzQxMUMgMTAuNDE0NSA3LjEzNDExIDMuNTU4NzIgNC4yNTc2IDAgMEMgMS4zMjUzOSAzLjgyMDQgMy43OTU1NiA3LjEzMDgxIDcuMDY4NiA5LjQ3MzAzQyAxMC4zNDE3IDExLjgxNTIgMTQuMjU1NyAxMy4wNzM0IDE4LjI2OSAxMy4wNzM0QyAyMi4yODIzIDEzLjA3MzQgMjYuMTk2MyAxMS44MTUyIDI5LjQ2OTQgOS40NzMwM0MgMzIuNzQyNCA3LjEzMDgxIDM1LjIxMjYgMy44MjA0IDM2LjUzOCAwQyAzMi45NzA1IDQuMjU3NiAyNi4xMTQ4IDcuMTM0MTEgMTguMjY0NiA3LjEzNDExWiIvPgogICAgICA8cGF0aCB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxNjM5LjczIDIyODUuNDgpIiBkPSJNIDE4LjI3MzMgNS45MzkzMUMgMjYuMTIzNSA1LjkzOTMxIDMyLjk3OTMgOC44MTU4MyAzNi41MzggMTMuMDczNEMgMzUuMjEyNiA5LjI1MzAzIDMyLjc0MjQgNS45NDI2MiAyOS40Njk0IDMuNjAwNEMgMjYuMTk2MyAxLjI1ODE4IDIyLjI4MjMgMCAxOC4yNjkgMEMgMTQuMjU1NyAwIDEwLjM0MTcgMS4yNTgxOCA3LjA2ODYgMy42MDA0QyAzLjc5NTU2IDUuOTQyNjIgMS4zMjUzOSA5LjI1MzAzIDAgMTMuMDczNEMgMy41Njc0NSA4LjgyNDYzIDEwLjQyMzIgNS45MzkzMSAxOC4yNzMzIDUuOTM5MzFaIi8+CiAgICA8L2c+CiAgICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgICA8cGF0aCB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxNjY5LjMgMjI4MS4zMSkiIGQ9Ik0gNS44OTM1MyAyLjg0NEMgNS45MTg4OSAzLjQzMTY1IDUuNzcwODUgNC4wMTM2NyA1LjQ2ODE1IDQuNTE2NDVDIDUuMTY1NDUgNS4wMTkyMiA0LjcyMTY4IDUuNDIwMTUgNC4xOTI5OSA1LjY2ODUxQyAzLjY2NDMgNS45MTY4OCAzLjA3NDQ0IDYuMDAxNTEgMi40OTgwNSA1LjkxMTcxQyAxLjkyMTY2IDUuODIxOSAxLjM4NDYzIDUuNTYxNyAwLjk1NDg5OCA1LjE2NDAxQyAwLjUyNTE3IDQuNzY2MzMgMC4yMjIwNTYgNC4yNDkwMyAwLjA4MzkwMzcgMy42Nzc1N0MgLTAuMDU0MjQ4MyAzLjEwNjExIC0wLjAyMTIzIDIuNTA2MTcgMC4xNzg3ODEgMS45NTM2NEMgMC4zNzg3OTMgMS40MDExIDAuNzM2ODA5IDAuOTIwODE3IDEuMjA3NTQgMC41NzM1MzhDIDEuNjc4MjYgMC4yMjYyNTkgMi4yNDA1NSAwLjAyNzU5MTkgMi44MjMyNiAwLjAwMjY3MjI5QyAzLjYwMzg5IC0wLjAzMDcxMTUgNC4zNjU3MyAwLjI0OTc4OSA0Ljk0MTQyIDAuNzgyNTUxQyA1LjUxNzExIDEuMzE1MzEgNS44NTk1NiAyLjA1Njc2IDUuODkzNTMgMi44NDRaIi8+CiAgICAgIDxwYXRoIHRyYW5zZm9ybT0idHJhbnNsYXRlKDE2MzkuOCAyMzIzLjgxKSIgZD0iTSA3LjQyNzg5IDMuNTgzMzhDIDcuNDYwMDggNC4zMjQzIDcuMjczNTUgNS4wNTgxOSA2Ljg5MTkzIDUuNjkyMTNDIDYuNTEwMzEgNi4zMjYwNyA1Ljk1MDc1IDYuODMxNTYgNS4yODQxMSA3LjE0NDZDIDQuNjE3NDcgNy40NTc2MyAzLjg3MzcxIDcuNTY0MTQgMy4xNDcwMiA3LjQ1MDYzQyAyLjQyMDMyIDcuMzM3MTIgMS43NDMzNiA3LjAwODcgMS4yMDE4NCA2LjUwNjk1QyAwLjY2MDMyOCA2LjAwNTIgMC4yNzg2MSA1LjM1MjY4IDAuMTA1MDE3IDQuNjMyMDJDIC0wLjA2ODU3NTcgMy45MTEzNSAtMC4wMjYyMzYxIDMuMTU0OTQgMC4yMjY2NzUgMi40NTg1NkMgMC40Nzk1ODcgMS43NjIxNyAwLjkzMTY5NyAxLjE1NzEzIDEuNTI1NzYgMC43MjAwMzNDIDIuMTE5ODMgMC4yODI5MzUgMi44MjkxNCAwLjAzMzQzOTUgMy41NjM4OSAwLjAwMzEzMzQ0QyA0LjU0NjY3IC0wLjAzNzQwMzMgNS41MDUyOSAwLjMxNjcwNiA2LjIyOTYxIDAuOTg3ODM1QyA2Ljk1MzkzIDEuNjU4OTYgNy4zODQ4NCAyLjU5MjM1IDcuNDI3ODkgMy41ODMzOEwgNy40Mjc4OSAzLjU4MzM4WiIvPgogICAgICA8cGF0aCB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxNjM4LjM2IDIyODYuMDYpIiBkPSJNIDIuMjc0NzEgNC4zOTYyOUMgMS44NDM2MyA0LjQxNTA4IDEuNDE2NzEgNC4zMDQ0NSAxLjA0Nzk5IDQuMDc4NDNDIDAuNjc5MjY4IDMuODUyNCAwLjM4NTMyOCAzLjUyMTE0IDAuMjAzMzcxIDMuMTI2NTZDIDAuMDIxNDEzNiAyLjczMTk4IC0wLjA0MDM3OTggMi4yOTE4MyAwLjAyNTgxMTYgMS44NjE4MUMgMC4wOTIwMDMxIDEuNDMxOCAwLjI4MzIwNCAxLjAzMTI2IDAuNTc1MjEzIDAuNzEwODgzQyAwLjg2NzIyMiAwLjM5MDUxIDEuMjQ2OTEgMC4xNjQ3MDggMS42NjYyMiAwLjA2MjA1OTJDIDIuMDg1NTMgLTAuMDQwNTg5NyAyLjUyNTYxIC0wLjAxNTQ3MTQgMi45MzA3NiAwLjEzNDIzNUMgMy4zMzU5MSAwLjI4Mzk0MSAzLjY4NzkyIDAuNTUxNTA1IDMuOTQyMjIgMC45MDMwNkMgNC4xOTY1MiAxLjI1NDYyIDQuMzQxNjkgMS42NzQzNiA0LjM1OTM1IDIuMTA5MTZDIDQuMzgyOTkgMi42OTEwNyA0LjE3Njc4IDMuMjU4NjkgMy43ODU5NyAzLjY4NzQ2QyAzLjM5NTE2IDQuMTE2MjQgMi44NTE2NiA0LjM3MTE2IDIuMjc0NzEgNC4zOTYyOUwgMi4yNzQ3MSA0LjM5NjI5WiIvPgogICAgPC9nPgogIDwvZz4+Cjwvc3ZnPgo=);
	  --jp-icon-jupyterlab-wordmark: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyMDAiIHZpZXdCb3g9IjAgMCAxODYwLjggNDc1Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjIiIGZpbGw9IiM0RTRFNEUiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDQ4MC4xMzY0MDEsIDY0LjI3MTQ5MykiPgogICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMC4wMDAwMDAsIDU4Ljg3NTU2NikiPgogICAgICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgwLjA4NzYwMywgMC4xNDAyOTQpIj4KICAgICAgICA8cGF0aCBkPSJNLTQyNi45LDE2OS44YzAsNDguNy0zLjcsNjQuNy0xMy42LDc2LjRjLTEwLjgsMTAtMjUsMTUuNS0zOS43LDE1LjVsMy43LDI5IGMyMi44LDAuMyw0NC44LTcuOSw2MS45LTIzLjFjMTcuOC0xOC41LDI0LTQ0LjEsMjQtODMuM1YwSC00Mjd2MTcwLjFMLTQyNi45LDE2OS44TC00MjYuOSwxNjkuOHoiLz4KICAgICAgPC9nPgogICAgPC9nPgogICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMTU1LjA0NTI5NiwgNTYuODM3MTA0KSI+CiAgICAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDEuNTYyNDUzLCAxLjc5OTg0MikiPgogICAgICAgIDxwYXRoIGQ9Ik0tMzEyLDE0OGMwLDIxLDAsMzkuNSwxLjcsNTUuNGgtMzEuOGwtMi4xLTMzLjNoLTAuOGMtNi43LDExLjYtMTYuNCwyMS4zLTI4LDI3LjkgYy0xMS42LDYuNi0yNC44LDEwLTM4LjIsOS44Yy0zMS40LDAtNjktMTcuNy02OS04OVYwaDM2LjR2MTEyLjdjMCwzOC43LDExLjYsNjQuNyw0NC42LDY0LjdjMTAuMy0wLjIsMjAuNC0zLjUsMjguOS05LjQgYzguNS01LjksMTUuMS0xNC4zLDE4LjktMjMuOWMyLjItNi4xLDMuMy0xMi41LDMuMy0xOC45VjAuMmgzNi40VjE0OEgtMzEyTC0zMTIsMTQ4eiIvPgogICAgICA8L2c+CiAgICA8L2c+CiAgICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgzOTAuMDEzMzIyLCA1My40Nzk2MzgpIj4KICAgICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMS43MDY0NTgsIDAuMjMxNDI1KSI+CiAgICAgICAgPHBhdGggZD0iTS00NzguNiw3MS40YzAtMjYtMC44LTQ3LTEuNy02Ni43aDMyLjdsMS43LDM0LjhoMC44YzcuMS0xMi41LDE3LjUtMjIuOCwzMC4xLTI5LjcgYzEyLjUtNywyNi43LTEwLjMsNDEtOS44YzQ4LjMsMCw4NC43LDQxLjcsODQuNywxMDMuM2MwLDczLjEtNDMuNywxMDkuMi05MSwxMDkuMmMtMTIuMSwwLjUtMjQuMi0yLjItMzUtNy44IGMtMTAuOC01LjYtMTkuOS0xMy45LTI2LjYtMjQuMmgtMC44VjI5MWgtMzZ2LTIyMEwtNDc4LjYsNzEuNEwtNDc4LjYsNzEuNHogTS00NDIuNiwxMjUuNmMwLjEsNS4xLDAuNiwxMC4xLDEuNywxNS4xIGMzLDEyLjMsOS45LDIzLjMsMTkuOCwzMS4xYzkuOSw3LjgsMjIuMSwxMi4xLDM0LjcsMTIuMWMzOC41LDAsNjAuNy0zMS45LDYwLjctNzguNWMwLTQwLjctMjEuMS03NS42LTU5LjUtNzUuNiBjLTEyLjksMC40LTI1LjMsNS4xLTM1LjMsMTMuNGMtOS45LDguMy0xNi45LDE5LjctMTkuNiwzMi40Yy0xLjUsNC45LTIuMywxMC0yLjUsMTUuMVYxMjUuNkwtNDQyLjYsMTI1LjZMLTQ0Mi42LDEyNS42eiIvPgogICAgICA8L2c+CiAgICA8L2c+CiAgICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSg2MDYuNzQwNzI2LCA1Ni44MzcxMDQpIj4KICAgICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMC43NTEyMjYsIDEuOTg5Mjk5KSI+CiAgICAgICAgPHBhdGggZD0iTS00NDAuOCwwbDQzLjcsMTIwLjFjNC41LDEzLjQsOS41LDI5LjQsMTIuOCw0MS43aDAuOGMzLjctMTIuMiw3LjktMjcuNywxMi44LTQyLjQgbDM5LjctMTE5LjJoMzguNUwtMzQ2LjksMTQ1Yy0yNiw2OS43LTQzLjcsMTA1LjQtNjguNiwxMjcuMmMtMTIuNSwxMS43LTI3LjksMjAtNDQuNiwyMy45bC05LjEtMzEuMSBjMTEuNy0zLjksMjIuNS0xMC4xLDMxLjgtMTguMWMxMy4yLTExLjEsMjMuNy0yNS4yLDMwLjYtNDEuMmMxLjUtMi44LDIuNS01LjcsMi45LTguOGMtMC4zLTMuMy0xLjItNi42LTIuNS05LjdMLTQ4MC4yLDAuMSBoMzkuN0wtNDQwLjgsMEwtNDQwLjgsMHoiLz4KICAgICAgPC9nPgogICAgPC9nPgogICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoODIyLjc0ODEwNCwgMC4wMDAwMDApIj4KICAgICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoMS40NjQwNTAsIDAuMzc4OTE0KSI+CiAgICAgICAgPHBhdGggZD0iTS00MTMuNywwdjU4LjNoNTJ2MjguMmgtNTJWMTk2YzAsMjUsNywzOS41LDI3LjMsMzkuNWM3LjEsMC4xLDE0LjItMC43LDIxLjEtMi41IGwxLjcsMjcuN2MtMTAuMywzLjctMjEuMyw1LjQtMzIuMiw1Yy03LjMsMC40LTE0LjYtMC43LTIxLjMtMy40Yy02LjgtMi43LTEyLjktNi44LTE3LjktMTIuMWMtMTAuMy0xMC45LTE0LjEtMjktMTQuMS01Mi45IFY4Ni41aC0zMVY1OC4zaDMxVjkuNkwtNDEzLjcsMEwtNDEzLjcsMHoiLz4KICAgICAgPC9nPgogICAgPC9nPgogICAgPGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoOTc0LjQzMzI4NiwgNTMuNDc5NjM4KSI+CiAgICAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDAuOTkwMDM0LCAwLjYxMDMzOSkiPgogICAgICAgIDxwYXRoIGQ9Ik0tNDQ1LjgsMTEzYzAuOCw1MCwzMi4yLDcwLjYsNjguNiw3MC42YzE5LDAuNiwzNy45LTMsNTUuMy0xMC41bDYuMiwyNi40IGMtMjAuOSw4LjktNDMuNSwxMy4xLTY2LjIsMTIuNmMtNjEuNSwwLTk4LjMtNDEuMi05OC4zLTEwMi41Qy00ODAuMiw0OC4yLTQ0NC43LDAtMzg2LjUsMGM2NS4yLDAsODIuNyw1OC4zLDgyLjcsOTUuNyBjLTAuMSw1LjgtMC41LDExLjUtMS4yLDE3LjJoLTE0MC42SC00NDUuOEwtNDQ1LjgsMTEzeiBNLTMzOS4yLDg2LjZjMC40LTIzLjUtOS41LTYwLjEtNTAuNC02MC4xIGMtMzYuOCwwLTUyLjgsMzQuNC01NS43LDYwLjFILTMzOS4yTC0zMzkuMiw4Ni42TC0zMzkuMiw4Ni42eiIvPgogICAgICA8L2c+CiAgICA8L2c+CiAgICA8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgxMjAxLjk2MTA1OCwgNTMuNDc5NjM4KSI+CiAgICAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDEuMTc5NjQwLCAwLjcwNTA2OCkiPgogICAgICAgIDxwYXRoIGQ9Ik0tNDc4LjYsNjhjMC0yMy45LTAuNC00NC41LTEuNy02My40aDMxLjhsMS4yLDM5LjloMS43YzkuMS0yNy4zLDMxLTQ0LjUsNTUuMy00NC41IGMzLjUtMC4xLDcsMC40LDEwLjMsMS4ydjM0LjhjLTQuMS0wLjktOC4yLTEuMy0xMi40LTEuMmMtMjUuNiwwLTQzLjcsMTkuNy00OC43LDQ3LjRjLTEsNS43LTEuNiwxMS41LTEuNywxNy4ydjEwOC4zaC0zNlY2OCBMLTQ3OC42LDY4eiIvPgogICAgICA8L2c+CiAgICA8L2c+CiAgPC9nPgoKICA8ZyBjbGFzcz0ianAtaWNvbi13YXJuMCIgZmlsbD0iI0YzNzcyNiI+CiAgICA8cGF0aCBkPSJNMTM1Mi4zLDMyNi4yaDM3VjI4aC0zN1YzMjYuMnogTTE2MDQuOCwzMjYuMmMtMi41LTEzLjktMy40LTMxLjEtMy40LTQ4Ljd2LTc2IGMwLTQwLjctMTUuMS04My4xLTc3LjMtODMuMWMtMjUuNiwwLTUwLDcuMS02Ni44LDE4LjFsOC40LDI0LjRjMTQuMy05LjIsMzQtMTUuMSw1My0xNS4xYzQxLjYsMCw0Ni4yLDMwLjIsNDYuMiw0N3Y0LjIgYy03OC42LTAuNC0xMjIuMywyNi41LTEyMi4zLDc1LjZjMCwyOS40LDIxLDU4LjQsNjIuMiw1OC40YzI5LDAsNTAuOS0xNC4zLDYyLjItMzAuMmgxLjNsMi45LDI1LjZIMTYwNC44eiBNMTU2NS43LDI1Ny43IGMwLDMuOC0wLjgsOC0yLjEsMTEuOGMtNS45LDE3LjItMjIuNywzNC00OS4yLDM0Yy0xOC45LDAtMzQuOS0xMS4zLTM0LjktMzUuM2MwLTM5LjUsNDUuOC00Ni42LDg2LjItNDUuOFYyNTcuN3ogTTE2OTguNSwzMjYuMiBsMS43LTMzLjZoMS4zYzE1LjEsMjYuOSwzOC43LDM4LjIsNjguMSwzOC4yYzQ1LjQsMCw5MS4yLTM2LjEsOTEuMi0xMDguOGMwLjQtNjEuNy0zNS4zLTEwMy43LTg1LjctMTAzLjcgYy0zMi44LDAtNTYuMywxNC43LTY5LjMsMzcuNGgtMC44VjI4aC0zNi42djI0NS43YzAsMTguMS0wLjgsMzguNi0xLjcsNTIuNUgxNjk4LjV6IE0xNzA0LjgsMjA4LjJjMC01LjksMS4zLTEwLjksMi4xLTE1LjEgYzcuNi0yOC4xLDMxLjEtNDUuNCw1Ni4zLTQ1LjRjMzkuNSwwLDYwLjUsMzQuOSw2MC41LDc1LjZjMCw0Ni42LTIzLjEsNzguMS02MS44LDc4LjFjLTI2LjksMC00OC4zLTE3LjYtNTUuNS00My4zIGMtMC44LTQuMi0xLjctOC44LTEuNy0xMy40VjIwOC4yeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-kernel: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uMiIgZmlsbD0iIzYxNjE2MSIgZD0iTTE1IDlIOXY2aDZWOXptLTIgNGgtMnYtMmgydjJ6bTgtMlY5aC0yVjdjMC0xLjEtLjktMi0yLTJoLTJWM2gtMnYyaC0yVjNIOXYySDdjLTEuMSAwLTIgLjktMiAydjJIM3YyaDJ2MkgzdjJoMnYyYzAgMS4xLjkgMiAyIDJoMnYyaDJ2LTJoMnYyaDJ2LTJoMmMxLjEgMCAyLS45IDItMnYtMmgydi0yaC0ydi0yaDJ6bS00IDZIN1Y3aDEwdjEweiIvPgo8L3N2Zz4K);
	  --jp-icon-keyboard: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMjAgNUg0Yy0xLjEgMC0xLjk5LjktMS45OSAyTDIgMTdjMCAxLjEuOSAyIDIgMmgxNmMxLjEgMCAyLS45IDItMlY3YzAtMS4xLS45LTItMi0yem0tOSAzaDJ2MmgtMlY4em0wIDNoMnYyaC0ydi0yek04IDhoMnYySDhWOHptMCAzaDJ2Mkg4di0yem0tMSAySDV2LTJoMnYyem0wLTNINVY4aDJ2MnptOSA3SDh2LTJoOHYyem0wLTRoLTJ2LTJoMnYyem0wLTNoLTJWOGgydjJ6bTMgM2gtMnYtMmgydjJ6bTAtM2gtMlY4aDJ2MnoiLz4KPC9zdmc+Cg==);
	  --jp-icon-launcher: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMTkgMTlINVY1aDdWM0g1YTIgMiAwIDAwLTIgMnYxNGEyIDIgMCAwMDIgMmgxNGMxLjEgMCAyLS45IDItMnYtN2gtMnY3ek0xNCAzdjJoMy41OWwtOS44MyA5LjgzIDEuNDEgMS40MUwxOSA2LjQxVjEwaDJWM2gtN3oiLz4KPC9zdmc+Cg==);
	  --jp-icon-line-form: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICAgIDxwYXRoIGZpbGw9IndoaXRlIiBkPSJNNS44OCA0LjEyTDEzLjc2IDEybC03Ljg4IDcuODhMOCAyMmwxMC0xMEw4IDJ6Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-link: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTMuOSAxMmMwLTEuNzEgMS4zOS0zLjEgMy4xLTMuMWg0VjdIN2MtMi43NiAwLTUgMi4yNC01IDVzMi4yNCA1IDUgNWg0di0xLjlIN2MtMS43MSAwLTMuMS0xLjM5LTMuMS0zLjF6TTggMTNoOHYtMkg4djJ6bTktNmgtNHYxLjloNGMxLjcxIDAgMy4xIDEuMzkgMy4xIDMuMXMtMS4zOSAzLjEtMy4xIDMuMWgtNFYxN2g0YzIuNzYgMCA1LTIuMjQgNS01cy0yLjI0LTUtNS01eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-list: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uMiBqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiM2MTYxNjEiIGQ9Ik0xOSA1djE0SDVWNWgxNG0xLjEtMkgzLjljLS41IDAtLjkuNC0uOS45djE2LjJjMCAuNC40LjkuOS45aDE2LjJjLjQgMCAuOS0uNS45LS45VjMuOWMwLS41LS41LS45LS45LS45ek0xMSA3aDZ2MmgtNlY3em0wIDRoNnYyaC02di0yem0wIDRoNnYyaC02ek03IDdoMnYySDd6bTAgNGgydjJIN3ptMCA0aDJ2Mkg3eiIvPgo8L3N2Zz4=);
	  --jp-icon-listings-info: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iaXNvLTg4NTktMSI/Pg0KPHN2ZyB2ZXJzaW9uPSIxLjEiIGlkPSJDYXBhXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4Ig0KCSB2aWV3Qm94PSIwIDAgNTAuOTc4IDUwLjk3OCIgc3R5bGU9ImVuYWJsZS1iYWNrZ3JvdW5kOm5ldyAwIDAgNTAuOTc4IDUwLjk3ODsiIHhtbDpzcGFjZT0icHJlc2VydmUiPg0KPGc+DQoJPGc+DQoJCTxnPg0KCQkJPHBhdGggc3R5bGU9ImZpbGw6IzAxMDAwMjsiIGQ9Ik00My41Miw3LjQ1OEMzOC43MTEsMi42NDgsMzIuMzA3LDAsMjUuNDg5LDBDMTguNjcsMCwxMi4yNjYsMi42NDgsNy40NTgsNy40NTgNCgkJCQljLTkuOTQzLDkuOTQxLTkuOTQzLDI2LjExOSwwLDM2LjA2MmM0LjgwOSw0LjgwOSwxMS4yMTIsNy40NTYsMTguMDMxLDcuNDU4YzAsMCwwLjAwMSwwLDAuMDAyLDANCgkJCQljNi44MTYsMCwxMy4yMjEtMi42NDgsMTguMDI5LTcuNDU4YzQuODA5LTQuODA5LDcuNDU3LTExLjIxMiw3LjQ1Ny0xOC4wM0M1MC45NzcsMTguNjcsNDguMzI4LDEyLjI2Niw0My41Miw3LjQ1OHoNCgkJCQkgTTQyLjEwNiw0Mi4xMDVjLTQuNDMyLDQuNDMxLTEwLjMzMiw2Ljg3Mi0xNi42MTUsNi44NzJoLTAuMDAyYy02LjI4NS0wLjAwMS0xMi4xODctMi40NDEtMTYuNjE3LTYuODcyDQoJCQkJYy05LjE2Mi05LjE2My05LjE2Mi0yNC4wNzEsMC0zMy4yMzNDMTMuMzAzLDQuNDQsMTkuMjA0LDIsMjUuNDg5LDJjNi4yODQsMCwxMi4xODYsMi40NCwxNi42MTcsNi44NzINCgkJCQljNC40MzEsNC40MzEsNi44NzEsMTAuMzMyLDYuODcxLDE2LjYxN0M0OC45NzcsMzEuNzcyLDQ2LjUzNiwzNy42NzUsNDIuMTA2LDQyLjEwNXoiLz4NCgkJPC9nPg0KCQk8Zz4NCgkJCTxwYXRoIHN0eWxlPSJmaWxsOiMwMTAwMDI7IiBkPSJNMjMuNTc4LDMyLjIxOGMtMC4wMjMtMS43MzQsMC4xNDMtMy4wNTksMC40OTYtMy45NzJjMC4zNTMtMC45MTMsMS4xMS0xLjk5NywyLjI3Mi0zLjI1Mw0KCQkJCWMwLjQ2OC0wLjUzNiwwLjkyMy0xLjA2MiwxLjM2Ny0xLjU3NWMwLjYyNi0wLjc1MywxLjEwNC0xLjQ3OCwxLjQzNi0yLjE3NWMwLjMzMS0wLjcwNywwLjQ5NS0xLjU0MSwwLjQ5NS0yLjUNCgkJCQljMC0xLjA5Ni0wLjI2LTIuMDg4LTAuNzc5LTIuOTc5Yy0wLjU2NS0wLjg3OS0xLjUwMS0xLjMzNi0yLjgwNi0xLjM2OWMtMS44MDIsMC4wNTctMi45ODUsMC42NjctMy41NSwxLjgzMg0KCQkJCWMtMC4zMDEsMC41MzUtMC41MDMsMS4xNDEtMC42MDcsMS44MTRjLTAuMTM5LDAuNzA3LTAuMjA3LDEuNDMyLTAuMjA3LDIuMTc0aC0yLjkzN2MtMC4wOTEtMi4yMDgsMC40MDctNC4xMTQsMS40OTMtNS43MTkNCgkJCQljMS4wNjItMS42NCwyLjg1NS0yLjQ4MSw1LjM3OC0yLjUyN2MyLjE2LDAuMDIzLDMuODc0LDAuNjA4LDUuMTQxLDEuNzU4YzEuMjc4LDEuMTYsMS45MjksMi43NjQsMS45NSw0LjgxMQ0KCQkJCWMwLDEuMTQyLTAuMTM3LDIuMTExLTAuNDEsMi45MTFjLTAuMzA5LDAuODQ1LTAuNzMxLDEuNTkzLTEuMjY4LDIuMjQzYy0wLjQ5MiwwLjY1LTEuMDY4LDEuMzE4LTEuNzMsMi4wMDINCgkJCQljLTAuNjUsMC42OTctMS4zMTMsMS40NzktMS45ODcsMi4zNDZjLTAuMjM5LDAuMzc3LTAuNDI5LDAuNzc3LTAuNTY1LDEuMTk5Yy0wLjE2LDAuOTU5LTAuMjE3LDEuOTUxLTAuMTcxLDIuOTc5DQoJCQkJQzI2LjU4OSwzMi4yMTgsMjMuNTc4LDMyLjIxOCwyMy41NzgsMzIuMjE4eiBNMjMuNTc4LDM4LjIydi0zLjQ4NGgzLjA3NnYzLjQ4NEgyMy41Nzh6Ii8+DQoJCTwvZz4NCgk8L2c+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8Zz4NCjwvZz4NCjxnPg0KPC9nPg0KPGc+DQo8L2c+DQo8L3N2Zz4NCg==);
	  --jp-icon-markdown: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8cGF0aCBjbGFzcz0ianAtaWNvbi1jb250cmFzdDAganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjN0IxRkEyIiBkPSJNNSAxNC45aDEybC02LjEgNnptOS40LTYuOGMwLTEuMy0uMS0yLjktLjEtNC41LS40IDEuNC0uOSAyLjktMS4zIDQuM2wtMS4zIDQuM2gtMkw4LjUgNy45Yy0uNC0xLjMtLjctMi45LTEtNC4zLS4xIDEuNi0uMSAzLjItLjIgNC42TDcgMTIuNEg0LjhsLjctMTFoMy4zTDEwIDVjLjQgMS4yLjcgMi43IDEgMy45LjMtMS4yLjctMi42IDEtMy45bDEuMi0zLjdoMy4zbC42IDExaC0yLjRsLS4zLTQuMnoiLz4KPC9zdmc+Cg==);
	  --jp-icon-new-folder: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTIwIDZoLThsLTItMkg0Yy0xLjExIDAtMS45OS44OS0xLjk5IDJMMiAxOGMwIDEuMTEuODkgMiAyIDJoMTZjMS4xMSAwIDItLjg5IDItMlY4YzAtMS4xMS0uODktMi0yLTJ6bS0xIDhoLTN2M2gtMnYtM2gtM3YtMmgzVjloMnYzaDN2MnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-not-trusted: url(data:image/svg+xml;base64,PHN2ZyBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI1IDI1Ij4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uMiIgc3Ryb2tlPSIjMzMzMzMzIiBzdHJva2Utd2lkdGg9IjIiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDMgMykiIGQ9Ik0xLjg2MDk0IDExLjQ0MDlDMC44MjY0NDggOC43NzAyNyAwLjg2Mzc3OSA2LjA1NzY0IDEuMjQ5MDcgNC4xOTkzMkMyLjQ4MjA2IDMuOTMzNDcgNC4wODA2OCAzLjQwMzQ3IDUuNjAxMDIgMi44NDQ5QzcuMjM1NDkgMi4yNDQ0IDguODU2NjYgMS41ODE1IDkuOTg3NiAxLjA5NTM5QzExLjA1OTcgMS41ODM0MSAxMi42MDk0IDIuMjQ0NCAxNC4yMTggMi44NDMzOUMxNS43NTAzIDMuNDEzOTQgMTcuMzk5NSAzLjk1MjU4IDE4Ljc1MzkgNC4yMTM4NUMxOS4xMzY0IDYuMDcxNzcgMTkuMTcwOSA4Ljc3NzIyIDE4LjEzOSAxMS40NDA5QzE3LjAzMDMgMTQuMzAzMiAxNC42NjY4IDE3LjE4NDQgOS45OTk5OSAxOC45MzU0QzUuMzMzMTkgMTcuMTg0NCAyLjk2OTY4IDE0LjMwMzIgMS44NjA5NCAxMS40NDA5WiIvPgogICAgPHBhdGggY2xhc3M9ImpwLWljb24yIiBzdHJva2U9IiMzMzMzMzMiIHN0cm9rZS13aWR0aD0iMiIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoOS4zMTU5MiA5LjMyMDMxKSIgZD0iTTcuMzY4NDIgMEwwIDcuMzY0NzkiLz4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uMiIgc3Ryb2tlPSIjMzMzMzMzIiBzdHJva2Utd2lkdGg9IjIiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDkuMzE1OTIgMTYuNjgzNikgc2NhbGUoMSAtMSkiIGQ9Ik03LjM2ODQyIDBMMCA3LjM2NDc5Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-notebook: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8ZyBjbGFzcz0ianAtaWNvbi13YXJuMCBqcC1pY29uLXNlbGVjdGFibGUiIGZpbGw9IiNFRjZDMDAiPgogICAgPHBhdGggZD0iTTE4LjcgMy4zdjE1LjRIMy4zVjMuM2gxNS40bTEuNS0xLjVIMS44djE4LjNoMTguM2wuMS0xOC4zeiIvPgogICAgPHBhdGggZD0iTTE2LjUgMTYuNWwtNS40LTQuMy01LjYgNC4zdi0xMWgxMXoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-numbering: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjIiIGhlaWdodD0iMjIiIHZpZXdCb3g9IjAgMCAyOCAyOCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CgkJPHBhdGggZD0iTTQgMTlINlYxOS41SDVWMjAuNUg2VjIxSDRWMjJIN1YxOEg0VjE5Wk01IDEwSDZWNkg0VjdINVYxMFpNNCAxM0g1LjhMNCAxNS4xVjE2SDdWMTVINS4yTDcgMTIuOVYxMkg0VjEzWk05IDdWOUgyM1Y3SDlaTTkgMjFIMjNWMTlIOVYyMVpNOSAxNUgyM1YxM0g5VjE1WiIvPgoJPC9nPgo8L3N2Zz4K);
	  --jp-icon-offline-bolt: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgd2lkdGg9IjE2Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEyIDIuMDJjLTUuNTEgMC05Ljk4IDQuNDctOS45OCA5Ljk4czQuNDcgOS45OCA5Ljk4IDkuOTggOS45OC00LjQ3IDkuOTgtOS45OFMxNy41MSAyLjAyIDEyIDIuMDJ6TTExLjQ4IDIwdi02LjI2SDhMMTMgNHY2LjI2aDMuMzVMMTEuNDggMjB6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-palette: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTE4IDEzVjIwSDRWNkg5LjAyQzkuMDcgNS4yOSA5LjI0IDQuNjIgOS41IDRINEMyLjkgNCAyIDQuOSAyIDZWMjBDMiAyMS4xIDIuOSAyMiA0IDIySDE4QzE5LjEgMjIgMjAgMjEuMSAyMCAyMFYxNUwxOCAxM1pNMTkuMyA4Ljg5QzE5Ljc0IDguMTkgMjAgNy4zOCAyMCA2LjVDMjAgNC4wMSAxNy45OSAyIDE1LjUgMkMxMy4wMSAyIDExIDQuMDEgMTEgNi41QzExIDguOTkgMTMuMDEgMTEgMTUuNDkgMTFDMTYuMzcgMTEgMTcuMTkgMTAuNzQgMTcuODggMTAuM0wyMSAxMy40MkwyMi40MiAxMkwxOS4zIDguODlaTTE1LjUgOUMxNC4xMiA5IDEzIDcuODggMTMgNi41QzEzIDUuMTIgMTQuMTIgNCAxNS41IDRDMTYuODggNCAxOCA1LjEyIDE4IDYuNUMxOCA3Ljg4IDE2Ljg4IDkgMTUuNSA5WiIvPgogICAgPHBhdGggZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiIGQ9Ik00IDZIOS4wMTg5NEM5LjAwNjM5IDYuMTY1MDIgOSA2LjMzMTc2IDkgNi41QzkgOC44MTU3NyAxMC4yMTEgMTAuODQ4NyAxMi4wMzQzIDEySDlWMTRIMTZWMTIuOTgxMUMxNi41NzAzIDEyLjkzNzcgMTcuMTIgMTIuODIwNyAxNy42Mzk2IDEyLjYzOTZMMTggMTNWMjBINFY2Wk04IDhINlYxMEg4VjhaTTYgMTJIOFYxNEg2VjEyWk04IDE2SDZWMThIOFYxNlpNOSAxNkgxNlYxOEg5VjE2WiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-paste: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTE5IDJoLTQuMThDMTQuNC44NCAxMy4zIDAgMTIgMGMtMS4zIDAtMi40Ljg0LTIuODIgMkg1Yy0xLjEgMC0yIC45LTIgMnYxNmMwIDEuMS45IDIgMiAyaDE0YzEuMSAwIDItLjkgMi0yVjRjMC0xLjEtLjktMi0yLTJ6bS03IDBjLjU1IDAgMSAuNDUgMSAxcy0uNDUgMS0xIDEtMS0uNDUtMS0xIC40NS0xIDEtMXptNyAxOEg1VjRoMnYzaDEwVjRoMnYxNnoiLz4KICAgIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-pdf: url(data:image/svg+xml;base64,PHN2ZwogICB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyMiAyMiIgd2lkdGg9IjE2Ij4KICAgIDxwYXRoIHRyYW5zZm9ybT0icm90YXRlKDQ1KSIgY2xhc3M9ImpwLWljb24tc2VsZWN0YWJsZSIgZmlsbD0iI0ZGMkEyQSIKICAgICAgIGQ9Im0gMjIuMzQ0MzY5LC0zLjAxNjM2NDIgaCA1LjYzODYwNCB2IDEuNTc5MjQzMyBoIC0zLjU0OTIyNyB2IDEuNTA4NjkyOTkgaCAzLjMzNzU3NiBWIDEuNjUwODE1NCBoIC0zLjMzNzU3NiB2IDMuNDM1MjYxMyBoIC0yLjA4OTM3NyB6IG0gLTcuMTM2NDQ0LDEuNTc5MjQzMyB2IDQuOTQzOTU0MyBoIDAuNzQ4OTIgcSAxLjI4MDc2MSwwIDEuOTUzNzAzLC0wLjYzNDk1MzUgMC42NzgzNjksLTAuNjM0OTUzNSAwLjY3ODM2OSwtMS44NDUxNjQxIDAsLTEuMjA0NzgzNTUgLTAuNjcyOTQyLC0xLjgzNDMxMDExIC0wLjY3Mjk0MiwtMC42Mjk1MjY1OSAtMS45NTkxMywtMC42Mjk1MjY1OSB6IG0gLTIuMDg5Mzc3LC0xLjU3OTI0MzMgaCAyLjIwMzM0MyBxIDEuODQ1MTY0LDAgMi43NDYwMzksMC4yNjU5MjA3IDAuOTA2MzAxLDAuMjYwNDkzNyAxLjU1MjEwOCwwLjg5MDAyMDMgMC41Njk4MywwLjU0ODEyMjMgMC44NDY2MDUsMS4yNjQ0ODAwNiAwLjI3Njc3NCwwLjcxNjM1NzgxIDAuMjc2Nzc0LDEuNjIyNjU4OTQgMCwwLjkxNzE1NTEgLTAuMjc2Nzc0LDEuNjM4OTM5OSAtMC4yNzY3NzUsMC43MTYzNTc4IC0wLjg0NjYwNSwxLjI2NDQ4IC0wLjY1MTIzNCwwLjYyOTUyNjYgLTEuNTYyOTYyLDAuODk1NDQ3MyAtMC45MTE3MjgsMC4yNjA0OTM3IC0yLjczNTE4NSwwLjI2MDQ5MzcgaCAtMi4yMDMzNDMgeiBtIC04LjE0NTg1NjUsMCBoIDMuNDY3ODIzIHEgMS41NDY2ODE2LDAgMi4zNzE1Nzg1LDAuNjg5MjIzIDAuODMwMzI0LDAuNjgzNzk2MSAwLjgzMDMyNCwxLjk1MzcwMzE0IDAsMS4yNzUzMzM5NyAtMC44MzAzMjQsMS45NjQ1NTcwNiBRIDkuOTg3MTk2MSwyLjI3NDkxNSA4LjQ0MDUxNDUsMi4yNzQ5MTUgSCA3LjA2MjA2ODQgViA1LjA4NjA3NjcgSCA0Ljk3MjY5MTUgWiBtIDIuMDg5Mzc2OSwxLjUxNDExOTkgdiAyLjI2MzAzOTQzIGggMS4xNTU5NDEgcSAwLjYwNzgxODgsMCAwLjkzODg2MjksLTAuMjkzMDU1NDcgMC4zMzEwNDQxLC0wLjI5ODQ4MjQxIDAuMzMxMDQ0MSwtMC44NDExNzc3MiAwLC0wLjU0MjY5NTMxIC0wLjMzMTA0NDEsLTAuODM1NzUwNzQgLTAuMzMxMDQ0MSwtMC4yOTMwNTU1IC0wLjkzODg2MjksLTAuMjkzMDU1NSB6IgovPgo8L3N2Zz4K);
	  --jp-icon-python: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8ZyBjbGFzcz0ianAtaWNvbi1icmFuZDAganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjMEQ0N0ExIj4KICAgIDxwYXRoIGQ9Ik0xMS4xIDYuOVY1LjhINi45YzAtLjUgMC0xLjMuMi0xLjYuNC0uNy44LTEuMSAxLjctMS40IDEuNy0uMyAyLjUtLjMgMy45LS4xIDEgLjEgMS45LjkgMS45IDEuOXY0LjJjMCAuNS0uOSAxLjYtMiAxLjZIOC44Yy0xLjUgMC0yLjQgMS40LTIuNCAyLjh2Mi4ySDQuN0MzLjUgMTUuMSAzIDE0IDMgMTMuMVY5Yy0uMS0xIC42LTIgMS44LTIgMS41LS4xIDYuMy0uMSA2LjMtLjF6Ii8+CiAgICA8cGF0aCBkPSJNMTAuOSAxNS4xdjEuMWg0LjJjMCAuNSAwIDEuMy0uMiAxLjYtLjQuNy0uOCAxLjEtMS43IDEuNC0xLjcuMy0yLjUuMy0zLjkuMS0xLS4xLTEuOS0uOS0xLjktMS45di00LjJjMC0uNS45LTEuNiAyLTEuNmgzLjhjMS41IDAgMi40LTEuNCAyLjQtMi44VjYuNmgxLjdDMTguNSA2LjkgMTkgOCAxOSA4LjlWMTNjMCAxLS43IDIuMS0xLjkgMi4xaC02LjJ6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-r-kernel: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8cGF0aCBjbGFzcz0ianAtaWNvbi1jb250cmFzdDMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjMjE5NkYzIiBkPSJNNC40IDIuNWMxLjItLjEgMi45LS4zIDQuOS0uMyAyLjUgMCA0LjEuNCA1LjIgMS4zIDEgLjcgMS41IDEuOSAxLjUgMy41IDAgMi0xLjQgMy41LTIuOSA0LjEgMS4yLjQgMS43IDEuNiAyLjIgMyAuNiAxLjkgMSAzLjkgMS4zIDQuNmgtMy44Yy0uMy0uNC0uOC0xLjctMS4yLTMuN3MtMS4yLTIuNi0yLjYtMi42aC0uOXY2LjRINC40VjIuNXptMy43IDYuOWgxLjRjMS45IDAgMi45LS45IDIuOS0yLjNzLTEtMi4zLTIuOC0yLjNjLS43IDAtMS4zIDAtMS42LjJ2NC41aC4xdi0uMXoiLz4KPC9zdmc+Cg==);
	  --jp-icon-react: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMTUwIDE1MCA1NDEuOSAyOTUuMyI+CiAgPGcgY2xhc3M9ImpwLWljb24tYnJhbmQyIGpwLWljb24tc2VsZWN0YWJsZSIgZmlsbD0iIzYxREFGQiI+CiAgICA8cGF0aCBkPSJNNjY2LjMgMjk2LjVjMC0zMi41LTQwLjctNjMuMy0xMDMuMS04Mi40IDE0LjQtNjMuNiA4LTExNC4yLTIwLjItMTMwLjQtNi41LTMuOC0xNC4xLTUuNi0yMi40LTUuNnYyMi4zYzQuNiAwIDguMy45IDExLjQgMi42IDEzLjYgNy44IDE5LjUgMzcuNSAxNC45IDc1LjctMS4xIDkuNC0yLjkgMTkuMy01LjEgMjkuNC0xOS42LTQuOC00MS04LjUtNjMuNS0xMC45LTEzLjUtMTguNS0yNy41LTM1LjMtNDEuNi01MCAzMi42LTMwLjMgNjMuMi00Ni45IDg0LTQ2LjlWNzhjLTI3LjUgMC02My41IDE5LjYtOTkuOSA1My42LTM2LjQtMzMuOC03Mi40LTUzLjItOTkuOS01My4ydjIyLjNjMjAuNyAwIDUxLjQgMTYuNSA4NCA0Ni42LTE0IDE0LjctMjggMzEuNC00MS4zIDQ5LjktMjIuNiAyLjQtNDQgNi4xLTYzLjYgMTEtMi4zLTEwLTQtMTkuNy01LjItMjktNC43LTM4LjIgMS4xLTY3LjkgMTQuNi03NS44IDMtMS44IDYuOS0yLjYgMTEuNS0yLjZWNzguNWMtOC40IDAtMTYgMS44LTIyLjYgNS42LTI4LjEgMTYuMi0zNC40IDY2LjctMTkuOSAxMzAuMS02Mi4yIDE5LjItMTAyLjcgNDkuOS0xMDIuNyA4Mi4zIDAgMzIuNSA0MC43IDYzLjMgMTAzLjEgODIuNC0xNC40IDYzLjYtOCAxMTQuMiAyMC4yIDEzMC40IDYuNSAzLjggMTQuMSA1LjYgMjIuNSA1LjYgMjcuNSAwIDYzLjUtMTkuNiA5OS45LTUzLjYgMzYuNCAzMy44IDcyLjQgNTMuMiA5OS45IDUzLjIgOC40IDAgMTYtMS44IDIyLjYtNS42IDI4LjEtMTYuMiAzNC40LTY2LjcgMTkuOS0xMzAuMSA2Mi0xOS4xIDEwMi41LTQ5LjkgMTAyLjUtODIuM3ptLTEzMC4yLTY2LjdjLTMuNyAxMi45LTguMyAyNi4yLTEzLjUgMzkuNS00LjEtOC04LjQtMTYtMTMuMS0yNC00LjYtOC05LjUtMTUuOC0xNC40LTIzLjQgMTQuMiAyLjEgMjcuOSA0LjcgNDEgNy45em0tNDUuOCAxMDYuNWMtNy44IDEzLjUtMTUuOCAyNi4zLTI0LjEgMzguMi0xNC45IDEuMy0zMCAyLTQ1LjIgMi0xNS4xIDAtMzAuMi0uNy00NS0xLjktOC4zLTExLjktMTYuNC0yNC42LTI0LjItMzgtNy42LTEzLjEtMTQuNS0yNi40LTIwLjgtMzkuOCA2LjItMTMuNCAxMy4yLTI2LjggMjAuNy0zOS45IDcuOC0xMy41IDE1LjgtMjYuMyAyNC4xLTM4LjIgMTQuOS0xLjMgMzAtMiA0NS4yLTIgMTUuMSAwIDMwLjIuNyA0NSAxLjkgOC4zIDExLjkgMTYuNCAyNC42IDI0LjIgMzggNy42IDEzLjEgMTQuNSAyNi40IDIwLjggMzkuOC02LjMgMTMuNC0xMy4yIDI2LjgtMjAuNyAzOS45em0zMi4zLTEzYzUuNCAxMy40IDEwIDI2LjggMTMuOCAzOS44LTEzLjEgMy4yLTI2LjkgNS45LTQxLjIgOCA0LjktNy43IDkuOC0xNS42IDE0LjQtMjMuNyA0LjYtOCA4LjktMTYuMSAxMy0yNC4xek00MjEuMiA0MzBjLTkuMy05LjYtMTguNi0yMC4zLTI3LjgtMzIgOSAuNCAxOC4yLjcgMjcuNS43IDkuNCAwIDE4LjctLjIgMjcuOC0uNy05IDExLjctMTguMyAyMi40LTI3LjUgMzJ6bS03NC40LTU4LjljLTE0LjItMi4xLTI3LjktNC43LTQxLTcuOSAzLjctMTIuOSA4LjMtMjYuMiAxMy41LTM5LjUgNC4xIDggOC40IDE2IDEzLjEgMjQgNC43IDggOS41IDE1LjggMTQuNCAyMy40ek00MjAuNyAxNjNjOS4zIDkuNiAxOC42IDIwLjMgMjcuOCAzMi05LS40LTE4LjItLjctMjcuNS0uNy05LjQgMC0xOC43LjItMjcuOC43IDktMTEuNyAxOC4zLTIyLjQgMjcuNS0zMnptLTc0IDU4LjljLTQuOSA3LjctOS44IDE1LjYtMTQuNCAyMy43LTQuNiA4LTguOSAxNi0xMyAyNC01LjQtMTMuNC0xMC0yNi44LTEzLjgtMzkuOCAxMy4xLTMuMSAyNi45LTUuOCA0MS4yLTcuOXptLTkwLjUgMTI1LjJjLTM1LjQtMTUuMS01OC4zLTM0LjktNTguMy01MC42IDAtMTUuNyAyMi45LTM1LjYgNTguMy01MC42IDguNi0zLjcgMTgtNyAyNy43LTEwLjEgNS43IDE5LjYgMTMuMiA0MCAyMi41IDYwLjktOS4yIDIwLjgtMTYuNiA0MS4xLTIyLjIgNjAuNi05LjktMy4xLTE5LjMtNi41LTI4LTEwLjJ6TTMxMCA0OTBjLTEzLjYtNy44LTE5LjUtMzcuNS0xNC45LTc1LjcgMS4xLTkuNCAyLjktMTkuMyA1LjEtMjkuNCAxOS42IDQuOCA0MSA4LjUgNjMuNSAxMC45IDEzLjUgMTguNSAyNy41IDM1LjMgNDEuNiA1MC0zMi42IDMwLjMtNjMuMiA0Ni45LTg0IDQ2LjktNC41LS4xLTguMy0xLTExLjMtMi43em0yMzcuMi03Ni4yYzQuNyAzOC4yLTEuMSA2Ny45LTE0LjYgNzUuOC0zIDEuOC02LjkgMi42LTExLjUgMi42LTIwLjcgMC01MS40LTE2LjUtODQtNDYuNiAxNC0xNC43IDI4LTMxLjQgNDEuMy00OS45IDIyLjYtMi40IDQ0LTYuMSA2My42LTExIDIuMyAxMC4xIDQuMSAxOS44IDUuMiAyOS4xem0zOC41LTY2LjdjLTguNiAzLjctMTggNy0yNy43IDEwLjEtNS43LTE5LjYtMTMuMi00MC0yMi41LTYwLjkgOS4yLTIwLjggMTYuNi00MS4xIDIyLjItNjAuNiA5LjkgMy4xIDE5LjMgNi41IDI4LjEgMTAuMiAzNS40IDE1LjEgNTguMyAzNC45IDU4LjMgNTAuNi0uMSAxNS43LTIzIDM1LjYtNTguNCA1MC42ek0zMjAuOCA3OC40eiIvPgogICAgPGNpcmNsZSBjeD0iNDIwLjkiIGN5PSIyOTYuNSIgcj0iNDUuNyIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-redo: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgd2lkdGg9IjE2Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgICA8cGF0aCBkPSJNMCAwaDI0djI0SDB6IiBmaWxsPSJub25lIi8+PHBhdGggZD0iTTE4LjQgMTAuNkMxNi41NSA4Ljk5IDE0LjE1IDggMTEuNSA4Yy00LjY1IDAtOC41OCAzLjAzLTkuOTYgNy4yMkwzLjkgMTZjMS4wNS0zLjE5IDQuMDUtNS41IDcuNi01LjUgMS45NSAwIDMuNzMuNzIgNS4xMiAxLjg4TDEzIDE2aDlWN2wtMy42IDMuNnoiLz4KICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-refresh: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDE4IDE4Ij4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTkgMTMuNWMtMi40OSAwLTQuNS0yLjAxLTQuNS00LjVTNi41MSA0LjUgOSA0LjVjMS4yNCAwIDIuMzYuNTIgMy4xNyAxLjMzTDEwIDhoNVYzbC0xLjc2IDEuNzZDMTIuMTUgMy42OCAxMC42NiAzIDkgMyA1LjY5IDMgMy4wMSA1LjY5IDMuMDEgOVM1LjY5IDE1IDkgMTVjMi45NyAwIDUuNDMtMi4xNiA1LjktNWgtMS41MmMtLjQ2IDItMi4yNCAzLjUtNC4zOCAzLjV6Ii8+CiAgICA8L2c+Cjwvc3ZnPgo=);
	  --jp-icon-regex: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIwIDIwIj4KICA8ZyBjbGFzcz0ianAtaWNvbjIiIGZpbGw9IiM0MTQxNDEiPgogICAgPHJlY3QgeD0iMiIgeT0iMiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ii8+CiAgPC9nPgoKICA8ZyBjbGFzcz0ianAtaWNvbi1hY2NlbnQyIiBmaWxsPSIjRkZGIj4KICAgIDxjaXJjbGUgY2xhc3M9InN0MiIgY3g9IjUuNSIgY3k9IjE0LjUiIHI9IjEuNSIvPgogICAgPHJlY3QgeD0iMTIiIHk9IjQiIGNsYXNzPSJzdDIiIHdpZHRoPSIxIiBoZWlnaHQ9IjgiLz4KICAgIDxyZWN0IHg9IjguNSIgeT0iNy41IiB0cmFuc2Zvcm09Im1hdHJpeCgwLjg2NiAtMC41IDAuNSAwLjg2NiAtMi4zMjU1IDcuMzIxOSkiIGNsYXNzPSJzdDIiIHdpZHRoPSI4IiBoZWlnaHQ9IjEiLz4KICAgIDxyZWN0IHg9IjEyIiB5PSI0IiB0cmFuc2Zvcm09Im1hdHJpeCgwLjUgLTAuODY2IDAuODY2IDAuNSAtMC42Nzc5IDE0LjgyNTIpIiBjbGFzcz0ic3QyIiB3aWR0aD0iMSIgaGVpZ2h0PSI4Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-run: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTggNXYxNGwxMS03eiIvPgogICAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-running: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDUxMiA1MTIiPgogIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICA8cGF0aCBkPSJNMjU2IDhDMTE5IDggOCAxMTkgOCAyNTZzMTExIDI0OCAyNDggMjQ4IDI0OC0xMTEgMjQ4LTI0OFMzOTMgOCAyNTYgOHptOTYgMzI4YzAgOC44LTcuMiAxNi0xNiAxNkgxNzZjLTguOCAwLTE2LTcuMi0xNi0xNlYxNzZjMC04LjggNy4yLTE2IDE2LTE2aDE2MGM4LjggMCAxNiA3LjIgMTYgMTZ2MTYweiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-save: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTE3IDNINWMtMS4xMSAwLTIgLjktMiAydjE0YzAgMS4xLjg5IDIgMiAyaDE0YzEuMSAwIDItLjkgMi0yVjdsLTQtNHptLTUgMTZjLTEuNjYgMC0zLTEuMzQtMy0zczEuMzQtMyAzLTMgMyAxLjM0IDMgMy0xLjM0IDMtMyAzem0zLTEwSDVWNWgxMHY0eiIvPgogICAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-search: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMTggMTgiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEyLjEsMTAuOWgtMC43bC0wLjItMC4yYzAuOC0wLjksMS4zLTIuMiwxLjMtMy41YzAtMy0yLjQtNS40LTUuNC01LjRTMS44LDQuMiwxLjgsNy4xczIuNCw1LjQsNS40LDUuNCBjMS4zLDAsMi41LTAuNSwzLjUtMS4zbDAuMiwwLjJ2MC43bDQuMSw0LjFsMS4yLTEuMkwxMi4xLDEwLjl6IE03LjEsMTAuOWMtMi4xLDAtMy43LTEuNy0zLjctMy43czEuNy0zLjcsMy43LTMuN3MzLjcsMS43LDMuNywzLjcgUzkuMiwxMC45LDcuMSwxMC45eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-settings: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMTkuNDMgMTIuOThjLjA0LS4zMi4wNy0uNjQuMDctLjk4cy0uMDMtLjY2LS4wNy0uOThsMi4xMS0xLjY1Yy4xOS0uMTUuMjQtLjQyLjEyLS42NGwtMi0zLjQ2Yy0uMTItLjIyLS4zOS0uMy0uNjEtLjIybC0yLjQ5IDFjLS41Mi0uNC0xLjA4LS43My0xLjY5LS45OGwtLjM4LTIuNjVBLjQ4OC40ODggMCAwMDE0IDJoLTRjLS4yNSAwLS40Ni4xOC0uNDkuNDJsLS4zOCAyLjY1Yy0uNjEuMjUtMS4xNy41OS0xLjY5Ljk4bC0yLjQ5LTFjLS4yMy0uMDktLjQ5IDAtLjYxLjIybC0yIDMuNDZjLS4xMy4yMi0uMDcuNDkuMTIuNjRsMi4xMSAxLjY1Yy0uMDQuMzItLjA3LjY1LS4wNy45OHMuMDMuNjYuMDcuOThsLTIuMTEgMS42NWMtLjE5LjE1LS4yNC40Mi0uMTIuNjRsMiAzLjQ2Yy4xMi4yMi4zOS4zLjYxLjIybDIuNDktMWMuNTIuNCAxLjA4LjczIDEuNjkuOThsLjM4IDIuNjVjLjAzLjI0LjI0LjQyLjQ5LjQyaDRjLjI1IDAgLjQ2LS4xOC40OS0uNDJsLjM4LTIuNjVjLjYxLS4yNSAxLjE3LS41OSAxLjY5LS45OGwyLjQ5IDFjLjIzLjA5LjQ5IDAgLjYxLS4yMmwyLTMuNDZjLjEyLS4yMi4wNy0uNDktLjEyLS42NGwtMi4xMS0xLjY1ek0xMiAxNS41Yy0xLjkzIDAtMy41LTEuNTctMy41LTMuNXMxLjU3LTMuNSAzLjUtMy41IDMuNSAxLjU3IDMuNSAzLjUtMS41NyAzLjUtMy41IDMuNXoiLz4KPC9zdmc+Cg==);
	  --jp-icon-spreadsheet: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8cGF0aCBjbGFzcz0ianAtaWNvbi1jb250cmFzdDEganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNENBRjUwIiBkPSJNMi4yIDIuMnYxNy42aDE3LjZWMi4ySDIuMnptMTUuNCA3LjdoLTUuNVY0LjRoNS41djUuNXpNOS45IDQuNHY1LjVINC40VjQuNGg1LjV6bS01LjUgNy43aDUuNXY1LjVINC40di01LjV6bTcuNyA1LjV2LTUuNWg1LjV2NS41aC01LjV6Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-stop: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTAgMGgyNHYyNEgweiIgZmlsbD0ibm9uZSIvPgogICAgICAgIDxwYXRoIGQ9Ik02IDZoMTJ2MTJINnoiLz4KICAgIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-tab: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTIxIDNIM2MtMS4xIDAtMiAuOS0yIDJ2MTRjMCAxLjEuOSAyIDIgMmgxOGMxLjEgMCAyLS45IDItMlY1YzAtMS4xLS45LTItMi0yem0wIDE2SDNWNWgxMHY0aDh2MTB6Ii8+CiAgPC9nPgo8L3N2Zz4K);
	  --jp-icon-table-rows: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTAgMGgyNHYyNEgweiIgZmlsbD0ibm9uZSIvPgogICAgICAgIDxwYXRoIGQ9Ik0yMSw4SDNWNGgxOFY4eiBNMjEsMTBIM3Y0aDE4VjEweiBNMjEsMTZIM3Y0aDE4VjE2eiIvPgogICAgPC9nPgo8L3N2Zz4=);
	  --jp-icon-tag: url(data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjgiIGhlaWdodD0iMjgiIHZpZXdCb3g9IjAgMCA0MyAyOCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KCTxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CgkJPHBhdGggZD0iTTI4LjgzMzIgMTIuMzM0TDMyLjk5OTggMTYuNTAwN0wzNy4xNjY1IDEyLjMzNEgyOC44MzMyWiIvPgoJCTxwYXRoIGQ9Ik0xNi4yMDk1IDIxLjYxMDRDMTUuNjg3MyAyMi4xMjk5IDE0Ljg0NDMgMjIuMTI5OSAxNC4zMjQ4IDIxLjYxMDRMNi45ODI5IDE0LjcyNDVDNi41NzI0IDE0LjMzOTQgNi4wODMxMyAxMy42MDk4IDYuMDQ3ODYgMTMuMDQ4MkM1Ljk1MzQ3IDExLjUyODggNi4wMjAwMiA4LjYxOTQ0IDYuMDY2MjEgNy4wNzY5NUM2LjA4MjgxIDYuNTE0NzcgNi41NTU0OCA2LjA0MzQ3IDcuMTE4MDQgNi4wMzA1NUM5LjA4ODYzIDUuOTg0NzMgMTMuMjYzOCA1LjkzNTc5IDEzLjY1MTggNi4zMjQyNUwyMS43MzY5IDEzLjYzOUMyMi4yNTYgMTQuMTU4NSAyMS43ODUxIDE1LjQ3MjQgMjEuMjYyIDE1Ljk5NDZMMTYuMjA5NSAyMS42MTA0Wk05Ljc3NTg1IDguMjY1QzkuMzM1NTEgNy44MjU2NiA4LjYyMzUxIDcuODI1NjYgOC4xODI4IDguMjY1QzcuNzQzNDYgOC43MDU3MSA3Ljc0MzQ2IDkuNDE3MzMgOC4xODI4IDkuODU2NjdDOC42MjM4MiAxMC4yOTY0IDkuMzM1ODIgMTAuMjk2NCA5Ljc3NTg1IDkuODU2NjdDMTAuMjE1NiA5LjQxNzMzIDEwLjIxNTYgOC43MDUzMyA5Ljc3NTg1IDguMjY1WiIvPgoJPC9nPgo8L3N2Zz4K);
	  --jp-icon-terminal: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0IiA+CiAgICA8cmVjdCBjbGFzcz0ianAtaWNvbjIganAtaWNvbi1zZWxlY3RhYmxlIiB3aWR0aD0iMjAiIGhlaWdodD0iMjAiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDIgMikiIGZpbGw9IiMzMzMzMzMiLz4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uLWFjY2VudDIganAtaWNvbi1zZWxlY3RhYmxlLWludmVyc2UiIGQ9Ik01LjA1NjY0IDguNzYxNzJDNS4wNTY2NCA4LjU5NzY2IDUuMDMxMjUgOC40NTMxMiA0Ljk4MDQ3IDguMzI4MTJDNC45MzM1OSA4LjE5OTIyIDQuODU1NDcgOC4wODIwMyA0Ljc0NjA5IDcuOTc2NTZDNC42NDA2MiA3Ljg3MTA5IDQuNSA3Ljc3NTM5IDQuMzI0MjIgNy42ODk0NUM0LjE1MjM0IDcuNTk5NjEgMy45NDMzNiA3LjUxMTcyIDMuNjk3MjcgNy40MjU3OEMzLjMwMjczIDcuMjg1MTYgMi45NDMzNiA3LjEzNjcyIDIuNjE5MTQgNi45ODA0N0MyLjI5NDkyIDYuODI0MjIgMi4wMTc1OCA2LjY0MjU4IDEuNzg3MTEgNi40MzU1NUMxLjU2MDU1IDYuMjI4NTIgMS4zODQ3NyA1Ljk4ODI4IDEuMjU5NzcgNS43MTQ4NEMxLjEzNDc3IDUuNDM3NSAxLjA3MjI3IDUuMTA5MzggMS4wNzIyNyA0LjczMDQ3QzEuMDcyMjcgNC4zOTg0NCAxLjEyODkxIDQuMDk1NyAxLjI0MjE5IDMuODIyMjdDMS4zNTU0NyAzLjU0NDkyIDEuNTE1NjIgMy4zMDQ2OSAxLjcyMjY2IDMuMTAxNTZDMS45Mjk2OSAyLjg5ODQ0IDIuMTc5NjkgMi43MzQzNyAyLjQ3MjY2IDIuNjA5MzhDMi43NjU2MiAyLjQ4NDM4IDMuMDkxOCAyLjQwNDMgMy40NTExNyAyLjM2OTE0VjEuMTA5MzhINC4zODg2N1YyLjM4MDg2QzQuNzQwMjMgMi40Mjc3MyA1LjA1NjY0IDIuNTIzNDQgNS4zMzc4OSAyLjY2Nzk3QzUuNjE5MTQgMi44MTI1IDUuODU3NDIgMy4wMDE5NSA2LjA1MjczIDMuMjM2MzNDNi4yNTE5NSAzLjQ2NjggNi40MDQzIDMuNzQwMjMgNi41MDk3NyA0LjA1NjY0QzYuNjE5MTQgNC4zNjkxNCA2LjY3MzgzIDQuNzIwNyA2LjY3MzgzIDUuMTExMzNINS4wNDQ5MkM1LjA0NDkyIDQuNjM4NjcgNC45Mzc1IDQuMjgxMjUgNC43MjI2NiA0LjAzOTA2QzQuNTA3ODEgMy43OTI5NyA0LjIxNjggMy42Njk5MiAzLjg0OTYxIDMuNjY5OTJDMy42NTAzOSAzLjY2OTkyIDMuNDc2NTYgMy42OTcyNyAzLjMyODEyIDMuNzUxOTVDMy4xODM1OSAzLjgwMjczIDMuMDY0NDUgMy44NzY5NSAyLjk3MDcgMy45NzQ2MUMyLjg3Njk1IDQuMDY4MzYgMi44MDY2NCA0LjE3OTY5IDIuNzU5NzcgNC4zMDg1OUMyLjcxNjggNC40Mzc1IDIuNjk1MzEgNC41NzgxMiAyLjY5NTMxIDQuNzMwNDdDMi42OTUzMSA0Ljg4MjgxIDIuNzE2OCA1LjAxOTUzIDIuNzU5NzcgNS4xNDA2MkMyLjgwNjY0IDUuMjU3ODEgMi44ODI4MSA1LjM2NzE5IDIuOTg4MjggNS40Njg3NUMzLjA5NzY2IDUuNTcwMzEgMy4yNDAyMyA1LjY2Nzk3IDMuNDE2MDIgNS43NjE3MkMzLjU5MTggNS44NTE1NiAzLjgxMDU1IDUuOTQzMzYgNC4wNzIyNyA2LjAzNzExQzQuNDY2OCA2LjE4NTU1IDQuODI0MjIgNi4zMzk4NCA1LjE0NDUzIDYuNUM1LjQ2NDg0IDYuNjU2MjUgNS43MzgyOCA2LjgzOTg0IDUuOTY0ODQgNy4wNTA3OEM2LjE5NTMxIDcuMjU3ODEgNi4zNzEwOSA3LjUgNi40OTIxOSA3Ljc3NzM0QzYuNjE3MTkgOC4wNTA3OCA2LjY3OTY5IDguMzc1IDYuNjc5NjkgOC43NUM2LjY3OTY5IDkuMDkzNzUgNi42MjMwNSA5LjQwNDMgNi41MDk3NyA5LjY4MTY0QzYuMzk2NDggOS45NTUwOCA2LjIzNDM4IDEwLjE5MTQgNi4wMjM0NCAxMC4zOTA2QzUuODEyNSAxMC41ODk4IDUuNTU4NTkgMTAuNzUgNS4yNjE3MiAxMC44NzExQzQuOTY0ODQgMTAuOTg4MyA0LjYzMjgxIDExLjA2NDUgNC4yNjU2MiAxMS4wOTk2VjEyLjI0OEgzLjMzMzk4VjExLjA5OTZDMy4wMDE5NSAxMS4wNjg0IDIuNjc5NjkgMTAuOTk2MSAyLjM2NzE5IDEwLjg4MjhDMi4wNTQ2OSAxMC43NjU2IDEuNzc3MzQgMTAuNTk3NyAxLjUzNTE2IDEwLjM3ODlDMS4yOTY4OCAxMC4xNjAyIDEuMTA1NDcgOS44ODQ3NyAwLjk2MDkzOCA5LjU1MjczQzAuODE2NDA2IDkuMjE2OCAwLjc0NDE0MSA4LjgxNDQ1IDAuNzQ0MTQxIDguMzQ1N0gyLjM3ODkxQzIuMzc4OTEgOC42MjY5NSAyLjQxOTkyIDguODYzMjggMi41MDE5NSA5LjA1NDY5QzIuNTgzOTggOS4yNDIxOSAyLjY4OTQ1IDkuMzkyNTggMi44MTgzNiA5LjUwNTg2QzIuOTUxMTcgOS42MTUyMyAzLjEwMTU2IDkuNjkzMzYgMy4yNjk1MyA5Ljc0MDIzQzMuNDM3NSA5Ljc4NzExIDMuNjA5MzggOS44MTA1NSAzLjc4NTE2IDkuODEwNTVDNC4yMDMxMiA5LjgxMDU1IDQuNTE5NTMgOS43MTI4OSA0LjczNDM4IDkuNTE3NThDNC45NDkyMiA5LjMyMjI3IDUuMDU2NjQgOS4wNzAzMSA1LjA1NjY0IDguNzYxNzJaTTEzLjQxOCAxMi4yNzE1SDguMDc0MjJWMTFIMTMuNDE4VjEyLjI3MTVaIiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgzLjk1MjY0IDYpIiBmaWxsPSJ3aGl0ZSIvPgo8L3N2Zz4K);
	  --jp-icon-text-editor: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI0Ij4KICA8cGF0aCBjbGFzcz0ianAtaWNvbjMganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjNjE2MTYxIiBkPSJNMTUgMTVIM3YyaDEydi0yem0wLThIM3YyaDEyVjd6TTMgMTNoMTh2LTJIM3Yyem0wIDhoMTh2LTJIM3Yyek0zIDN2MmgxOFYzSDN6Ii8+Cjwvc3ZnPgo=);
	  --jp-icon-toc: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB2ZXJzaW9uPSIxLjEiIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgoJPHBhdGggZD0iTTcsNUgyMVY3SDdWNU03LDEzVjExSDIxVjEzSDdNNCw0LjVBMS41LDEuNSAwIDAsMSA1LjUsNkExLjUsMS41IDAgMCwxIDQsNy41QTEuNSwxLjUgMCAwLDEgMi41LDZBMS41LDEuNSAwIDAsMSA0LDQuNU00LDEwLjVBMS41LDEuNSAwIDAsMSA1LjUsMTJBMS41LDEuNSAwIDAsMSA0LDEzLjVBMS41LDEuNSAwIDAsMSAyLjUsMTJBMS41LDEuNSAwIDAsMSA0LDEwLjVNNywxOVYxN0gyMVYxOUg3TTQsMTYuNUExLjUsMS41IDAgMCwxIDUuNSwxOEExLjUsMS41IDAgMCwxIDQsMTkuNUExLjUsMS41IDAgMCwxIDIuNSwxOEExLjUsMS41IDAgMCwxIDQsMTYuNVoiIC8+Cjwvc3ZnPgo=);
	  --jp-icon-tree-view: url(data:image/svg+xml;base64,PHN2ZyBoZWlnaHQ9IjI0IiB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICAgIDxnIGNsYXNzPSJqcC1pY29uMyIgZmlsbD0iIzYxNjE2MSI+CiAgICAgICAgPHBhdGggZD0iTTAgMGgyNHYyNEgweiIgZmlsbD0ibm9uZSIvPgogICAgICAgIDxwYXRoIGQ9Ik0yMiAxMVYzaC03djNIOVYzSDJ2OGg3VjhoMnYxMGg0djNoN3YtOGgtN3YzaC0yVjhoMnYzeiIvPgogICAgPC9nPgo8L3N2Zz4=);
	  --jp-icon-trusted: url(data:image/svg+xml;base64,PHN2ZyBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDI0IDI1Ij4KICAgIDxwYXRoIGNsYXNzPSJqcC1pY29uMiIgc3Ryb2tlPSIjMzMzMzMzIiBzdHJva2Utd2lkdGg9IjIiIHRyYW5zZm9ybT0idHJhbnNsYXRlKDIgMykiIGQ9Ik0xLjg2MDk0IDExLjQ0MDlDMC44MjY0NDggOC43NzAyNyAwLjg2Mzc3OSA2LjA1NzY0IDEuMjQ5MDcgNC4xOTkzMkMyLjQ4MjA2IDMuOTMzNDcgNC4wODA2OCAzLjQwMzQ3IDUuNjAxMDIgMi44NDQ5QzcuMjM1NDkgMi4yNDQ0IDguODU2NjYgMS41ODE1IDkuOTg3NiAxLjA5NTM5QzExLjA1OTcgMS41ODM0MSAxMi42MDk0IDIuMjQ0NCAxNC4yMTggMi44NDMzOUMxNS43NTAzIDMuNDEzOTQgMTcuMzk5NSAzLjk1MjU4IDE4Ljc1MzkgNC4yMTM4NUMxOS4xMzY0IDYuMDcxNzcgMTkuMTcwOSA4Ljc3NzIyIDE4LjEzOSAxMS40NDA5QzE3LjAzMDMgMTQuMzAzMiAxNC42NjY4IDE3LjE4NDQgOS45OTk5OSAxOC45MzU0QzUuMzMzMiAxNy4xODQ0IDIuOTY5NjggMTQuMzAzMiAxLjg2MDk0IDExLjQ0MDlaIi8+CiAgICA8cGF0aCBjbGFzcz0ianAtaWNvbjIiIGZpbGw9IiMzMzMzMzMiIHN0cm9rZT0iIzMzMzMzMyIgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoOCA5Ljg2NzE5KSIgZD0iTTIuODYwMTUgNC44NjUzNUwwLjcyNjU0OSAyLjk5OTU5TDAgMy42MzA0NUwyLjg2MDE1IDYuMTMxNTdMOCAwLjYzMDg3Mkw3LjI3ODU3IDBMMi44NjAxNSA0Ljg2NTM1WiIvPgo8L3N2Zz4K);
	  --jp-icon-undo: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMjQgMjQiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEyLjUgOGMtMi42NSAwLTUuMDUuOTktNi45IDIuNkwyIDd2OWg5bC0zLjYyLTMuNjJjMS4zOS0xLjE2IDMuMTYtMS44OCA1LjEyLTEuODggMy41NCAwIDYuNTUgMi4zMSA3LjYgNS41bDIuMzctLjc4QzIxLjA4IDExLjAzIDE3LjE1IDggMTIuNSA4eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-vega: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8ZyBjbGFzcz0ianAtaWNvbjEganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjMjEyMTIxIj4KICAgIDxwYXRoIGQ9Ik0xMC42IDUuNGwyLjItMy4ySDIuMnY3LjNsNC02LjZ6Ii8+CiAgICA8cGF0aCBkPSJNMTUuOCAyLjJsLTQuNCA2LjZMNyA2LjNsLTQuOCA4djUuNWgxNy42VjIuMmgtNHptLTcgMTUuNEg1LjV2LTQuNGgzLjN2NC40em00LjQgMEg5LjhWOS44aDMuNHY3Ljh6bTQuNCAwaC0zLjRWNi41aDMuNHYxMS4xeiIvPgogIDwvZz4KPC9zdmc+Cg==);
	  --jp-icon-yaml: url(data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxNiIgdmlld0JveD0iMCAwIDIyIDIyIj4KICA8ZyBjbGFzcz0ianAtaWNvbi1jb250cmFzdDIganAtaWNvbi1zZWxlY3RhYmxlIiBmaWxsPSIjRDgxQjYwIj4KICAgIDxwYXRoIGQ9Ik03LjIgMTguNnYtNS40TDMgNS42aDMuM2wxLjQgMy4xYy4zLjkuNiAxLjYgMSAyLjUuMy0uOC42LTEuNiAxLTIuNWwxLjQtMy4xaDMuNGwtNC40IDcuNnY1LjVsLTIuOS0uMXoiLz4KICAgIDxjaXJjbGUgY2xhc3M9InN0MCIgY3g9IjE3LjYiIGN5PSIxNi41IiByPSIyLjEiLz4KICAgIDxjaXJjbGUgY2xhc3M9InN0MCIgY3g9IjE3LjYiIGN5PSIxMSIgcj0iMi4xIi8+CiAgPC9nPgo8L3N2Zz4K);
	}
	
	/* Icon CSS class declarations */
	
	.jp-AddIcon {
	  background-image: var(--jp-icon-add);
	}
	.jp-BugIcon {
	  background-image: var(--jp-icon-bug);
	}
	.jp-BuildIcon {
	  background-image: var(--jp-icon-build);
	}
	.jp-CaretDownEmptyIcon {
	  background-image: var(--jp-icon-caret-down-empty);
	}
	.jp-CaretDownEmptyThinIcon {
	  background-image: var(--jp-icon-caret-down-empty-thin);
	}
	.jp-CaretDownIcon {
	  background-image: var(--jp-icon-caret-down);
	}
	.jp-CaretLeftIcon {
	  background-image: var(--jp-icon-caret-left);
	}
	.jp-CaretRightIcon {
	  background-image: var(--jp-icon-caret-right);
	}
	.jp-CaretUpEmptyThinIcon {
	  background-image: var(--jp-icon-caret-up-empty-thin);
	}
	.jp-CaretUpIcon {
	  background-image: var(--jp-icon-caret-up);
	}
	.jp-CaseSensitiveIcon {
	  background-image: var(--jp-icon-case-sensitive);
	}
	.jp-CheckIcon {
	  background-image: var(--jp-icon-check);
	}
	.jp-CircleEmptyIcon {
	  background-image: var(--jp-icon-circle-empty);
	}
	.jp-CircleIcon {
	  background-image: var(--jp-icon-circle);
	}
	.jp-ClearIcon {
	  background-image: var(--jp-icon-clear);
	}
	.jp-CloseIcon {
	  background-image: var(--jp-icon-close);
	}
	.jp-CodeIcon {
	  background-image: var(--jp-icon-code);
	}
	.jp-ConsoleIcon {
	  background-image: var(--jp-icon-console);
	}
	.jp-CopyIcon {
	  background-image: var(--jp-icon-copy);
	}
	.jp-CutIcon {
	  background-image: var(--jp-icon-cut);
	}
	.jp-DownloadIcon {
	  background-image: var(--jp-icon-download);
	}
	.jp-EditIcon {
	  background-image: var(--jp-icon-edit);
	}
	.jp-EllipsesIcon {
	  background-image: var(--jp-icon-ellipses);
	}
	.jp-ExtensionIcon {
	  background-image: var(--jp-icon-extension);
	}
	.jp-FastForwardIcon {
	  background-image: var(--jp-icon-fast-forward);
	}
	.jp-FileIcon {
	  background-image: var(--jp-icon-file);
	}
	.jp-FileUploadIcon {
	  background-image: var(--jp-icon-file-upload);
	}
	.jp-FilterListIcon {
	  background-image: var(--jp-icon-filter-list);
	}
	.jp-FolderIcon {
	  background-image: var(--jp-icon-folder);
	}
	.jp-Html5Icon {
	  background-image: var(--jp-icon-html5);
	}
	.jp-ImageIcon {
	  background-image: var(--jp-icon-image);
	}
	.jp-InspectorIcon {
	  background-image: var(--jp-icon-inspector);
	}
	.jp-JsonIcon {
	  background-image: var(--jp-icon-json);
	}
	.jp-JupyterFaviconIcon {
	  background-image: var(--jp-icon-jupyter-favicon);
	}
	.jp-JupyterIcon {
	  background-image: var(--jp-icon-jupyter);
	}
	.jp-JupyterlabWordmarkIcon {
	  background-image: var(--jp-icon-jupyterlab-wordmark);
	}
	.jp-KernelIcon {
	  background-image: var(--jp-icon-kernel);
	}
	.jp-KeyboardIcon {
	  background-image: var(--jp-icon-keyboard);
	}
	.jp-LauncherIcon {
	  background-image: var(--jp-icon-launcher);
	}
	.jp-LineFormIcon {
	  background-image: var(--jp-icon-line-form);
	}
	.jp-LinkIcon {
	  background-image: var(--jp-icon-link);
	}
	.jp-ListIcon {
	  background-image: var(--jp-icon-list);
	}
	.jp-ListingsInfoIcon {
	  background-image: var(--jp-icon-listings-info);
	}
	.jp-MarkdownIcon {
	  background-image: var(--jp-icon-markdown);
	}
	.jp-NewFolderIcon {
	  background-image: var(--jp-icon-new-folder);
	}
	.jp-NotTrustedIcon {
	  background-image: var(--jp-icon-not-trusted);
	}
	.jp-NotebookIcon {
	  background-image: var(--jp-icon-notebook);
	}
	.jp-NumberingIcon {
	  background-image: var(--jp-icon-numbering);
	}
	.jp-OfflineBoltIcon {
	  background-image: var(--jp-icon-offline-bolt);
	}
	.jp-PaletteIcon {
	  background-image: var(--jp-icon-palette);
	}
	.jp-PasteIcon {
	  background-image: var(--jp-icon-paste);
	}
	.jp-PdfIcon {
	  background-image: var(--jp-icon-pdf);
	}
	.jp-PythonIcon {
	  background-image: var(--jp-icon-python);
	}
	.jp-RKernelIcon {
	  background-image: var(--jp-icon-r-kernel);
	}
	.jp-ReactIcon {
	  background-image: var(--jp-icon-react);
	}
	.jp-RedoIcon {
	  background-image: var(--jp-icon-redo);
	}
	.jp-RefreshIcon {
	  background-image: var(--jp-icon-refresh);
	}
	.jp-RegexIcon {
	  background-image: var(--jp-icon-regex);
	}
	.jp-RunIcon {
	  background-image: var(--jp-icon-run);
	}
	.jp-RunningIcon {
	  background-image: var(--jp-icon-running);
	}
	.jp-SaveIcon {
	  background-image: var(--jp-icon-save);
	}
	.jp-SearchIcon {
	  background-image: var(--jp-icon-search);
	}
	.jp-SettingsIcon {
	  background-image: var(--jp-icon-settings);
	}
	.jp-SpreadsheetIcon {
	  background-image: var(--jp-icon-spreadsheet);
	}
	.jp-StopIcon {
	  background-image: var(--jp-icon-stop);
	}
	.jp-TabIcon {
	  background-image: var(--jp-icon-tab);
	}
	.jp-TableRowsIcon {
	  background-image: var(--jp-icon-table-rows);
	}
	.jp-TagIcon {
	  background-image: var(--jp-icon-tag);
	}
	.jp-TerminalIcon {
	  background-image: var(--jp-icon-terminal);
	}
	.jp-TextEditorIcon {
	  background-image: var(--jp-icon-text-editor);
	}
	.jp-TocIcon {
	  background-image: var(--jp-icon-toc);
	}
	.jp-TreeViewIcon {
	  background-image: var(--jp-icon-tree-view);
	}
	.jp-TrustedIcon {
	  background-image: var(--jp-icon-trusted);
	}
	.jp-UndoIcon {
	  background-image: var(--jp-icon-undo);
	}
	.jp-VegaIcon {
	  background-image: var(--jp-icon-vega);
	}
	.jp-YamlIcon {
	  background-image: var(--jp-icon-yaml);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/**
	 * (DEPRECATED) Support for consuming icons as CSS background images
	 */
	
	:root {
	  --jp-icon-search-white: url(data:image/svg+xml;base64,PHN2ZyB2aWV3Qm94PSIwIDAgMTggMTgiIHdpZHRoPSIxNiIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZyBjbGFzcz0ianAtaWNvbjMiIGZpbGw9IiM2MTYxNjEiPgogICAgPHBhdGggZD0iTTEyLjEsMTAuOWgtMC43bC0wLjItMC4yYzAuOC0wLjksMS4zLTIuMiwxLjMtMy41YzAtMy0yLjQtNS40LTUuNC01LjRTMS44LDQuMiwxLjgsNy4xczIuNCw1LjQsNS40LDUuNCBjMS4zLDAsMi41LTAuNSwzLjUtMS4zbDAuMiwwLjJ2MC43bDQuMSw0LjFsMS4yLTEuMkwxMi4xLDEwLjl6IE03LjEsMTAuOWMtMi4xLDAtMy43LTEuNy0zLjctMy43czEuNy0zLjcsMy43LTMuN3MzLjcsMS43LDMuNywzLjcgUzkuMiwxMC45LDcuMSwxMC45eiIvPgogIDwvZz4KPC9zdmc+Cg==);
	}
	
	.jp-Icon,
	.jp-MaterialIcon {
	  background-position: center;
	  background-repeat: no-repeat;
	  background-size: 16px;
	  min-width: 16px;
	  min-height: 16px;
	}
	
	.jp-Icon-cover {
	  background-position: center;
	  background-repeat: no-repeat;
	  background-size: cover;
	}
	
	/**
	 * (DEPRECATED) Support for specific CSS icon sizes
	 */
	
	.jp-Icon-16 {
	  background-size: 16px;
	  min-width: 16px;
	  min-height: 16px;
	}
	
	.jp-Icon-18 {
	  background-size: 18px;
	  min-width: 18px;
	  min-height: 18px;
	}
	
	.jp-Icon-20 {
	  background-size: 20px;
	  min-width: 20px;
	  min-height: 20px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/**
	 * Support for icons as inline SVG HTMLElements
	 */
	
	/* recolor the primary elements of an icon */
	.jp-icon0[fill] {
	  fill: var(--jp-inverse-layout-color0);
	}
	.jp-icon1[fill] {
	  fill: var(--jp-inverse-layout-color1);
	}
	.jp-icon2[fill] {
	  fill: var(--jp-inverse-layout-color2);
	}
	.jp-icon3[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	.jp-icon4[fill] {
	  fill: var(--jp-inverse-layout-color4);
	}
	
	.jp-icon0[stroke] {
	  stroke: var(--jp-inverse-layout-color0);
	}
	.jp-icon1[stroke] {
	  stroke: var(--jp-inverse-layout-color1);
	}
	.jp-icon2[stroke] {
	  stroke: var(--jp-inverse-layout-color2);
	}
	.jp-icon3[stroke] {
	  stroke: var(--jp-inverse-layout-color3);
	}
	.jp-icon4[stroke] {
	  stroke: var(--jp-inverse-layout-color4);
	}
	/* recolor the accent elements of an icon */
	.jp-icon-accent0[fill] {
	  fill: var(--jp-layout-color0);
	}
	.jp-icon-accent1[fill] {
	  fill: var(--jp-layout-color1);
	}
	.jp-icon-accent2[fill] {
	  fill: var(--jp-layout-color2);
	}
	.jp-icon-accent3[fill] {
	  fill: var(--jp-layout-color3);
	}
	.jp-icon-accent4[fill] {
	  fill: var(--jp-layout-color4);
	}
	
	.jp-icon-accent0[stroke] {
	  stroke: var(--jp-layout-color0);
	}
	.jp-icon-accent1[stroke] {
	  stroke: var(--jp-layout-color1);
	}
	.jp-icon-accent2[stroke] {
	  stroke: var(--jp-layout-color2);
	}
	.jp-icon-accent3[stroke] {
	  stroke: var(--jp-layout-color3);
	}
	.jp-icon-accent4[stroke] {
	  stroke: var(--jp-layout-color4);
	}
	/* set the color of an icon to transparent */
	.jp-icon-none[fill] {
	  fill: none;
	}
	
	.jp-icon-none[stroke] {
	  stroke: none;
	}
	/* brand icon colors. Same for light and dark */
	.jp-icon-brand0[fill] {
	  fill: var(--jp-brand-color0);
	}
	.jp-icon-brand1[fill] {
	  fill: var(--jp-brand-color1);
	}
	.jp-icon-brand2[fill] {
	  fill: var(--jp-brand-color2);
	}
	.jp-icon-brand3[fill] {
	  fill: var(--jp-brand-color3);
	}
	.jp-icon-brand4[fill] {
	  fill: var(--jp-brand-color4);
	}
	
	.jp-icon-brand0[stroke] {
	  stroke: var(--jp-brand-color0);
	}
	.jp-icon-brand1[stroke] {
	  stroke: var(--jp-brand-color1);
	}
	.jp-icon-brand2[stroke] {
	  stroke: var(--jp-brand-color2);
	}
	.jp-icon-brand3[stroke] {
	  stroke: var(--jp-brand-color3);
	}
	.jp-icon-brand4[stroke] {
	  stroke: var(--jp-brand-color4);
	}
	/* warn icon colors. Same for light and dark */
	.jp-icon-warn0[fill] {
	  fill: var(--jp-warn-color0);
	}
	.jp-icon-warn1[fill] {
	  fill: var(--jp-warn-color1);
	}
	.jp-icon-warn2[fill] {
	  fill: var(--jp-warn-color2);
	}
	.jp-icon-warn3[fill] {
	  fill: var(--jp-warn-color3);
	}
	
	.jp-icon-warn0[stroke] {
	  stroke: var(--jp-warn-color0);
	}
	.jp-icon-warn1[stroke] {
	  stroke: var(--jp-warn-color1);
	}
	.jp-icon-warn2[stroke] {
	  stroke: var(--jp-warn-color2);
	}
	.jp-icon-warn3[stroke] {
	  stroke: var(--jp-warn-color3);
	}
	/* icon colors that contrast well with each other and most backgrounds */
	.jp-icon-contrast0[fill] {
	  fill: var(--jp-icon-contrast-color0);
	}
	.jp-icon-contrast1[fill] {
	  fill: var(--jp-icon-contrast-color1);
	}
	.jp-icon-contrast2[fill] {
	  fill: var(--jp-icon-contrast-color2);
	}
	.jp-icon-contrast3[fill] {
	  fill: var(--jp-icon-contrast-color3);
	}
	
	.jp-icon-contrast0[stroke] {
	  stroke: var(--jp-icon-contrast-color0);
	}
	.jp-icon-contrast1[stroke] {
	  stroke: var(--jp-icon-contrast-color1);
	}
	.jp-icon-contrast2[stroke] {
	  stroke: var(--jp-icon-contrast-color2);
	}
	.jp-icon-contrast3[stroke] {
	  stroke: var(--jp-icon-contrast-color3);
	}
	
	/* CSS for icons in selected items in the settings editor */
	#setting-editor .jp-PluginList .jp-mod-selected .jp-icon-selectable[fill] {
	  fill: #fff;
	}
	#setting-editor
	  .jp-PluginList
	  .jp-mod-selected
	  .jp-icon-selectable-inverse[fill] {
	  fill: var(--jp-brand-color1);
	}
	
	/* CSS for icons in selected filebrowser listing items */
	.jp-DirListing-item.jp-mod-selected .jp-icon-selectable[fill] {
	  fill: #fff;
	}
	.jp-DirListing-item.jp-mod-selected .jp-icon-selectable-inverse[fill] {
	  fill: var(--jp-brand-color1);
	}
	
	/* CSS for icons in selected tabs in the sidebar tab manager */
	#tab-manager .lm-TabBar-tab.jp-mod-active .jp-icon-selectable[fill] {
	  fill: #fff;
	}
	
	#tab-manager .lm-TabBar-tab.jp-mod-active .jp-icon-selectable-inverse[fill] {
	  fill: var(--jp-brand-color1);
	}
	#tab-manager
	  .lm-TabBar-tab.jp-mod-active
	  .jp-icon-hover
	  :hover
	  .jp-icon-selectable[fill] {
	  fill: var(--jp-brand-color1);
	}
	
	#tab-manager
	  .lm-TabBar-tab.jp-mod-active
	  .jp-icon-hover
	  :hover
	  .jp-icon-selectable-inverse[fill] {
	  fill: #fff;
	}
	
	/**
	 * TODO: come up with non css-hack solution for showing the busy icon on top
	 *  of the close icon
	 * CSS for complex behavior of close icon of tabs in the sidebar tab manager
	 */
	#tab-manager
	  .lm-TabBar-tab.jp-mod-dirty
	  > .lm-TabBar-tabCloseIcon
	  > :not(:hover)
	  > .jp-icon3[fill] {
	  fill: none;
	}
	#tab-manager
	  .lm-TabBar-tab.jp-mod-dirty
	  > .lm-TabBar-tabCloseIcon
	  > :not(:hover)
	  > .jp-icon-busy[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	
	#tab-manager
	  .lm-TabBar-tab.jp-mod-dirty.jp-mod-active
	  > .lm-TabBar-tabCloseIcon
	  > :not(:hover)
	  > .jp-icon-busy[fill] {
	  fill: #fff;
	}
	
	/**
	* TODO: come up with non css-hack solution for showing the busy icon on top
	*  of the close icon
	* CSS for complex behavior of close icon of tabs in the main area tabbar
	*/
	.lm-DockPanel-tabBar
	  .lm-TabBar-tab.lm-mod-closable.jp-mod-dirty
	  > .lm-TabBar-tabCloseIcon
	  > :not(:hover)
	  > .jp-icon3[fill] {
	  fill: none;
	}
	.lm-DockPanel-tabBar
	  .lm-TabBar-tab.lm-mod-closable.jp-mod-dirty
	  > .lm-TabBar-tabCloseIcon
	  > :not(:hover)
	  > .jp-icon-busy[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	
	/* CSS for icons in status bar */
	#jp-main-statusbar .jp-mod-selected .jp-icon-selectable[fill] {
	  fill: #fff;
	}
	
	#jp-main-statusbar .jp-mod-selected .jp-icon-selectable-inverse[fill] {
	  fill: var(--jp-brand-color1);
	}
	/* special handling for splash icon CSS. While the theme CSS reloads during
	   splash, the splash icon can loose theming. To prevent that, we set a
	   default for its color variable */
	:root {
	  --jp-warn-color0: var(--md-orange-700);
	}
	
	/* not sure what to do with this one, used in filebrowser listing */
	.jp-DragIcon {
	  margin-right: 4px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/**
	 * Support for alt colors for icons as inline SVG HTMLElements
	 */
	
	/* alt recolor the primary elements of an icon */
	.jp-icon-alt .jp-icon0[fill] {
	  fill: var(--jp-layout-color0);
	}
	.jp-icon-alt .jp-icon1[fill] {
	  fill: var(--jp-layout-color1);
	}
	.jp-icon-alt .jp-icon2[fill] {
	  fill: var(--jp-layout-color2);
	}
	.jp-icon-alt .jp-icon3[fill] {
	  fill: var(--jp-layout-color3);
	}
	.jp-icon-alt .jp-icon4[fill] {
	  fill: var(--jp-layout-color4);
	}
	
	.jp-icon-alt .jp-icon0[stroke] {
	  stroke: var(--jp-layout-color0);
	}
	.jp-icon-alt .jp-icon1[stroke] {
	  stroke: var(--jp-layout-color1);
	}
	.jp-icon-alt .jp-icon2[stroke] {
	  stroke: var(--jp-layout-color2);
	}
	.jp-icon-alt .jp-icon3[stroke] {
	  stroke: var(--jp-layout-color3);
	}
	.jp-icon-alt .jp-icon4[stroke] {
	  stroke: var(--jp-layout-color4);
	}
	
	/* alt recolor the accent elements of an icon */
	.jp-icon-alt .jp-icon-accent0[fill] {
	  fill: var(--jp-inverse-layout-color0);
	}
	.jp-icon-alt .jp-icon-accent1[fill] {
	  fill: var(--jp-inverse-layout-color1);
	}
	.jp-icon-alt .jp-icon-accent2[fill] {
	  fill: var(--jp-inverse-layout-color2);
	}
	.jp-icon-alt .jp-icon-accent3[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	.jp-icon-alt .jp-icon-accent4[fill] {
	  fill: var(--jp-inverse-layout-color4);
	}
	
	.jp-icon-alt .jp-icon-accent0[stroke] {
	  stroke: var(--jp-inverse-layout-color0);
	}
	.jp-icon-alt .jp-icon-accent1[stroke] {
	  stroke: var(--jp-inverse-layout-color1);
	}
	.jp-icon-alt .jp-icon-accent2[stroke] {
	  stroke: var(--jp-inverse-layout-color2);
	}
	.jp-icon-alt .jp-icon-accent3[stroke] {
	  stroke: var(--jp-inverse-layout-color3);
	}
	.jp-icon-alt .jp-icon-accent4[stroke] {
	  stroke: var(--jp-inverse-layout-color4);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-icon-hoverShow:not(:hover) svg {
	  display: none !important;
	}
	
	/**
	 * Support for hover colors for icons as inline SVG HTMLElements
	 */
	
	/**
	 * regular colors
	 */
	
	/* recolor the primary elements of an icon */
	.jp-icon-hover :hover .jp-icon0-hover[fill] {
	  fill: var(--jp-inverse-layout-color0);
	}
	.jp-icon-hover :hover .jp-icon1-hover[fill] {
	  fill: var(--jp-inverse-layout-color1);
	}
	.jp-icon-hover :hover .jp-icon2-hover[fill] {
	  fill: var(--jp-inverse-layout-color2);
	}
	.jp-icon-hover :hover .jp-icon3-hover[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	.jp-icon-hover :hover .jp-icon4-hover[fill] {
	  fill: var(--jp-inverse-layout-color4);
	}
	
	.jp-icon-hover :hover .jp-icon0-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color0);
	}
	.jp-icon-hover :hover .jp-icon1-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color1);
	}
	.jp-icon-hover :hover .jp-icon2-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color2);
	}
	.jp-icon-hover :hover .jp-icon3-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color3);
	}
	.jp-icon-hover :hover .jp-icon4-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color4);
	}
	
	/* recolor the accent elements of an icon */
	.jp-icon-hover :hover .jp-icon-accent0-hover[fill] {
	  fill: var(--jp-layout-color0);
	}
	.jp-icon-hover :hover .jp-icon-accent1-hover[fill] {
	  fill: var(--jp-layout-color1);
	}
	.jp-icon-hover :hover .jp-icon-accent2-hover[fill] {
	  fill: var(--jp-layout-color2);
	}
	.jp-icon-hover :hover .jp-icon-accent3-hover[fill] {
	  fill: var(--jp-layout-color3);
	}
	.jp-icon-hover :hover .jp-icon-accent4-hover[fill] {
	  fill: var(--jp-layout-color4);
	}
	
	.jp-icon-hover :hover .jp-icon-accent0-hover[stroke] {
	  stroke: var(--jp-layout-color0);
	}
	.jp-icon-hover :hover .jp-icon-accent1-hover[stroke] {
	  stroke: var(--jp-layout-color1);
	}
	.jp-icon-hover :hover .jp-icon-accent2-hover[stroke] {
	  stroke: var(--jp-layout-color2);
	}
	.jp-icon-hover :hover .jp-icon-accent3-hover[stroke] {
	  stroke: var(--jp-layout-color3);
	}
	.jp-icon-hover :hover .jp-icon-accent4-hover[stroke] {
	  stroke: var(--jp-layout-color4);
	}
	
	/* set the color of an icon to transparent */
	.jp-icon-hover :hover .jp-icon-none-hover[fill] {
	  fill: none;
	}
	
	.jp-icon-hover :hover .jp-icon-none-hover[stroke] {
	  stroke: none;
	}
	
	/**
	 * inverse colors
	 */
	
	/* inverse recolor the primary elements of an icon */
	.jp-icon-hover.jp-icon-alt :hover .jp-icon0-hover[fill] {
	  fill: var(--jp-layout-color0);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon1-hover[fill] {
	  fill: var(--jp-layout-color1);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon2-hover[fill] {
	  fill: var(--jp-layout-color2);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon3-hover[fill] {
	  fill: var(--jp-layout-color3);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon4-hover[fill] {
	  fill: var(--jp-layout-color4);
	}
	
	.jp-icon-hover.jp-icon-alt :hover .jp-icon0-hover[stroke] {
	  stroke: var(--jp-layout-color0);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon1-hover[stroke] {
	  stroke: var(--jp-layout-color1);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon2-hover[stroke] {
	  stroke: var(--jp-layout-color2);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon3-hover[stroke] {
	  stroke: var(--jp-layout-color3);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon4-hover[stroke] {
	  stroke: var(--jp-layout-color4);
	}
	
	/* inverse recolor the accent elements of an icon */
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent0-hover[fill] {
	  fill: var(--jp-inverse-layout-color0);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent1-hover[fill] {
	  fill: var(--jp-inverse-layout-color1);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent2-hover[fill] {
	  fill: var(--jp-inverse-layout-color2);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent3-hover[fill] {
	  fill: var(--jp-inverse-layout-color3);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent4-hover[fill] {
	  fill: var(--jp-inverse-layout-color4);
	}
	
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent0-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color0);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent1-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color1);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent2-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color2);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent3-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color3);
	}
	.jp-icon-hover.jp-icon-alt :hover .jp-icon-accent4-hover[stroke] {
	  stroke: var(--jp-inverse-layout-color4);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-switch {
	  display: flex;
	  align-items: center;
	  padding-left: 4px;
	  padding-right: 4px;
	  font-size: var(--jp-ui-font-size1);
	  background-color: transparent;
	  color: var(--jp-ui-font-color1);
	  border: none;
	  height: 20px;
	}
	
	.jp-switch:hover {
	  background-color: var(--jp-layout-color2);
	}
	
	.jp-switch-label {
	  margin-right: 5px;
	}
	
	.jp-switch-track {
	  cursor: pointer;
	  background-color: var(--jp-border-color1);
	  -webkit-transition: 0.4s;
	  transition: 0.4s;
	  border-radius: 34px;
	  height: 16px;
	  width: 35px;
	  position: relative;
	}
	
	.jp-switch-track::before {
	  content: '';
	  position: absolute;
	  height: 10px;
	  width: 10px;
	  margin: 3px;
	  left: 0px;
	  background-color: var(--jp-ui-inverse-font-color1);
	  -webkit-transition: 0.4s;
	  transition: 0.4s;
	  border-radius: 50%;
	}
	
	.jp-switch[aria-checked='true'] .jp-switch-track {
	  background-color: var(--jp-warn-color0);
	}
	
	.jp-switch[aria-checked='true'] .jp-switch-track::before {
	  /* track width (35) - margins (3 + 3) - thumb width (10) */
	  left: 19px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/* Sibling imports */
	
	/* Override Blueprint's _reset.scss styles */
	html {
	  box-sizing: unset;
	}
	
	*,
	*::before,
	*::after {
	  box-sizing: unset;
	}
	
	body {
	  color: unset;
	  font-family: var(--jp-ui-font-family);
	}
	
	p {
	  margin-top: unset;
	  margin-bottom: unset;
	}
	
	small {
	  font-size: unset;
	}
	
	strong {
	  font-weight: unset;
	}
	
	/* Override Blueprint's _typography.scss styles */
	a {
	  text-decoration: unset;
	  color: unset;
	}
	a:hover {
	  text-decoration: unset;
	  color: unset;
	}
	
	/* Override Blueprint's _accessibility.scss styles */
	:focus {
	  outline: unset;
	  outline-offset: unset;
	  -moz-outline-radius: unset;
	}
	
	/* Styles for ui-components */
	.jp-Button {
	  border-radius: var(--jp-border-radius);
	  padding: 0px 12px;
	  font-size: var(--jp-ui-font-size1);
	}
	
	/* Use our own theme for hover styles */
	button.jp-Button.bp3-button.bp3-minimal:hover {
	  background-color: var(--jp-layout-color2);
	}
	.jp-Button.minimal {
	  color: unset !important;
	}
	
	.jp-Button.jp-ToolbarButtonComponent {
	  text-transform: none;
	}
	
	.jp-InputGroup input {
	  box-sizing: border-box;
	  border-radius: 0;
	  background-color: transparent;
	  color: var(--jp-ui-font-color0);
	  box-shadow: inset 0 0 0 var(--jp-border-width) var(--jp-input-border-color);
	}
	
	.jp-InputGroup input:focus {
	  box-shadow: inset 0 0 0 var(--jp-border-width)
		  var(--jp-input-active-box-shadow-color),
		inset 0 0 0 3px var(--jp-input-active-box-shadow-color);
	}
	
	.jp-InputGroup input::placeholder,
	input::placeholder {
	  color: var(--jp-ui-font-color3);
	}
	
	.jp-BPIcon {
	  display: inline-block;
	  vertical-align: middle;
	  margin: auto;
	}
	
	/* Stop blueprint futzing with our icon fills */
	.bp3-icon.jp-BPIcon > svg:not([fill]) {
	  fill: var(--jp-inverse-layout-color3);
	}
	
	.jp-InputGroupAction {
	  padding: 6px;
	}
	
	.jp-HTMLSelect.jp-DefaultStyle select {
	  background-color: initial;
	  border: none;
	  border-radius: 0;
	  box-shadow: none;
	  color: var(--jp-ui-font-color0);
	  display: block;
	  font-size: var(--jp-ui-font-size1);
	  height: 24px;
	  line-height: 14px;
	  padding: 0 25px 0 10px;
	  text-align: left;
	  -moz-appearance: none;
	  -webkit-appearance: none;
	}
	
	/* Use our own theme for hover and option styles */
	.jp-HTMLSelect.jp-DefaultStyle select:hover,
	.jp-HTMLSelect.jp-DefaultStyle select > option {
	  background-color: var(--jp-layout-color2);
	  color: var(--jp-ui-font-color0);
	}
	select {
	  box-sizing: border-box;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-Collapse {
	  display: flex;
	  flex-direction: column;
	  align-items: stretch;
	  border-top: 1px solid var(--jp-border-color2);
	  border-bottom: 1px solid var(--jp-border-color2);
	}
	
	.jp-Collapse-header {
	  padding: 1px 12px;
	  color: var(--jp-ui-font-color1);
	  background-color: var(--jp-layout-color1);
	  font-size: var(--jp-ui-font-size2);
	}
	
	.jp-Collapse-header:hover {
	  background-color: var(--jp-layout-color2);
	}
	
	.jp-Collapse-contents {
	  padding: 0px 12px 0px 12px;
	  background-color: var(--jp-layout-color1);
	  color: var(--jp-ui-font-color1);
	  overflow: auto;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Variables
	|----------------------------------------------------------------------------*/
	
	:root {
	  --jp-private-commandpalette-search-height: 28px;
	}
	
	/*-----------------------------------------------------------------------------
	| Overall styles
	|----------------------------------------------------------------------------*/
	
	.lm-CommandPalette {
	  padding-bottom: 0px;
	  color: var(--jp-ui-font-color1);
	  background: var(--jp-layout-color1);
	  /* This is needed so that all font sizing of children done in ems is
	   * relative to this base size */
	  font-size: var(--jp-ui-font-size1);
	}
	
	/*-----------------------------------------------------------------------------
	| Modal variant
	|----------------------------------------------------------------------------*/
	
	.jp-ModalCommandPalette {
	  position: absolute;
	  z-index: 10000;
	  top: 38px;
	  left: 30%;
	  margin: 0;
	  padding: 4px;
	  width: 40%;
	  box-shadow: var(--jp-elevation-z4);
	  border-radius: 4px;
	  background: var(--jp-layout-color0);
	}
	
	.jp-ModalCommandPalette .lm-CommandPalette {
	  max-height: 40vh;
	}
	
	.jp-ModalCommandPalette .lm-CommandPalette .lm-close-icon::after {
	  display: none;
	}
	
	.jp-ModalCommandPalette .lm-CommandPalette .lm-CommandPalette-header {
	  display: none;
	}
	
	.jp-ModalCommandPalette .lm-CommandPalette .lm-CommandPalette-item {
	  margin-left: 4px;
	  margin-right: 4px;
	}
	
	.jp-ModalCommandPalette
	  .lm-CommandPalette
	  .lm-CommandPalette-item.lm-mod-disabled {
	  display: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Search
	|----------------------------------------------------------------------------*/
	
	.lm-CommandPalette-search {
	  padding: 4px;
	  background-color: var(--jp-layout-color1);
	  z-index: 2;
	}
	
	.lm-CommandPalette-wrapper {
	  overflow: overlay;
	  padding: 0px 9px;
	  background-color: var(--jp-input-active-background);
	  height: 30px;
	  box-shadow: inset 0 0 0 var(--jp-border-width) var(--jp-input-border-color);
	}
	
	.lm-CommandPalette.lm-mod-focused .lm-CommandPalette-wrapper {
	  box-shadow: inset 0 0 0 1px var(--jp-input-active-box-shadow-color),
		inset 0 0 0 3px var(--jp-input-active-box-shadow-color);
	}
	
	.lm-CommandPalette-wrapper::after {
	  content: ' ';
	  color: white;
	  background-color: var(--jp-brand-color1);
	  position: absolute;
	  top: 4px;
	  right: 4px;
	  height: 30px;
	  width: 10px;
	  padding: 0px 10px;
	  background-image: var(--jp-icon-search-white);
	  background-size: 20px;
	  background-repeat: no-repeat;
	  background-position: center;
	}
	
	.lm-CommandPalette-input {
	  background: transparent;
	  width: calc(100% - 18px);
	  float: left;
	  border: none;
	  outline: none;
	  font-size: var(--jp-ui-font-size1);
	  color: var(--jp-ui-font-color0);
	  line-height: var(--jp-private-commandpalette-search-height);
	}
	
	.lm-CommandPalette-input::-webkit-input-placeholder,
	.lm-CommandPalette-input::-moz-placeholder,
	.lm-CommandPalette-input:-ms-input-placeholder {
	  color: var(--jp-ui-font-color3);
	  font-size: var(--jp-ui-font-size1);
	}
	
	/*-----------------------------------------------------------------------------
	| Results
	|----------------------------------------------------------------------------*/
	
	.lm-CommandPalette-header:first-child {
	  margin-top: 0px;
	}
	
	.lm-CommandPalette-header {
	  border-bottom: solid var(--jp-border-width) var(--jp-border-color2);
	  color: var(--jp-ui-font-color1);
	  cursor: pointer;
	  display: flex;
	  font-size: var(--jp-ui-font-size0);
	  font-weight: 600;
	  letter-spacing: 1px;
	  margin-top: 8px;
	  padding: 8px 0 8px 12px;
	  text-transform: uppercase;
	}
	
	.lm-CommandPalette-header.lm-mod-active {
	  background: var(--jp-layout-color2);
	}
	
	.lm-CommandPalette-header > mark {
	  background-color: transparent;
	  font-weight: bold;
	  color: var(--jp-ui-font-color1);
	}
	
	.lm-CommandPalette-item {
	  padding: 4px 12px 4px 4px;
	  color: var(--jp-ui-font-color1);
	  font-size: var(--jp-ui-font-size1);
	  font-weight: 400;
	  display: flex;
	}
	
	.lm-CommandPalette-item.lm-mod-disabled {
	  color: var(--jp-ui-font-color3);
	}
	
	.lm-CommandPalette-item.lm-mod-active {
	  background: var(--jp-layout-color3);
	}
	
	.lm-CommandPalette-item.lm-mod-active:hover:not(.lm-mod-disabled) {
	  background: var(--jp-layout-color4);
	}
	
	.lm-CommandPalette-item:hover:not(.lm-mod-active):not(.lm-mod-disabled) {
	  background: var(--jp-layout-color2);
	}
	
	.lm-CommandPalette-itemContent {
	  overflow: hidden;
	}
	
	.lm-CommandPalette-itemLabel > mark {
	  color: var(--jp-ui-font-color0);
	  background-color: transparent;
	  font-weight: bold;
	}
	
	.lm-CommandPalette-item.lm-mod-disabled mark {
	  color: var(--jp-ui-font-color3);
	}
	
	.lm-CommandPalette-item .lm-CommandPalette-itemIcon {
	  margin: 0 4px 0 0;
	  position: relative;
	  width: 16px;
	  top: 2px;
	  flex: 0 0 auto;
	}
	
	.lm-CommandPalette-item.lm-mod-disabled .lm-CommandPalette-itemIcon {
	  opacity: 0.4;
	}
	
	.lm-CommandPalette-item .lm-CommandPalette-itemShortcut {
	  flex: 0 0 auto;
	}
	
	.lm-CommandPalette-itemCaption {
	  display: none;
	}
	
	.lm-CommandPalette-content {
	  background-color: var(--jp-layout-color1);
	}
	
	.lm-CommandPalette-content:empty:after {
	  content: 'No results';
	  margin: auto;
	  margin-top: 20px;
	  width: 100px;
	  display: block;
	  font-size: var(--jp-ui-font-size2);
	  font-family: var(--jp-ui-font-family);
	  font-weight: lighter;
	}
	
	.lm-CommandPalette-emptyMessage {
	  text-align: center;
	  margin-top: 24px;
	  line-height: 1.32;
	  padding: 0px 8px;
	  color: var(--jp-content-font-color3);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2017, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-Dialog {
	  position: absolute;
	  z-index: 10000;
	  display: flex;
	  flex-direction: column;
	  align-items: center;
	  justify-content: center;
	  top: 0px;
	  left: 0px;
	  margin: 0;
	  padding: 0;
	  width: 100%;
	  height: 100%;
	  background: var(--jp-dialog-background);
	}
	
	.jp-Dialog-content {
	  display: flex;
	  flex-direction: column;
	  margin-left: auto;
	  margin-right: auto;
	  background: var(--jp-layout-color1);
	  padding: 24px;
	  padding-bottom: 12px;
	  min-width: 300px;
	  min-height: 150px;
	  max-width: 1000px;
	  max-height: 500px;
	  box-sizing: border-box;
	  box-shadow: var(--jp-elevation-z20);
	  word-wrap: break-word;
	  border-radius: var(--jp-border-radius);
	  /* This is needed so that all font sizing of children done in ems is
	   * relative to this base size */
	  font-size: var(--jp-ui-font-size1);
	  color: var(--jp-ui-font-color1);
	  resize: both;
	}
	
	.jp-Dialog-button {
	  overflow: visible;
	}
	
	button.jp-Dialog-button:focus {
	  outline: 1px solid var(--jp-brand-color1);
	  outline-offset: 4px;
	  -moz-outline-radius: 0px;
	}
	
	button.jp-Dialog-button:focus::-moz-focus-inner {
	  border: 0;
	}
	
	button.jp-Dialog-close-button {
	  padding: 0;
	  height: 100%;
	  min-width: unset;
	  min-height: unset;
	}
	
	.jp-Dialog-header {
	  display: flex;
	  justify-content: space-between;
	  flex: 0 0 auto;
	  padding-bottom: 12px;
	  font-size: var(--jp-ui-font-size3);
	  font-weight: 400;
	  color: var(--jp-ui-font-color0);
	}
	
	.jp-Dialog-body {
	  display: flex;
	  flex-direction: column;
	  flex: 1 1 auto;
	  font-size: var(--jp-ui-font-size1);
	  background: var(--jp-layout-color1);
	  overflow: auto;
	}
	
	.jp-Dialog-footer {
	  display: flex;
	  flex-direction: row;
	  justify-content: flex-end;
	  flex: 0 0 auto;
	  margin-left: -12px;
	  margin-right: -12px;
	  padding: 12px;
	}
	
	.jp-Dialog-title {
	  overflow: hidden;
	  white-space: nowrap;
	  text-overflow: ellipsis;
	}
	
	.jp-Dialog-body > .jp-select-wrapper {
	  width: 100%;
	}
	
	.jp-Dialog-body > button {
	  padding: 0px 16px;
	}
	
	.jp-Dialog-body > label {
	  line-height: 1.4;
	  color: var(--jp-ui-font-color0);
	}
	
	.jp-Dialog-button.jp-mod-styled:not(:last-child) {
	  margin-right: 12px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2016, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-HoverBox {
	  position: fixed;
	}
	
	.jp-HoverBox.jp-mod-outofview {
	  display: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-IFrame {
	  width: 100%;
	  height: 100%;
	}
	
	.jp-IFrame > iframe {
	  border: none;
	}
	
	body.lm-mod-override-cursor .jp-IFrame {
	  position: relative;
	}
	
	body.lm-mod-override-cursor .jp-IFrame:before {
	  content: '';
	  position: absolute;
	  top: 0;
	  left: 0;
	  right: 0;
	  bottom: 0;
	  background: transparent;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2016, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-MainAreaWidget > :focus {
	  outline: none;
	}
	
	/**
	 * google-material-color v1.2.6
	 * https://github.com/danlevan/google-material-color
	 */
	:root {
	  --md-red-50: #ffebee;
	  --md-red-100: #ffcdd2;
	  --md-red-200: #ef9a9a;
	  --md-red-300: #e57373;
	  --md-red-400: #ef5350;
	  --md-red-500: #f44336;
	  --md-red-600: #e53935;
	  --md-red-700: #d32f2f;
	  --md-red-800: #c62828;
	  --md-red-900: #b71c1c;
	  --md-red-A100: #ff8a80;
	  --md-red-A200: #ff5252;
	  --md-red-A400: #ff1744;
	  --md-red-A700: #d50000;
	
	  --md-pink-50: #fce4ec;
	  --md-pink-100: #f8bbd0;
	  --md-pink-200: #f48fb1;
	  --md-pink-300: #f06292;
	  --md-pink-400: #ec407a;
	  --md-pink-500: #e91e63;
	  --md-pink-600: #d81b60;
	  --md-pink-700: #c2185b;
	  --md-pink-800: #ad1457;
	  --md-pink-900: #880e4f;
	  --md-pink-A100: #ff80ab;
	  --md-pink-A200: #ff4081;
	  --md-pink-A400: #f50057;
	  --md-pink-A700: #c51162;
	
	  --md-purple-50: #f3e5f5;
	  --md-purple-100: #e1bee7;
	  --md-purple-200: #ce93d8;
	  --md-purple-300: #ba68c8;
	  --md-purple-400: #ab47bc;
	  --md-purple-500: #9c27b0;
	  --md-purple-600: #8e24aa;
	  --md-purple-700: #7b1fa2;
	  --md-purple-800: #6a1b9a;
	  --md-purple-900: #4a148c;
	  --md-purple-A100: #ea80fc;
	  --md-purple-A200: #e040fb;
	  --md-purple-A400: #d500f9;
	  --md-purple-A700: #aa00ff;
	
	  --md-deep-purple-50: #ede7f6;
	  --md-deep-purple-100: #d1c4e9;
	  --md-deep-purple-200: #b39ddb;
	  --md-deep-purple-300: #9575cd;
	  --md-deep-purple-400: #7e57c2;
	  --md-deep-purple-500: #673ab7;
	  --md-deep-purple-600: #5e35b1;
	  --md-deep-purple-700: #512da8;
	  --md-deep-purple-800: #4527a0;
	  --md-deep-purple-900: #311b92;
	  --md-deep-purple-A100: #b388ff;
	  --md-deep-purple-A200: #7c4dff;
	  --md-deep-purple-A400: #651fff;
	  --md-deep-purple-A700: #6200ea;
	
	  --md-indigo-50: #e8eaf6;
	  --md-indigo-100: #c5cae9;
	  --md-indigo-200: #9fa8da;
	  --md-indigo-300: #7986cb;
	  --md-indigo-400: #5c6bc0;
	  --md-indigo-500: #3f51b5;
	  --md-indigo-600: #3949ab;
	  --md-indigo-700: #303f9f;
	  --md-indigo-800: #283593;
	  --md-indigo-900: #1a237e;
	  --md-indigo-A100: #8c9eff;
	  --md-indigo-A200: #536dfe;
	  --md-indigo-A400: #3d5afe;
	  --md-indigo-A700: #304ffe;
	
	  --md-blue-50: #e3f2fd;
	  --md-blue-100: #bbdefb;
	  --md-blue-200: #90caf9;
	  --md-blue-300: #64b5f6;
	  --md-blue-400: #42a5f5;
	  --md-blue-500: #2196f3;
	  --md-blue-600: #1e88e5;
	  --md-blue-700: #1976d2;
	  --md-blue-800: #1565c0;
	  --md-blue-900: #0d47a1;
	  --md-blue-A100: #82b1ff;
	  --md-blue-A200: #448aff;
	  --md-blue-A400: #2979ff;
	  --md-blue-A700: #2962ff;
	
	  --md-light-blue-50: #e1f5fe;
	  --md-light-blue-100: #b3e5fc;
	  --md-light-blue-200: #81d4fa;
	  --md-light-blue-300: #4fc3f7;
	  --md-light-blue-400: #29b6f6;
	  --md-light-blue-500: #03a9f4;
	  --md-light-blue-600: #039be5;
	  --md-light-blue-700: #0288d1;
	  --md-light-blue-800: #0277bd;
	  --md-light-blue-900: #01579b;
	  --md-light-blue-A100: #80d8ff;
	  --md-light-blue-A200: #40c4ff;
	  --md-light-blue-A400: #00b0ff;
	  --md-light-blue-A700: #0091ea;
	
	  --md-cyan-50: #e0f7fa;
	  --md-cyan-100: #b2ebf2;
	  --md-cyan-200: #80deea;
	  --md-cyan-300: #4dd0e1;
	  --md-cyan-400: #26c6da;
	  --md-cyan-500: #00bcd4;
	  --md-cyan-600: #00acc1;
	  --md-cyan-700: #0097a7;
	  --md-cyan-800: #00838f;
	  --md-cyan-900: #006064;
	  --md-cyan-A100: #84ffff;
	  --md-cyan-A200: #18ffff;
	  --md-cyan-A400: #00e5ff;
	  --md-cyan-A700: #00b8d4;
	
	  --md-teal-50: #e0f2f1;
	  --md-teal-100: #b2dfdb;
	  --md-teal-200: #80cbc4;
	  --md-teal-300: #4db6ac;
	  --md-teal-400: #26a69a;
	  --md-teal-500: #009688;
	  --md-teal-600: #00897b;
	  --md-teal-700: #00796b;
	  --md-teal-800: #00695c;
	  --md-teal-900: #004d40;
	  --md-teal-A100: #a7ffeb;
	  --md-teal-A200: #64ffda;
	  --md-teal-A400: #1de9b6;
	  --md-teal-A700: #00bfa5;
	
	  --md-green-50: #e8f5e9;
	  --md-green-100: #c8e6c9;
	  --md-green-200: #a5d6a7;
	  --md-green-300: #81c784;
	  --md-green-400: #66bb6a;
	  --md-green-500: #4caf50;
	  --md-green-600: #43a047;
	  --md-green-700: #388e3c;
	  --md-green-800: #2e7d32;
	  --md-green-900: #1b5e20;
	  --md-green-A100: #b9f6ca;
	  --md-green-A200: #69f0ae;
	  --md-green-A400: #00e676;
	  --md-green-A700: #00c853;
	
	  --md-light-green-50: #f1f8e9;
	  --md-light-green-100: #dcedc8;
	  --md-light-green-200: #c5e1a5;
	  --md-light-green-300: #aed581;
	  --md-light-green-400: #9ccc65;
	  --md-light-green-500: #8bc34a;
	  --md-light-green-600: #7cb342;
	  --md-light-green-700: #689f38;
	  --md-light-green-800: #558b2f;
	  --md-light-green-900: #33691e;
	  --md-light-green-A100: #ccff90;
	  --md-light-green-A200: #b2ff59;
	  --md-light-green-A400: #76ff03;
	  --md-light-green-A700: #64dd17;
	
	  --md-lime-50: #f9fbe7;
	  --md-lime-100: #f0f4c3;
	  --md-lime-200: #e6ee9c;
	  --md-lime-300: #dce775;
	  --md-lime-400: #d4e157;
	  --md-lime-500: #cddc39;
	  --md-lime-600: #c0ca33;
	  --md-lime-700: #afb42b;
	  --md-lime-800: #9e9d24;
	  --md-lime-900: #827717;
	  --md-lime-A100: #f4ff81;
	  --md-lime-A200: #eeff41;
	  --md-lime-A400: #c6ff00;
	  --md-lime-A700: #aeea00;
	
	  --md-yellow-50: #fffde7;
	  --md-yellow-100: #fff9c4;
	  --md-yellow-200: #fff59d;
	  --md-yellow-300: #fff176;
	  --md-yellow-400: #ffee58;
	  --md-yellow-500: #ffeb3b;
	  --md-yellow-600: #fdd835;
	  --md-yellow-700: #fbc02d;
	  --md-yellow-800: #f9a825;
	  --md-yellow-900: #f57f17;
	  --md-yellow-A100: #ffff8d;
	  --md-yellow-A200: #ffff00;
	  --md-yellow-A400: #ffea00;
	  --md-yellow-A700: #ffd600;
	
	  --md-amber-50: #fff8e1;
	  --md-amber-100: #ffecb3;
	  --md-amber-200: #ffe082;
	  --md-amber-300: #ffd54f;
	  --md-amber-400: #ffca28;
	  --md-amber-500: #ffc107;
	  --md-amber-600: #ffb300;
	  --md-amber-700: #ffa000;
	  --md-amber-800: #ff8f00;
	  --md-amber-900: #ff6f00;
	  --md-amber-A100: #ffe57f;
	  --md-amber-A200: #ffd740;
	  --md-amber-A400: #ffc400;
	  --md-amber-A700: #ffab00;
	
	  --md-orange-50: #fff3e0;
	  --md-orange-100: #ffe0b2;
	  --md-orange-200: #ffcc80;
	  --md-orange-300: #ffb74d;
	  --md-orange-400: #ffa726;
	  --md-orange-500: #ff9800;
	  --md-orange-600: #fb8c00;
	  --md-orange-700: #f57c00;
	  --md-orange-800: #ef6c00;
	  --md-orange-900: #e65100;
	  --md-orange-A100: #ffd180;
	  --md-orange-A200: #ffab40;
	  --md-orange-A400: #ff9100;
	  --md-orange-A700: #ff6d00;
	
	  --md-deep-orange-50: #fbe9e7;
	  --md-deep-orange-100: #ffccbc;
	  --md-deep-orange-200: #ffab91;
	  --md-deep-orange-300: #ff8a65;
	  --md-deep-orange-400: #ff7043;
	  --md-deep-orange-500: #ff5722;
	  --md-deep-orange-600: #f4511e;
	  --md-deep-orange-700: #e64a19;
	  --md-deep-orange-800: #d84315;
	  --md-deep-orange-900: #bf360c;
	  --md-deep-orange-A100: #ff9e80;
	  --md-deep-orange-A200: #ff6e40;
	  --md-deep-orange-A400: #ff3d00;
	  --md-deep-orange-A700: #dd2c00;
	
	  --md-brown-50: #efebe9;
	  --md-brown-100: #d7ccc8;
	  --md-brown-200: #bcaaa4;
	  --md-brown-300: #a1887f;
	  --md-brown-400: #8d6e63;
	  --md-brown-500: #795548;
	  --md-brown-600: #6d4c41;
	  --md-brown-700: #5d4037;
	  --md-brown-800: #4e342e;
	  --md-brown-900: #3e2723;
	
	  --md-grey-50: #fafafa;
	  --md-grey-100: #f5f5f5;
	  --md-grey-200: #eeeeee;
	  --md-grey-300: #e0e0e0;
	  --md-grey-400: #bdbdbd;
	  --md-grey-500: #9e9e9e;
	  --md-grey-600: #757575;
	  --md-grey-700: #616161;
	  --md-grey-800: #424242;
	  --md-grey-900: #212121;
	
	  --md-blue-grey-50: #eceff1;
	  --md-blue-grey-100: #cfd8dc;
	  --md-blue-grey-200: #b0bec5;
	  --md-blue-grey-300: #90a4ae;
	  --md-blue-grey-400: #78909c;
	  --md-blue-grey-500: #607d8b;
	  --md-blue-grey-600: #546e7a;
	  --md-blue-grey-700: #455a64;
	  --md-blue-grey-800: #37474f;
	  --md-blue-grey-900: #263238;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2017, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-Spinner {
	  position: absolute;
	  display: flex;
	  justify-content: center;
	  align-items: center;
	  z-index: 10;
	  left: 0;
	  top: 0;
	  width: 100%;
	  height: 100%;
	  background: var(--jp-layout-color0);
	  outline: none;
	}
	
	.jp-SpinnerContent {
	  font-size: 10px;
	  margin: 50px auto;
	  text-indent: -9999em;
	  width: 3em;
	  height: 3em;
	  border-radius: 50%;
	  background: var(--jp-brand-color3);
	  background: linear-gradient(
		to right,
		#f37626 10%,
		rgba(255, 255, 255, 0) 42%
	  );
	  position: relative;
	  animation: load3 1s infinite linear, fadeIn 1s;
	}
	
	.jp-SpinnerContent:before {
	  width: 50%;
	  height: 50%;
	  background: #f37626;
	  border-radius: 100% 0 0 0;
	  position: absolute;
	  top: 0;
	  left: 0;
	  content: '';
	}
	
	.jp-SpinnerContent:after {
	  background: var(--jp-layout-color0);
	  width: 75%;
	  height: 75%;
	  border-radius: 50%;
	  content: '';
	  margin: auto;
	  position: absolute;
	  top: 0;
	  left: 0;
	  bottom: 0;
	  right: 0;
	}
	
	@keyframes fadeIn {
	  0% {
		opacity: 0;
	  }
	  100% {
		opacity: 1;
	  }
	}
	
	@keyframes load3 {
	  0% {
		transform: rotate(0deg);
	  }
	  100% {
		transform: rotate(360deg);
	  }
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2017, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	button.jp-mod-styled {
	  font-size: var(--jp-ui-font-size1);
	  color: var(--jp-ui-font-color0);
	  border: none;
	  box-sizing: border-box;
	  text-align: center;
	  line-height: 32px;
	  height: 32px;
	  padding: 0px 12px;
	  letter-spacing: 0.8px;
	  outline: none;
	  appearance: none;
	  -webkit-appearance: none;
	  -moz-appearance: none;
	}
	
	input.jp-mod-styled {
	  background: var(--jp-input-background);
	  height: 28px;
	  box-sizing: border-box;
	  border: var(--jp-border-width) solid var(--jp-border-color1);
	  padding-left: 7px;
	  padding-right: 7px;
	  font-size: var(--jp-ui-font-size2);
	  color: var(--jp-ui-font-color0);
	  outline: none;
	  appearance: none;
	  -webkit-appearance: none;
	  -moz-appearance: none;
	}
	
	input.jp-mod-styled:focus {
	  border: var(--jp-border-width) solid var(--md-blue-500);
	  box-shadow: inset 0 0 4px var(--md-blue-300);
	}
	
	.jp-select-wrapper {
	  display: flex;
	  position: relative;
	  flex-direction: column;
	  padding: 1px;
	  background-color: var(--jp-layout-color1);
	  height: 28px;
	  box-sizing: border-box;
	  margin-bottom: 12px;
	}
	
	.jp-select-wrapper.jp-mod-focused select.jp-mod-styled {
	  border: var(--jp-border-width) solid var(--jp-input-active-border-color);
	  box-shadow: var(--jp-input-box-shadow);
	  background-color: var(--jp-input-active-background);
	}
	
	select.jp-mod-styled:hover {
	  background-color: var(--jp-layout-color1);
	  cursor: pointer;
	  color: var(--jp-ui-font-color0);
	  background-color: var(--jp-input-hover-background);
	  box-shadow: inset 0 0px 1px rgba(0, 0, 0, 0.5);
	}
	
	select.jp-mod-styled {
	  flex: 1 1 auto;
	  height: 32px;
	  width: 100%;
	  font-size: var(--jp-ui-font-size2);
	  background: var(--jp-input-background);
	  color: var(--jp-ui-font-color0);
	  padding: 0 25px 0 8px;
	  border: var(--jp-border-width) solid var(--jp-input-border-color);
	  border-radius: 0px;
	  outline: none;
	  appearance: none;
	  -webkit-appearance: none;
	  -moz-appearance: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2016, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	:root {
	  --jp-private-toolbar-height: calc(
		28px + var(--jp-border-width)
	  ); /* leave 28px for content */
	}
	
	.jp-Toolbar {
	  color: var(--jp-ui-font-color1);
	  flex: 0 0 auto;
	  display: flex;
	  flex-direction: row;
	  border-bottom: var(--jp-border-width) solid var(--jp-toolbar-border-color);
	  box-shadow: var(--jp-toolbar-box-shadow);
	  background: var(--jp-toolbar-background);
	  min-height: var(--jp-toolbar-micro-height);
	  padding: 2px;
	  z-index: 1;
	  overflow-x: hidden;
	}
	
	.jp-Toolbar:hover {
	  overflow-x: auto;
	}
	
	/* Toolbar items */
	
	.jp-Toolbar > .jp-Toolbar-item.jp-Toolbar-spacer {
	  flex-grow: 1;
	  flex-shrink: 1;
	}
	
	.jp-Toolbar-item.jp-Toolbar-kernelStatus {
	  display: inline-block;
	  width: 32px;
	  background-repeat: no-repeat;
	  background-position: center;
	  background-size: 16px;
	}
	
	.jp-Toolbar > .jp-Toolbar-item {
	  flex: 0 0 auto;
	  display: flex;
	  padding-left: 1px;
	  padding-right: 1px;
	  font-size: var(--jp-ui-font-size1);
	  line-height: var(--jp-private-toolbar-height);
	  height: 100%;
	}
	
	/* Toolbar buttons */
	
	/* This is the div we use to wrap the react component into a Widget */
	div.jp-ToolbarButton {
	  color: transparent;
	  border: none;
	  box-sizing: border-box;
	  outline: none;
	  appearance: none;
	  -webkit-appearance: none;
	  -moz-appearance: none;
	  padding: 0px;
	  margin: 0px;
	}
	
	button.jp-ToolbarButtonComponent {
	  background: var(--jp-layout-color1);
	  border: none;
	  box-sizing: border-box;
	  outline: none;
	  appearance: none;
	  -webkit-appearance: none;
	  -moz-appearance: none;
	  padding: 0px 6px;
	  margin: 0px;
	  height: 24px;
	  border-radius: var(--jp-border-radius);
	  display: flex;
	  align-items: center;
	  text-align: center;
	  font-size: 14px;
	  min-width: unset;
	  min-height: unset;
	}
	
	button.jp-ToolbarButtonComponent:disabled {
	  opacity: 0.4;
	}
	
	button.jp-ToolbarButtonComponent span {
	  padding: 0px;
	  flex: 0 0 auto;
	}
	
	button.jp-ToolbarButtonComponent .jp-ToolbarButtonComponent-label {
	  font-size: var(--jp-ui-font-size1);
	  line-height: 100%;
	  padding-left: 2px;
	  color: var(--jp-ui-font-color1);
	}
	
	#jp-main-dock-panel[data-mode='single-document']
	  .jp-MainAreaWidget
	  > .jp-Toolbar.jp-Toolbar-micro {
	  padding: 0;
	  min-height: 0;
	}
	
	#jp-main-dock-panel[data-mode='single-document']
	  .jp-MainAreaWidget
	  > .jp-Toolbar {
	  border: none;
	  box-shadow: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2017, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Copyright (c) 2014-2017, PhosphorJS Contributors
	|
	| Distributed under the terms of the BSD 3-Clause License.
	|
	| The full license is in the file LICENSE, distributed with this software.
	|----------------------------------------------------------------------------*/
	
	
	/* <DEPRECATED> */ body.p-mod-override-cursor *, /* </DEPRECATED> */
	body.lm-mod-override-cursor * {
	  cursor: inherit !important;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) 2014-2016, Jupyter Development Team.
	|
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-JSONEditor {
	  display: flex;
	  flex-direction: column;
	  width: 100%;
	}
	
	.jp-JSONEditor-host {
	  flex: 1 1 auto;
	  border: var(--jp-border-width) solid var(--jp-input-border-color);
	  border-radius: 0px;
	  background: var(--jp-layout-color0);
	  min-height: 50px;
	  padding: 1px;
	}
	
	.jp-JSONEditor.jp-mod-error .jp-JSONEditor-host {
	  border-color: red;
	  outline-color: red;
	}
	
	.jp-JSONEditor-header {
	  display: flex;
	  flex: 1 0 auto;
	  padding: 0 0 0 12px;
	}
	
	.jp-JSONEditor-header label {
	  flex: 0 0 auto;
	}
	
	.jp-JSONEditor-commitButton {
	  height: 16px;
	  width: 16px;
	  background-size: 18px;
	  background-repeat: no-repeat;
	  background-position: center;
	}
	
	.jp-JSONEditor-host.jp-mod-focused {
	  background-color: var(--jp-input-active-background);
	  border: 1px solid var(--jp-input-active-border-color);
	  box-shadow: var(--jp-input-box-shadow);
	}
	
	.jp-Editor.jp-mod-dropTarget {
	  border: var(--jp-border-width) solid var(--jp-input-active-border-color);
	  box-shadow: var(--jp-input-box-shadow);
	}
	
	/* BASICS */
	
	.CodeMirror {
	  /* Set height, width, borders, and global font properties here */
	  font-family: monospace;
	  height: 300px;
	  color: black;
	  direction: ltr;
	}
	
	/* PADDING */
	
	.CodeMirror-lines {
	  padding: 4px 0; /* Vertical padding around content */
	}
	.CodeMirror pre.CodeMirror-line,
	.CodeMirror pre.CodeMirror-line-like {
	  padding: 0 4px; /* Horizontal padding of content */
	}
	
	.CodeMirror-scrollbar-filler, .CodeMirror-gutter-filler {
	  background-color: white; /* The little square between H and V scrollbars */
	}
	
	/* GUTTER */
	
	.CodeMirror-gutters {
	  border-right: 1px solid #ddd;
	  background-color: #f7f7f7;
	  white-space: nowrap;
	}
	.CodeMirror-linenumbers {}
	.CodeMirror-linenumber {
	  padding: 0 3px 0 5px;
	  min-width: 20px;
	  text-align: right;
	  color: #999;
	  white-space: nowrap;
	}
	
	.CodeMirror-guttermarker { color: black; }
	.CodeMirror-guttermarker-subtle { color: #999; }
	
	/* CURSOR */
	
	.CodeMirror-cursor {
	  border-left: 1px solid black;
	  border-right: none;
	  width: 0;
	}
	/* Shown when moving in bi-directional text */
	.CodeMirror div.CodeMirror-secondarycursor {
	  border-left: 1px solid silver;
	}
	.cm-fat-cursor .CodeMirror-cursor {
	  width: auto;
	  border: 0 !important;
	  background: #7e7;
	}
	.cm-fat-cursor div.CodeMirror-cursors {
	  z-index: 1;
	}
	.cm-fat-cursor-mark {
	  background-color: rgba(20, 255, 20, 0.5);
	  -webkit-animation: blink 1.06s steps(1) infinite;
	  -moz-animation: blink 1.06s steps(1) infinite;
	  animation: blink 1.06s steps(1) infinite;
	}
	.cm-animate-fat-cursor {
	  width: auto;
	  border: 0;
	  -webkit-animation: blink 1.06s steps(1) infinite;
	  -moz-animation: blink 1.06s steps(1) infinite;
	  animation: blink 1.06s steps(1) infinite;
	  background-color: #7e7;
	}
	@-moz-keyframes blink {
	  0% {}
	  50% { background-color: transparent; }
	  100% {}
	}
	@-webkit-keyframes blink {
	  0% {}
	  50% { background-color: transparent; }
	  100% {}
	}
	@keyframes blink {
	  0% {}
	  50% { background-color: transparent; }
	  100% {}
	}
	
	/* Can style cursor different in overwrite (non-insert) mode */
	.CodeMirror-overwrite .CodeMirror-cursor {}
	
	.cm-tab { display: inline-block; text-decoration: inherit; }
	
	.CodeMirror-rulers {
	  position: absolute;
	  left: 0; right: 0; top: -50px; bottom: 0;
	  overflow: hidden;
	}
	.CodeMirror-ruler {
	  border-left: 1px solid #ccc;
	  top: 0; bottom: 0;
	  position: absolute;
	}
	
	/* DEFAULT THEME */
	
	.cm-s-default .cm-header {color: blue;}
	.cm-s-default .cm-quote {color: #090;}
	.cm-negative {color: #d44;}
	.cm-positive {color: #292;}
	.cm-header, .cm-strong {font-weight: bold;}
	.cm-em {font-style: italic;}
	.cm-link {text-decoration: underline;}
	.cm-strikethrough {text-decoration: line-through;}
	
	.cm-s-default .cm-keyword {color: #708;}
	.cm-s-default .cm-atom {color: #219;}
	.cm-s-default .cm-number {color: #164;}
	.cm-s-default .cm-def {color: #00f;}
	.cm-s-default .cm-variable,
	.cm-s-default .cm-punctuation,
	.cm-s-default .cm-property,
	.cm-s-default .cm-operator {}
	.cm-s-default .cm-variable-2 {color: #05a;}
	.cm-s-default .cm-variable-3, .cm-s-default .cm-type {color: #085;}
	.cm-s-default .cm-comment {color: #a50;}
	.cm-s-default .cm-string {color: #a11;}
	.cm-s-default .cm-string-2 {color: #f50;}
	.cm-s-default .cm-meta {color: #555;}
	.cm-s-default .cm-qualifier {color: #555;}
	.cm-s-default .cm-builtin {color: #30a;}
	.cm-s-default .cm-bracket {color: #997;}
	.cm-s-default .cm-tag {color: #170;}
	.cm-s-default .cm-attribute {color: #00c;}
	.cm-s-default .cm-hr {color: #999;}
	.cm-s-default .cm-link {color: #00c;}
	
	.cm-s-default .cm-error {color: #f00;}
	.cm-invalidchar {color: #f00;}
	
	.CodeMirror-composing { border-bottom: 2px solid; }
	
	/* Default styles for common addons */
	
	div.CodeMirror span.CodeMirror-matchingbracket {color: #0b0;}
	div.CodeMirror span.CodeMirror-nonmatchingbracket {color: #a22;}
	.CodeMirror-matchingtag { background: rgba(255, 150, 0, .3); }
	.CodeMirror-activeline-background {background: #e8f2ff;}
	
	/* STOP */
	
	/* The rest of this file contains styles related to the mechanics of
	   the editor. You probably shouldn't touch them. */
	
	.CodeMirror {
	  position: relative;
	  overflow: hidden;
	  background: white;
	}
	
	.CodeMirror-scroll {
	  overflow: scroll !important; /* Things will break if this is overridden */
	  /* 50px is the magic margin used to hide the element's real scrollbars */
	  /* See overflow: hidden in .CodeMirror */
	  margin-bottom: -50px; margin-right: -50px;
	  padding-bottom: 50px;
	  height: 100%;
	  outline: none; /* Prevent dragging from highlighting the element */
	  position: relative;
	}
	.CodeMirror-sizer {
	  position: relative;
	  border-right: 50px solid transparent;
	}
	
	/* The fake, visible scrollbars. Used to force redraw during scrolling
	   before actual scrolling happens, thus preventing shaking and
	   flickering artifacts. */
	.CodeMirror-vscrollbar, .CodeMirror-hscrollbar, .CodeMirror-scrollbar-filler, .CodeMirror-gutter-filler {
	  position: absolute;
	  z-index: 6;
	  display: none;
	}
	.CodeMirror-vscrollbar {
	  right: 0; top: 0;
	  overflow-x: hidden;
	  overflow-y: scroll;
	}
	.CodeMirror-hscrollbar {
	  bottom: 0; left: 0;
	  overflow-y: hidden;
	  overflow-x: scroll;
	}
	.CodeMirror-scrollbar-filler {
	  right: 0; bottom: 0;
	}
	.CodeMirror-gutter-filler {
	  left: 0; bottom: 0;
	}
	
	.CodeMirror-gutters {
	  position: absolute; left: 0; top: 0;
	  min-height: 100%;
	  z-index: 3;
	}
	.CodeMirror-gutter {
	  white-space: normal;
	  height: 100%;
	  display: inline-block;
	  vertical-align: top;
	  margin-bottom: -50px;
	}
	.CodeMirror-gutter-wrapper {
	  position: absolute;
	  z-index: 4;
	  background: none !important;
	  border: none !important;
	}
	.CodeMirror-gutter-background {
	  position: absolute;
	  top: 0; bottom: 0;
	  z-index: 4;
	}
	.CodeMirror-gutter-elt {
	  position: absolute;
	  cursor: default;
	  z-index: 4;
	}
	.CodeMirror-gutter-wrapper ::selection { background-color: transparent }
	.CodeMirror-gutter-wrapper ::-moz-selection { background-color: transparent }
	
	.CodeMirror-lines {
	  cursor: text;
	  min-height: 1px; /* prevents collapsing before first draw */
	}
	.CodeMirror pre.CodeMirror-line,
	.CodeMirror pre.CodeMirror-line-like {
	  /* Reset some styles that the rest of the page might have set */
	  -moz-border-radius: 0; -webkit-border-radius: 0; border-radius: 0;
	  border-width: 0;
	  background: transparent;
	  font-family: inherit;
	  font-size: inherit;
	  margin: 0;
	  white-space: pre;
	  word-wrap: normal;
	  line-height: inherit;
	  color: inherit;
	  z-index: 2;
	  position: relative;
	  overflow: visible;
	  -webkit-tap-highlight-color: transparent;
	  -webkit-font-variant-ligatures: contextual;
	  font-variant-ligatures: contextual;
	}
	.CodeMirror-wrap pre.CodeMirror-line,
	.CodeMirror-wrap pre.CodeMirror-line-like {
	  word-wrap: break-word;
	  white-space: pre-wrap;
	  word-break: normal;
	}
	
	.CodeMirror-linebackground {
	  position: absolute;
	  left: 0; right: 0; top: 0; bottom: 0;
	  z-index: 0;
	}
	
	.CodeMirror-linewidget {
	  position: relative;
	  z-index: 2;
	  padding: 0.1px; /* Force widget margins to stay inside of the container */
	}
	
	.CodeMirror-widget {}
	
	.CodeMirror-rtl pre { direction: rtl; }
	
	.CodeMirror-code {
	  outline: none;
	}
	
	/* Force content-box sizing for the elements where we expect it */
	.CodeMirror-scroll,
	.CodeMirror-sizer,
	.CodeMirror-gutter,
	.CodeMirror-gutters,
	.CodeMirror-linenumber {
	  -moz-box-sizing: content-box;
	  box-sizing: content-box;
	}
	
	.CodeMirror-measure {
	  position: absolute;
	  width: 100%;
	  height: 0;
	  overflow: hidden;
	  visibility: hidden;
	}
	
	.CodeMirror-cursor {
	  position: absolute;
	  pointer-events: none;
	}
	.CodeMirror-measure pre { position: static; }
	
	div.CodeMirror-cursors {
	  visibility: hidden;
	  position: relative;
	  z-index: 3;
	}
	div.CodeMirror-dragcursors {
	  visibility: visible;
	}
	
	.CodeMirror-focused div.CodeMirror-cursors {
	  visibility: visible;
	}
	
	.CodeMirror-selected { background: #d9d9d9; }
	.CodeMirror-focused .CodeMirror-selected { background: #d7d4f0; }
	.CodeMirror-crosshair { cursor: crosshair; }
	.CodeMirror-line::selection, .CodeMirror-line > span::selection, .CodeMirror-line > span > span::selection { background: #d7d4f0; }
	.CodeMirror-line::-moz-selection, .CodeMirror-line > span::-moz-selection, .CodeMirror-line > span > span::-moz-selection { background: #d7d4f0; }
	
	.cm-searching {
	  background-color: #ffa;
	  background-color: rgba(255, 255, 0, .4);
	}
	
	/* Used to force a border model for a node */
	.cm-force-border { padding-right: .1px; }
	
	@media print {
	  /* Hide the cursor when printing */
	  .CodeMirror div.CodeMirror-cursors {
		visibility: hidden;
	  }
	}
	
	/* See issue #2901 */
	.cm-tab-wrap-hack:after { content: ''; }
	
	/* Help users use markselection to safely style text background */
	span.CodeMirror-selectedtext { background: none; }
	
	.CodeMirror-dialog {
	  position: absolute;
	  left: 0; right: 0;
	  background: inherit;
	  z-index: 15;
	  padding: .1em .8em;
	  overflow: hidden;
	  color: inherit;
	}
	
	.CodeMirror-dialog-top {
	  border-bottom: 1px solid #eee;
	  top: 0;
	}
	
	.CodeMirror-dialog-bottom {
	  border-top: 1px solid #eee;
	  bottom: 0;
	}
	
	.CodeMirror-dialog input {
	  border: none;
	  outline: none;
	  background: transparent;
	  width: 20em;
	  color: inherit;
	  font-family: monospace;
	}
	
	.CodeMirror-dialog button {
	  font-size: 70%;
	}
	
	.CodeMirror-foldmarker {
	  color: blue;
	  text-shadow: #b9f 1px 1px 2px, #b9f -1px -1px 2px, #b9f 1px -1px 2px, #b9f -1px 1px 2px;
	  font-family: arial;
	  line-height: .3;
	  cursor: pointer;
	}
	.CodeMirror-foldgutter {
	  width: .7em;
	}
	.CodeMirror-foldgutter-open,
	.CodeMirror-foldgutter-folded {
	  cursor: pointer;
	}
	.CodeMirror-foldgutter-open:after {
	  content: "\25BE";
	}
	.CodeMirror-foldgutter-folded:after {
	  content: "\25B8";
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.CodeMirror {
	  line-height: var(--jp-code-line-height);
	  font-size: var(--jp-code-font-size);
	  font-family: var(--jp-code-font-family);
	  border: 0;
	  border-radius: 0;
	  height: auto;
	  /* Changed to auto to autogrow */
	}
	
	.CodeMirror pre {
	  padding: 0 var(--jp-code-padding);
	}
	
	.jp-CodeMirrorEditor[data-type='inline'] .CodeMirror-dialog {
	  background-color: var(--jp-layout-color0);
	  color: var(--jp-content-font-color1);
	}
	
	/* This causes https://github.com/jupyter/jupyterlab/issues/522 */
	/* May not cause it not because we changed it! */
	.CodeMirror-lines {
	  padding: var(--jp-code-padding) 0;
	}
	
	.CodeMirror-linenumber {
	  padding: 0 8px;
	}
	
	.jp-CodeMirrorEditor {
	  cursor: text;
	}
	
	.jp-CodeMirrorEditor[data-type='inline'] .CodeMirror-cursor {
	  border-left: var(--jp-code-cursor-width0) solid var(--jp-editor-cursor-color);
	}
	
	/* When zoomed out 67% and 33% on a screen of 1440 width x 900 height */
	@media screen and (min-width: 2138px) and (max-width: 4319px) {
	  .jp-CodeMirrorEditor[data-type='inline'] .CodeMirror-cursor {
		border-left: var(--jp-code-cursor-width1) solid
		  var(--jp-editor-cursor-color);
	  }
	}
	
	/* When zoomed out less than 33% */
	@media screen and (min-width: 4320px) {
	  .jp-CodeMirrorEditor[data-type='inline'] .CodeMirror-cursor {
		border-left: var(--jp-code-cursor-width2) solid
		  var(--jp-editor-cursor-color);
	  }
	}
	
	.CodeMirror.jp-mod-readOnly .CodeMirror-cursor {
	  display: none;
	}
	
	.CodeMirror-gutters {
	  border-right: 1px solid var(--jp-border-color2);
	  background-color: var(--jp-layout-color0);
	}
	
	.jp-CollaboratorCursor {
	  border-left: 5px solid transparent;
	  border-right: 5px solid transparent;
	  border-top: none;
	  border-bottom: 3px solid;
	  background-clip: content-box;
	  margin-left: -5px;
	  margin-right: -5px;
	}
	
	.CodeMirror-selectedtext.cm-searching {
	  background-color: var(--jp-search-selected-match-background-color) !important;
	  color: var(--jp-search-selected-match-color) !important;
	}
	
	.cm-searching {
	  background-color: var(
		--jp-search-unselected-match-background-color
	  ) !important;
	  color: var(--jp-search-unselected-match-color) !important;
	}
	
	.CodeMirror-focused .CodeMirror-selected {
	  background-color: var(--jp-editor-selected-focused-background);
	}
	
	.CodeMirror-selected {
	  background-color: var(--jp-editor-selected-background);
	}
	
	.jp-CollaboratorCursor-hover {
	  position: absolute;
	  z-index: 1;
	  transform: translateX(-50%);
	  color: white;
	  border-radius: 3px;
	  padding-left: 4px;
	  padding-right: 4px;
	  padding-top: 1px;
	  padding-bottom: 1px;
	  text-align: center;
	  font-size: var(--jp-ui-font-size1);
	  white-space: nowrap;
	}
	
	.jp-CodeMirror-ruler {
	  border-left: 1px dashed var(--jp-border-color2);
	}
	
	/**
	 * Here is our jupyter theme for CodeMirror syntax highlighting
	 * This is used in our marked.js syntax highlighting and CodeMirror itself
	 * The string "jupyter" is set in ../codemirror/widget.DEFAULT_CODEMIRROR_THEME
	 * This came from the classic notebook, which came form highlight.js/GitHub
	 */
	
	/**
	 * CodeMirror themes are handling the background/color in this way. This works
	 * fine for CodeMirror editors outside the notebook, but the notebook styles
	 * these things differently.
	 */
	.CodeMirror.cm-s-jupyter {
	  background: var(--jp-layout-color0);
	  color: var(--jp-content-font-color1);
	}
	
	/* In the notebook, we want this styling to be handled by its container */
	.jp-CodeConsole .CodeMirror.cm-s-jupyter,
	.jp-Notebook .CodeMirror.cm-s-jupyter {
	  background: transparent;
	}
	
	.cm-s-jupyter .CodeMirror-cursor {
	  border-left: var(--jp-code-cursor-width0) solid var(--jp-editor-cursor-color);
	}
	.cm-s-jupyter span.cm-keyword {
	  color: var(--jp-mirror-editor-keyword-color);
	  font-weight: bold;
	}
	.cm-s-jupyter span.cm-atom {
	  color: var(--jp-mirror-editor-atom-color);
	}
	.cm-s-jupyter span.cm-number {
	  color: var(--jp-mirror-editor-number-color);
	}
	.cm-s-jupyter span.cm-def {
	  color: var(--jp-mirror-editor-def-color);
	}
	.cm-s-jupyter span.cm-variable {
	  color: var(--jp-mirror-editor-variable-color);
	}
	.cm-s-jupyter span.cm-variable-2 {
	  color: var(--jp-mirror-editor-variable-2-color);
	}
	.cm-s-jupyter span.cm-variable-3 {
	  color: var(--jp-mirror-editor-variable-3-color);
	}
	.cm-s-jupyter span.cm-punctuation {
	  color: var(--jp-mirror-editor-punctuation-color);
	}
	.cm-s-jupyter span.cm-property {
	  color: var(--jp-mirror-editor-property-color);
	}
	.cm-s-jupyter span.cm-operator {
	  color: var(--jp-mirror-editor-operator-color);
	  font-weight: bold;
	}
	.cm-s-jupyter span.cm-comment {
	  color: var(--jp-mirror-editor-comment-color);
	  font-style: italic;
	}
	.cm-s-jupyter span.cm-string {
	  color: var(--jp-mirror-editor-string-color);
	}
	.cm-s-jupyter span.cm-string-2 {
	  color: var(--jp-mirror-editor-string-2-color);
	}
	.cm-s-jupyter span.cm-meta {
	  color: var(--jp-mirror-editor-meta-color);
	}
	.cm-s-jupyter span.cm-qualifier {
	  color: var(--jp-mirror-editor-qualifier-color);
	}
	.cm-s-jupyter span.cm-builtin {
	  color: var(--jp-mirror-editor-builtin-color);
	}
	.cm-s-jupyter span.cm-bracket {
	  color: var(--jp-mirror-editor-bracket-color);
	}
	.cm-s-jupyter span.cm-tag {
	  color: var(--jp-mirror-editor-tag-color);
	}
	.cm-s-jupyter span.cm-attribute {
	  color: var(--jp-mirror-editor-attribute-color);
	}
	.cm-s-jupyter span.cm-header {
	  color: var(--jp-mirror-editor-header-color);
	}
	.cm-s-jupyter span.cm-quote {
	  color: var(--jp-mirror-editor-quote-color);
	}
	.cm-s-jupyter span.cm-link {
	  color: var(--jp-mirror-editor-link-color);
	}
	.cm-s-jupyter span.cm-error {
	  color: var(--jp-mirror-editor-error-color);
	}
	.cm-s-jupyter span.cm-hr {
	  color: #999;
	}
	
	.cm-s-jupyter span.cm-tab {
	  background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAMCAYAAAAkuj5RAAAAAXNSR0IArs4c6QAAAGFJREFUSMft1LsRQFAQheHPowAKoACx3IgEKtaEHujDjORSgWTH/ZOdnZOcM/sgk/kFFWY0qV8foQwS4MKBCS3qR6ixBJvElOobYAtivseIE120FaowJPN75GMu8j/LfMwNjh4HUpwg4LUAAAAASUVORK5CYII=);
	  background-position: right;
	  background-repeat: no-repeat;
	}
	
	.cm-s-jupyter .CodeMirror-activeline-background,
	.cm-s-jupyter .CodeMirror-gutter {
	  background-color: var(--jp-layout-color2);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| RenderedText
	|----------------------------------------------------------------------------*/
	
	:root {
	  /* This is the padding value to fill the gaps between lines containing spans with background color. */
	  --jp-private-code-span-padding: calc(
		(var(--jp-code-line-height) - 1) * var(--jp-code-font-size) / 2
	  );
	}
	
	.jp-RenderedText {
	  text-align: left;
	  padding-left: var(--jp-code-padding);
	  line-height: var(--jp-code-line-height);
	  font-family: var(--jp-code-font-family);
	}
	
	.jp-RenderedText pre,
	.jp-RenderedJavaScript pre,
	.jp-RenderedHTMLCommon pre {
	  color: var(--jp-content-font-color1);
	  font-size: var(--jp-code-font-size);
	  border: none;
	  margin: 0px;
	  padding: 0px;
	}
	
	.jp-RenderedText pre a:link {
	  text-decoration: none;
	  color: var(--jp-content-link-color);
	}
	.jp-RenderedText pre a:hover {
	  text-decoration: underline;
	  color: var(--jp-content-link-color);
	}
	.jp-RenderedText pre a:visited {
	  text-decoration: none;
	  color: var(--jp-content-link-color);
	}
	
	/* console foregrounds and backgrounds */
	.jp-RenderedText pre .ansi-black-fg {
	  color: #3e424d;
	}
	.jp-RenderedText pre .ansi-red-fg {
	  color: #e75c58;
	}
	.jp-RenderedText pre .ansi-green-fg {
	  color: #00a250;
	}
	.jp-RenderedText pre .ansi-yellow-fg {
	  color: #ddb62b;
	}
	.jp-RenderedText pre .ansi-blue-fg {
	  color: #208ffb;
	}
	.jp-RenderedText pre .ansi-magenta-fg {
	  color: #d160c4;
	}
	.jp-RenderedText pre .ansi-cyan-fg {
	  color: #60c6c8;
	}
	.jp-RenderedText pre .ansi-white-fg {
	  color: #c5c1b4;
	}
	
	.jp-RenderedText pre .ansi-black-bg {
	  background-color: #3e424d;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-red-bg {
	  background-color: #e75c58;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-green-bg {
	  background-color: #00a250;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-yellow-bg {
	  background-color: #ddb62b;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-blue-bg {
	  background-color: #208ffb;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-magenta-bg {
	  background-color: #d160c4;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-cyan-bg {
	  background-color: #60c6c8;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-white-bg {
	  background-color: #c5c1b4;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	
	.jp-RenderedText pre .ansi-black-intense-fg {
	  color: #282c36;
	}
	.jp-RenderedText pre .ansi-red-intense-fg {
	  color: #b22b31;
	}
	.jp-RenderedText pre .ansi-green-intense-fg {
	  color: #007427;
	}
	.jp-RenderedText pre .ansi-yellow-intense-fg {
	  color: #b27d12;
	}
	.jp-RenderedText pre .ansi-blue-intense-fg {
	  color: #0065ca;
	}
	.jp-RenderedText pre .ansi-magenta-intense-fg {
	  color: #a03196;
	}
	.jp-RenderedText pre .ansi-cyan-intense-fg {
	  color: #258f8f;
	}
	.jp-RenderedText pre .ansi-white-intense-fg {
	  color: #a1a6b2;
	}
	
	.jp-RenderedText pre .ansi-black-intense-bg {
	  background-color: #282c36;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-red-intense-bg {
	  background-color: #b22b31;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-green-intense-bg {
	  background-color: #007427;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-yellow-intense-bg {
	  background-color: #b27d12;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-blue-intense-bg {
	  background-color: #0065ca;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-magenta-intense-bg {
	  background-color: #a03196;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-cyan-intense-bg {
	  background-color: #258f8f;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	.jp-RenderedText pre .ansi-white-intense-bg {
	  background-color: #a1a6b2;
	  padding: var(--jp-private-code-span-padding) 0;
	}
	
	.jp-RenderedText pre .ansi-default-inverse-fg {
	  color: var(--jp-ui-inverse-font-color0);
	}
	.jp-RenderedText pre .ansi-default-inverse-bg {
	  background-color: var(--jp-inverse-layout-color0);
	  padding: var(--jp-private-code-span-padding) 0;
	}
	
	.jp-RenderedText pre .ansi-bold {
	  font-weight: bold;
	}
	.jp-RenderedText pre .ansi-underline {
	  text-decoration: underline;
	}
	
	.jp-RenderedText[data-mime-type='application/vnd.jupyter.stderr'] {
	  background: var(--jp-rendermime-error-background);
	  padding-top: var(--jp-code-padding);
	}
	
	/*-----------------------------------------------------------------------------
	| RenderedLatex
	|----------------------------------------------------------------------------*/
	
	.jp-RenderedLatex {
	  color: var(--jp-content-font-color1);
	  font-size: var(--jp-content-font-size1);
	  line-height: var(--jp-content-line-height);
	}
	
	/* Left-justify outputs.*/
	.jp-OutputArea-output.jp-RenderedLatex {
	  padding: var(--jp-code-padding);
	  text-align: left;
	}
	
	/*-----------------------------------------------------------------------------
	| RenderedHTML
	|----------------------------------------------------------------------------*/
	
	.jp-RenderedHTMLCommon {
	  color: var(--jp-content-font-color1);
	  font-family: var(--jp-content-font-family);
	  font-size: var(--jp-content-font-size1);
	  line-height: var(--jp-content-line-height);
	  /* Give a bit more R padding on Markdown text to keep line lengths reasonable */
	  padding-right: 20px;
	}
	
	.jp-RenderedHTMLCommon em {
	  font-style: italic;
	}
	
	.jp-RenderedHTMLCommon strong {
	  font-weight: bold;
	}
	
	.jp-RenderedHTMLCommon u {
	  text-decoration: underline;
	}
	
	.jp-RenderedHTMLCommon a:link {
	  text-decoration: none;
	  color: var(--jp-content-link-color);
	}
	
	.jp-RenderedHTMLCommon a:hover {
	  text-decoration: underline;
	  color: var(--jp-content-link-color);
	}
	
	.jp-RenderedHTMLCommon a:visited {
	  text-decoration: none;
	  color: var(--jp-content-link-color);
	}
	
	/* Headings */
	
	.jp-RenderedHTMLCommon h1,
	.jp-RenderedHTMLCommon h2,
	.jp-RenderedHTMLCommon h3,
	.jp-RenderedHTMLCommon h4,
	.jp-RenderedHTMLCommon h5,
	.jp-RenderedHTMLCommon h6 {
	  line-height: var(--jp-content-heading-line-height);
	  font-weight: var(--jp-content-heading-font-weight);
	  font-style: normal;
	  margin: var(--jp-content-heading-margin-top) 0
		var(--jp-content-heading-margin-bottom) 0;
	}
	
	.jp-RenderedHTMLCommon h1:first-child,
	.jp-RenderedHTMLCommon h2:first-child,
	.jp-RenderedHTMLCommon h3:first-child,
	.jp-RenderedHTMLCommon h4:first-child,
	.jp-RenderedHTMLCommon h5:first-child,
	.jp-RenderedHTMLCommon h6:first-child {
	  margin-top: calc(0.5 * var(--jp-content-heading-margin-top));
	}
	
	.jp-RenderedHTMLCommon h1:last-child,
	.jp-RenderedHTMLCommon h2:last-child,
	.jp-RenderedHTMLCommon h3:last-child,
	.jp-RenderedHTMLCommon h4:last-child,
	.jp-RenderedHTMLCommon h5:last-child,
	.jp-RenderedHTMLCommon h6:last-child {
	  margin-bottom: calc(0.5 * var(--jp-content-heading-margin-bottom));
	}
	
	.jp-RenderedHTMLCommon h1 {
	  font-size: var(--jp-content-font-size5);
	}
	
	.jp-RenderedHTMLCommon h2 {
	  font-size: var(--jp-content-font-size4);
	}
	
	.jp-RenderedHTMLCommon h3 {
	  font-size: var(--jp-content-font-size3);
	}
	
	.jp-RenderedHTMLCommon h4 {
	  font-size: var(--jp-content-font-size2);
	}
	
	.jp-RenderedHTMLCommon h5 {
	  font-size: var(--jp-content-font-size1);
	}
	
	.jp-RenderedHTMLCommon h6 {
	  font-size: var(--jp-content-font-size0);
	}
	
	/* Lists */
	
	.jp-RenderedHTMLCommon ul:not(.list-inline),
	.jp-RenderedHTMLCommon ol:not(.list-inline) {
	  padding-left: 2em;
	}
	
	.jp-RenderedHTMLCommon ul {
	  list-style: disc;
	}
	
	.jp-RenderedHTMLCommon ul ul {
	  list-style: square;
	}
	
	.jp-RenderedHTMLCommon ul ul ul {
	  list-style: circle;
	}
	
	.jp-RenderedHTMLCommon ol {
	  list-style: decimal;
	}
	
	.jp-RenderedHTMLCommon ol ol {
	  list-style: upper-alpha;
	}
	
	.jp-RenderedHTMLCommon ol ol ol {
	  list-style: lower-alpha;
	}
	
	.jp-RenderedHTMLCommon ol ol ol ol {
	  list-style: lower-roman;
	}
	
	.jp-RenderedHTMLCommon ol ol ol ol ol {
	  list-style: decimal;
	}
	
	.jp-RenderedHTMLCommon ol,
	.jp-RenderedHTMLCommon ul {
	  margin-bottom: 1em;
	}
	
	.jp-RenderedHTMLCommon ul ul,
	.jp-RenderedHTMLCommon ul ol,
	.jp-RenderedHTMLCommon ol ul,
	.jp-RenderedHTMLCommon ol ol {
	  margin-bottom: 0em;
	}
	
	.jp-RenderedHTMLCommon hr {
	  color: var(--jp-border-color2);
	  background-color: var(--jp-border-color1);
	  margin-top: 1em;
	  margin-bottom: 1em;
	}
	
	.jp-RenderedHTMLCommon > pre {
	  margin: 1.5em 2em;
	}
	
	.jp-RenderedHTMLCommon pre,
	.jp-RenderedHTMLCommon code {
	  border: 0;
	  background-color: var(--jp-layout-color0);
	  color: var(--jp-content-font-color1);
	  font-family: var(--jp-code-font-family);
	  font-size: inherit;
	  line-height: var(--jp-code-line-height);
	  padding: 0;
	  white-space: pre-wrap;
	}
	
	.jp-RenderedHTMLCommon :not(pre) > code {
	  background-color: var(--jp-layout-color2);
	  padding: 1px 5px;
	}
	
	/* Tables */
	
	.jp-RenderedHTMLCommon table {
	  border-collapse: collapse;
	  border-spacing: 0;
	  border: none;
	  color: var(--jp-ui-font-color1);
	  font-size: 12px;
	  table-layout: fixed;
	  margin-left: auto;
	  margin-right: auto;
	}
	
	.jp-RenderedHTMLCommon thead {
	  border-bottom: var(--jp-border-width) solid var(--jp-border-color1);
	  vertical-align: bottom;
	}
	
	.jp-RenderedHTMLCommon td,
	.jp-RenderedHTMLCommon th,
	.jp-RenderedHTMLCommon tr {
	  vertical-align: middle;
	  padding: 0.5em 0.5em;
	  line-height: normal;
	  white-space: normal;
	  max-width: none;
	  border: none;
	}
	
	.jp-RenderedMarkdown.jp-RenderedHTMLCommon td,
	.jp-RenderedMarkdown.jp-RenderedHTMLCommon th {
	  max-width: none;
	}
	
	:not(.jp-RenderedMarkdown).jp-RenderedHTMLCommon td,
	:not(.jp-RenderedMarkdown).jp-RenderedHTMLCommon th,
	:not(.jp-RenderedMarkdown).jp-RenderedHTMLCommon tr {
	  text-align: right;
	}
	
	.jp-RenderedHTMLCommon th {
	  font-weight: bold;
	}
	
	.jp-RenderedHTMLCommon tbody tr:nth-child(odd) {
	  background: var(--jp-layout-color0);
	}
	
	.jp-RenderedHTMLCommon tbody tr:nth-child(even) {
	  background: var(--jp-rendermime-table-row-background);
	}
	
	.jp-RenderedHTMLCommon tbody tr:hover {
	  background: var(--jp-rendermime-table-row-hover-background);
	}
	
	.jp-RenderedHTMLCommon table {
	  margin-bottom: 1em;
	}
	
	.jp-RenderedHTMLCommon p {
	  text-align: left;
	  margin: 0px;
	}
	
	.jp-RenderedHTMLCommon p {
	  margin-bottom: 1em;
	}
	
	.jp-RenderedHTMLCommon img {
	  -moz-force-broken-image-icon: 1;
	}
	
	/* Restrict to direct children as other images could be nested in other content. */
	.jp-RenderedHTMLCommon > img {
	  display: block;
	  margin-left: 0;
	  margin-right: 0;
	  margin-bottom: 1em;
	}
	
	/* Change color behind transparent images if they need it... */
	[data-jp-theme-light='false'] .jp-RenderedImage img.jp-needs-light-background {
	  background-color: var(--jp-inverse-layout-color1);
	}
	[data-jp-theme-light='true'] .jp-RenderedImage img.jp-needs-dark-background {
	  background-color: var(--jp-inverse-layout-color1);
	}
	/* ...or leave it untouched if they don't */
	[data-jp-theme-light='false'] .jp-RenderedImage img.jp-needs-dark-background {
	}
	[data-jp-theme-light='true'] .jp-RenderedImage img.jp-needs-light-background {
	}
	
	.jp-RenderedHTMLCommon img,
	.jp-RenderedImage img,
	.jp-RenderedHTMLCommon svg,
	.jp-RenderedSVG svg {
	  max-width: 100%;
	  height: auto;
	}
	
	.jp-RenderedHTMLCommon img.jp-mod-unconfined,
	.jp-RenderedImage img.jp-mod-unconfined,
	.jp-RenderedHTMLCommon svg.jp-mod-unconfined,
	.jp-RenderedSVG svg.jp-mod-unconfined {
	  max-width: none;
	}
	
	.jp-RenderedHTMLCommon .alert {
	  padding: var(--jp-notebook-padding);
	  border: var(--jp-border-width) solid transparent;
	  border-radius: var(--jp-border-radius);
	  margin-bottom: 1em;
	}
	
	.jp-RenderedHTMLCommon .alert-info {
	  color: var(--jp-info-color0);
	  background-color: var(--jp-info-color3);
	  border-color: var(--jp-info-color2);
	}
	.jp-RenderedHTMLCommon .alert-info hr {
	  border-color: var(--jp-info-color3);
	}
	.jp-RenderedHTMLCommon .alert-info > p:last-child,
	.jp-RenderedHTMLCommon .alert-info > ul:last-child {
	  margin-bottom: 0;
	}
	
	.jp-RenderedHTMLCommon .alert-warning {
	  color: var(--jp-warn-color0);
	  background-color: var(--jp-warn-color3);
	  border-color: var(--jp-warn-color2);
	}
	.jp-RenderedHTMLCommon .alert-warning hr {
	  border-color: var(--jp-warn-color3);
	}
	.jp-RenderedHTMLCommon .alert-warning > p:last-child,
	.jp-RenderedHTMLCommon .alert-warning > ul:last-child {
	  margin-bottom: 0;
	}
	
	.jp-RenderedHTMLCommon .alert-success {
	  color: var(--jp-success-color0);
	  background-color: var(--jp-success-color3);
	  border-color: var(--jp-success-color2);
	}
	.jp-RenderedHTMLCommon .alert-success hr {
	  border-color: var(--jp-success-color3);
	}
	.jp-RenderedHTMLCommon .alert-success > p:last-child,
	.jp-RenderedHTMLCommon .alert-success > ul:last-child {
	  margin-bottom: 0;
	}
	
	.jp-RenderedHTMLCommon .alert-danger {
	  color: var(--jp-error-color0);
	  background-color: var(--jp-error-color3);
	  border-color: var(--jp-error-color2);
	}
	.jp-RenderedHTMLCommon .alert-danger hr {
	  border-color: var(--jp-error-color3);
	}
	.jp-RenderedHTMLCommon .alert-danger > p:last-child,
	.jp-RenderedHTMLCommon .alert-danger > ul:last-child {
	  margin-bottom: 0;
	}
	
	.jp-RenderedHTMLCommon blockquote {
	  margin: 1em 2em;
	  padding: 0 1em;
	  border-left: 5px solid var(--jp-border-color2);
	}
	
	a.jp-InternalAnchorLink {
	  visibility: hidden;
	  margin-left: 8px;
	  color: var(--md-blue-800);
	}
	
	h1:hover .jp-InternalAnchorLink,
	h2:hover .jp-InternalAnchorLink,
	h3:hover .jp-InternalAnchorLink,
	h4:hover .jp-InternalAnchorLink,
	h5:hover .jp-InternalAnchorLink,
	h6:hover .jp-InternalAnchorLink {
	  visibility: visible;
	}
	
	.jp-RenderedHTMLCommon kbd {
	  background-color: var(--jp-rendermime-table-row-background);
	  border: 1px solid var(--jp-border-color0);
	  border-bottom-color: var(--jp-border-color2);
	  border-radius: 3px;
	  box-shadow: inset 0 -1px 0 rgba(0, 0, 0, 0.25);
	  display: inline-block;
	  font-size: 0.8em;
	  line-height: 1em;
	  padding: 0.2em 0.5em;
	}
	
	/* Most direct children of .jp-RenderedHTMLCommon have a margin-bottom of 1.0.
	 * At the bottom of cells this is a bit too much as there is also spacing
	 * between cells. Going all the way to 0 gets too tight between markdown and
	 * code cells.
	 */
	.jp-RenderedHTMLCommon > *:last-child {
	  margin-bottom: 0.5em;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-MimeDocument {
	  outline: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Variables
	|----------------------------------------------------------------------------*/
	
	:root {
	  --jp-private-filebrowser-button-height: 28px;
	  --jp-private-filebrowser-button-width: 48px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-FileBrowser {
	  display: flex;
	  flex-direction: column;
	  color: var(--jp-ui-font-color1);
	  background: var(--jp-layout-color1);
	  /* This is needed so that all font sizing of children done in ems is
	   * relative to this base size */
	  font-size: var(--jp-ui-font-size1);
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar {
	  border-bottom: none;
	  height: auto;
	  margin: var(--jp-toolbar-header-margin);
	  box-shadow: none;
	}
	
	.jp-BreadCrumbs {
	  flex: 0 0 auto;
	  margin: 8px 12px 8px 12px;
	}
	
	.jp-BreadCrumbs-item {
	  margin: 0px 2px;
	  padding: 0px 2px;
	  border-radius: var(--jp-border-radius);
	  cursor: pointer;
	}
	
	.jp-BreadCrumbs-item:hover {
	  background-color: var(--jp-layout-color2);
	}
	
	.jp-BreadCrumbs-item:first-child {
	  margin-left: 0px;
	}
	
	.jp-BreadCrumbs-item.jp-mod-dropTarget {
	  background-color: var(--jp-brand-color2);
	  opacity: 0.7;
	}
	
	/*-----------------------------------------------------------------------------
	| Buttons
	|----------------------------------------------------------------------------*/
	
	.jp-FileBrowser-toolbar.jp-Toolbar {
	  padding: 0px;
	  margin: 8px 12px 0px 12px;
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar {
	  justify-content: flex-start;
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar .jp-Toolbar-item {
	  flex: 0 0 auto;
	  padding-left: 0px;
	  padding-right: 2px;
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar .jp-ToolbarButtonComponent {
	  width: 40px;
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar
	  .jp-Toolbar-item:first-child
	  .jp-ToolbarButtonComponent {
	  width: 72px;
	  background: var(--jp-brand-color1);
	}
	
	.jp-FileBrowser-toolbar.jp-Toolbar
	  .jp-Toolbar-item:first-child
	  .jp-ToolbarButtonComponent
	  .jp-icon3 {
	  fill: white;
	}
	
	/*-----------------------------------------------------------------------------
	| Other styles
	|----------------------------------------------------------------------------*/
	
	.jp-FileDialog.jp-mod-conflict input {
	  color: red;
	}
	
	.jp-FileDialog .jp-new-name-title {
	  margin-top: 12px;
	}
	
	.jp-LastModified-hidden {
	  display: none;
	}
	
	.jp-FileBrowser-filterBox {
	  padding: 0px;
	  flex: 0 0 auto;
	  margin: 8px 12px 0px 12px;
	}
	
	/*-----------------------------------------------------------------------------
	| DirListing
	|----------------------------------------------------------------------------*/
	
	.jp-DirListing {
	  flex: 1 1 auto;
	  display: flex;
	  flex-direction: column;
	  outline: 0;
	}
	
	.jp-DirListing-header {
	  flex: 0 0 auto;
	  display: flex;
	  flex-direction: row;
	  overflow: hidden;
	  border-top: var(--jp-border-width) solid var(--jp-border-color2);
	  border-bottom: var(--jp-border-width) solid var(--jp-border-color1);
	  box-shadow: var(--jp-toolbar-box-shadow);
	  z-index: 2;
	}
	
	.jp-DirListing-headerItem {
	  padding: 4px 12px 2px 12px;
	  font-weight: 500;
	}
	
	.jp-DirListing-headerItem:hover {
	  background: var(--jp-layout-color2);
	}
	
	.jp-DirListing-headerItem.jp-id-name {
	  flex: 1 0 84px;
	}
	
	.jp-DirListing-headerItem.jp-id-modified {
	  flex: 0 0 112px;
	  border-left: var(--jp-border-width) solid var(--jp-border-color2);
	  text-align: right;
	}
	
	.jp-id-narrow {
	  display: none;
	  flex: 0 0 5px;
	  padding: 4px 4px;
	  border-left: var(--jp-border-width) solid var(--jp-border-color2);
	  text-align: right;
	  color: var(--jp-border-color2);
	}
	
	.jp-DirListing-narrow .jp-id-narrow {
	  display: block;
	}
	
	.jp-DirListing-narrow .jp-id-modified,
	.jp-DirListing-narrow .jp-DirListing-itemModified {
	  display: none;
	}
	
	.jp-DirListing-headerItem.jp-mod-selected {
	  font-weight: 600;
	}
	
	/* increase specificity to override bundled default */
	.jp-DirListing-content {
	  flex: 1 1 auto;
	  margin: 0;
	  padding: 0;
	  list-style-type: none;
	  overflow: auto;
	  background-color: var(--jp-layout-color1);
	}
	
	.jp-DirListing-content mark {
	  color: var(--jp-ui-font-color0);
	  background-color: transparent;
	  font-weight: bold;
	}
	
	/* Style the directory listing content when a user drops a file to upload */
	.jp-DirListing.jp-mod-native-drop .jp-DirListing-content {
	  outline: 5px dashed rgba(128, 128, 128, 0.5);
	  outline-offset: -10px;
	  cursor: copy;
	}
	
	.jp-DirListing-item {
	  display: flex;
	  flex-direction: row;
	  padding: 4px 12px;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	.jp-DirListing-item[data-is-dot] {
	  opacity: 75%;
	}
	
	.jp-DirListing-item.jp-mod-selected {
	  color: white;
	  background: var(--jp-brand-color1);
	}
	
	.jp-DirListing-item.jp-mod-dropTarget {
	  background: var(--jp-brand-color3);
	}
	
	.jp-DirListing-item:hover:not(.jp-mod-selected) {
	  background: var(--jp-layout-color2);
	}
	
	.jp-DirListing-itemIcon {
	  flex: 0 0 20px;
	  margin-right: 4px;
	}
	
	.jp-DirListing-itemText {
	  flex: 1 0 64px;
	  white-space: nowrap;
	  overflow: hidden;
	  text-overflow: ellipsis;
	  user-select: none;
	}
	
	.jp-DirListing-itemModified {
	  flex: 0 0 125px;
	  text-align: right;
	}
	
	.jp-DirListing-editor {
	  flex: 1 0 64px;
	  outline: none;
	  border: none;
	}
	
	.jp-DirListing-item.jp-mod-running .jp-DirListing-itemIcon:before {
	  color: limegreen;
	  content: '\25CF';
	  font-size: 8px;
	  position: absolute;
	  left: -8px;
	}
	
	.jp-DirListing-item.lm-mod-drag-image,
	.jp-DirListing-item.jp-mod-selected.lm-mod-drag-image {
	  font-size: var(--jp-ui-font-size1);
	  padding-left: 4px;
	  margin-left: 4px;
	  width: 160px;
	  background-color: var(--jp-ui-inverse-font-color2);
	  box-shadow: var(--jp-elevation-z2);
	  border-radius: 0px;
	  color: var(--jp-ui-font-color1);
	  transform: translateX(-40%) translateY(-58%);
	}
	
	.jp-DirListing-deadSpace {
	  flex: 1 1 auto;
	  margin: 0;
	  padding: 0;
	  list-style-type: none;
	  overflow: auto;
	  background-color: var(--jp-layout-color1);
	}
	
	.jp-Document {
	  min-width: 120px;
	  min-height: 120px;
	  outline: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Private CSS variables
	|----------------------------------------------------------------------------*/
	
	:root {
	}
	
	/*-----------------------------------------------------------------------------
	| Main OutputArea
	| OutputArea has a list of Outputs
	|----------------------------------------------------------------------------*/
	
	.jp-OutputArea {
	  overflow-y: auto;
	}
	
	.jp-OutputArea-child {
	  display: flex;
	  flex-direction: row;
	}
	
	.jp-OutputPrompt {
	  flex: 0 0 var(--jp-cell-prompt-width);
	  color: var(--jp-cell-outprompt-font-color);
	  font-family: var(--jp-cell-prompt-font-family);
	  padding: var(--jp-code-padding);
	  letter-spacing: var(--jp-cell-prompt-letter-spacing);
	  line-height: var(--jp-code-line-height);
	  font-size: var(--jp-code-font-size);
	  border: var(--jp-border-width) solid transparent;
	  opacity: var(--jp-cell-prompt-opacity);
	  /* Right align prompt text, don't wrap to handle large prompt numbers */
	  text-align: right;
	  white-space: nowrap;
	  overflow: hidden;
	  text-overflow: ellipsis;
	  /* Disable text selection */
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	.jp-OutputArea-output {
	  height: auto;
	  overflow: auto;
	  user-select: text;
	  -moz-user-select: text;
	  -webkit-user-select: text;
	  -ms-user-select: text;
	}
	
	.jp-OutputArea-child .jp-OutputArea-output {
	  flex-grow: 1;
	  flex-shrink: 1;
	}
	
	/**
	 * Isolated output.
	 */
	.jp-OutputArea-output.jp-mod-isolated {
	  width: 100%;
	  display: block;
	}
	
	body.lm-mod-override-cursor .jp-OutputArea-output.jp-mod-isolated {
	  position: relative;
	}
	
	body.lm-mod-override-cursor .jp-OutputArea-output.jp-mod-isolated:before {
	  content: '';
	  position: absolute;
	  top: 0;
	  left: 0;
	  right: 0;
	  bottom: 0;
	  background: transparent;
	}
	
	/* pre */
	
	.jp-OutputArea-output pre {
	  border: none;
	  margin: 0px;
	  padding: 0px;
	  overflow-x: auto;
	  overflow-y: auto;
	  word-break: break-all;
	  word-wrap: break-word;
	  white-space: pre-wrap;
	}
	
	/* tables */
	
	.jp-OutputArea-output.jp-RenderedHTMLCommon table {
	  margin-left: 0;
	  margin-right: 0;
	}
	
	/* description lists */
	
	.jp-OutputArea-output dl,
	.jp-OutputArea-output dt,
	.jp-OutputArea-output dd {
	  display: block;
	}
	
	.jp-OutputArea-output dl {
	  width: 100%;
	  overflow: hidden;
	  padding: 0;
	  margin: 0;
	}
	
	.jp-OutputArea-output dt {
	  font-weight: bold;
	  float: left;
	  width: 20%;
	  padding: 0;
	  margin: 0;
	}
	
	.jp-OutputArea-output dd {
	  float: left;
	  width: 80%;
	  padding: 0;
	  margin: 0;
	}
	
	/* Hide the gutter in case of
	 *  - nested output areas (e.g. in the case of output widgets)
	 *  - mirrored output areas
	 */
	.jp-OutputArea .jp-OutputArea .jp-OutputArea-prompt {
	  display: none;
	}
	
	/*-----------------------------------------------------------------------------
	| executeResult is added to any Output-result for the display of the object
	| returned by a cell
	|----------------------------------------------------------------------------*/
	
	.jp-OutputArea-output.jp-OutputArea-executeResult {
	  margin-left: 0px;
	  flex: 1 1 auto;
	}
	
	/* Text output with the Out[] prompt needs a top padding to match the
	 * alignment of the Out[] prompt itself.
	 */
	.jp-OutputArea-executeResult .jp-RenderedText.jp-OutputArea-output {
	  padding-top: var(--jp-code-padding);
	  border-top: var(--jp-border-width) solid transparent;
	}
	
	/*-----------------------------------------------------------------------------
	| The Stdin output
	|----------------------------------------------------------------------------*/
	
	.jp-OutputArea-stdin {
	  line-height: var(--jp-code-line-height);
	  padding-top: var(--jp-code-padding);
	  display: flex;
	}
	
	.jp-Stdin-prompt {
	  color: var(--jp-content-font-color0);
	  padding-right: var(--jp-code-padding);
	  vertical-align: baseline;
	  flex: 0 0 auto;
	}
	
	.jp-Stdin-input {
	  font-family: var(--jp-code-font-family);
	  font-size: inherit;
	  color: inherit;
	  background-color: inherit;
	  width: 42%;
	  min-width: 200px;
	  /* make sure input baseline aligns with prompt */
	  vertical-align: baseline;
	  /* padding + margin = 0.5em between prompt and cursor */
	  padding: 0em 0.25em;
	  margin: 0em 0.25em;
	  flex: 0 0 70%;
	}
	
	.jp-Stdin-input:focus {
	  box-shadow: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Output Area View
	|----------------------------------------------------------------------------*/
	
	.jp-LinkedOutputView .jp-OutputArea {
	  height: 100%;
	  display: block;
	}
	
	.jp-LinkedOutputView .jp-OutputArea-output:only-child {
	  height: 100%;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	.jp-Collapser {
	  flex: 0 0 var(--jp-cell-collapser-width);
	  padding: 0px;
	  margin: 0px;
	  border: none;
	  outline: none;
	  background: transparent;
	  border-radius: var(--jp-border-radius);
	  opacity: 1;
	}
	
	.jp-Collapser-child {
	  display: block;
	  width: 100%;
	  box-sizing: border-box;
	  /* height: 100% doesn't work because the height of its parent is computed from content */
	  position: absolute;
	  top: 0px;
	  bottom: 0px;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Header/Footer
	|----------------------------------------------------------------------------*/
	
	/* Hidden by zero height by default */
	.jp-CellHeader,
	.jp-CellFooter {
	  height: 0px;
	  width: 100%;
	  padding: 0px;
	  margin: 0px;
	  border: none;
	  outline: none;
	  background: transparent;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Input
	|----------------------------------------------------------------------------*/
	
	/* All input areas */
	.jp-InputArea {
	  display: flex;
	  flex-direction: row;
	  overflow: hidden;
	}
	
	.jp-InputArea-editor {
	  flex: 1 1 auto;
	  overflow: hidden;
	}
	
	.jp-InputArea-editor {
	  /* This is the non-active, default styling */
	  border: var(--jp-border-width) solid var(--jp-cell-editor-border-color);
	  border-radius: 0px;
	  background: var(--jp-cell-editor-background);
	}
	
	.jp-InputPrompt {
	  flex: 0 0 var(--jp-cell-prompt-width);
	  color: var(--jp-cell-inprompt-font-color);
	  font-family: var(--jp-cell-prompt-font-family);
	  padding: var(--jp-code-padding);
	  letter-spacing: var(--jp-cell-prompt-letter-spacing);
	  opacity: var(--jp-cell-prompt-opacity);
	  line-height: var(--jp-code-line-height);
	  font-size: var(--jp-code-font-size);
	  border: var(--jp-border-width) solid transparent;
	  opacity: var(--jp-cell-prompt-opacity);
	  /* Right align prompt text, don't wrap to handle large prompt numbers */
	  text-align: right;
	  white-space: nowrap;
	  overflow: hidden;
	  text-overflow: ellipsis;
	  /* Disable text selection */
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Placeholder
	|----------------------------------------------------------------------------*/
	
	.jp-Placeholder {
	  display: flex;
	  flex-direction: row;
	  flex: 1 1 auto;
	}
	
	.jp-Placeholder-prompt {
	  box-sizing: border-box;
	}
	
	.jp-Placeholder-content {
	  flex: 1 1 auto;
	  border: none;
	  background: transparent;
	  height: 20px;
	  box-sizing: border-box;
	}
	
	.jp-Placeholder-content .jp-MoreHorizIcon {
	  width: 32px;
	  height: 16px;
	  border: 1px solid transparent;
	  border-radius: var(--jp-border-radius);
	}
	
	.jp-Placeholder-content .jp-MoreHorizIcon:hover {
	  border: 1px solid var(--jp-border-color1);
	  box-shadow: 0px 0px 2px 0px rgba(0, 0, 0, 0.25);
	  background-color: var(--jp-layout-color0);
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Private CSS variables
	|----------------------------------------------------------------------------*/
	
	:root {
	  --jp-private-cell-scrolling-output-offset: 5px;
	}
	
	/*-----------------------------------------------------------------------------
	| Cell
	|----------------------------------------------------------------------------*/
	
	.jp-Cell {
	  padding: var(--jp-cell-padding);
	  margin: 0px;
	  border: none;
	  outline: none;
	  background: transparent;
	}
	
	/*-----------------------------------------------------------------------------
	| Common input/output
	|----------------------------------------------------------------------------*/
	
	.jp-Cell-inputWrapper,
	.jp-Cell-outputWrapper {
	  display: flex;
	  flex-direction: row;
	  padding: 0px;
	  margin: 0px;
	  /* Added to reveal the box-shadow on the input and output collapsers. */
	  overflow: visible;
	}
	
	/* Only input/output areas inside cells */
	.jp-Cell-inputArea,
	.jp-Cell-outputArea {
	  flex: 1 1 auto;
	}
	
	/*-----------------------------------------------------------------------------
	| Collapser
	|----------------------------------------------------------------------------*/
	
	/* Make the output collapser disappear when there is not output, but do so
	 * in a manner that leaves it in the layout and preserves its width.
	 */
	.jp-Cell.jp-mod-noOutputs .jp-Cell-outputCollapser {
	  border: none !important;
	  background: transparent !important;
	}
	
	.jp-Cell:not(.jp-mod-noOutputs) .jp-Cell-outputCollapser {
	  min-height: var(--jp-cell-collapser-min-height);
	}
	
	/*-----------------------------------------------------------------------------
	| Output
	|----------------------------------------------------------------------------*/
	
	/* Put a space between input and output when there IS output */
	.jp-Cell:not(.jp-mod-noOutputs) .jp-Cell-outputWrapper {
	  margin-top: 5px;
	}
	
	.jp-CodeCell.jp-mod-outputsScrolled .jp-Cell-outputArea {
	  overflow-y: auto;
	  max-height: 200px;
	  box-shadow: inset 0 0 6px 2px rgba(0, 0, 0, 0.3);
	  margin-left: var(--jp-private-cell-scrolling-output-offset);
	}
	
	.jp-CodeCell.jp-mod-outputsScrolled .jp-OutputArea-prompt {
	  flex: 0 0
		calc(
		  var(--jp-cell-prompt-width) -
			var(--jp-private-cell-scrolling-output-offset)
		);
	}
	
	/*-----------------------------------------------------------------------------
	| CodeCell
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| MarkdownCell
	|----------------------------------------------------------------------------*/
	
	.jp-MarkdownOutput {
	  flex: 1 1 auto;
	  margin-top: 0;
	  margin-bottom: 0;
	  padding-left: var(--jp-code-padding);
	}
	
	.jp-MarkdownOutput.jp-RenderedHTMLCommon {
	  overflow: auto;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Variables
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	
	/*-----------------------------------------------------------------------------
	| Styles
	|----------------------------------------------------------------------------*/
	
	.jp-NotebookPanel-toolbar {
	  padding: 2px;
	}
	
	.jp-Toolbar-item.jp-Notebook-toolbarCellType .jp-select-wrapper.jp-mod-focused {
	  border: none;
	  box-shadow: none;
	}
	
	.jp-Notebook-toolbarCellTypeDropdown select {
	  height: 24px;
	  font-size: var(--jp-ui-font-size1);
	  line-height: 14px;
	  border-radius: 0;
	  display: block;
	}
	
	.jp-Notebook-toolbarCellTypeDropdown span {
	  top: 5px !important;
	}
	
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Private CSS variables
	|----------------------------------------------------------------------------*/
	
	:root {
	  --jp-private-notebook-dragImage-width: 304px;
	  --jp-private-notebook-dragImage-height: 36px;
	  --jp-private-notebook-selected-color: var(--md-blue-400);
	  --jp-private-notebook-active-color: var(--md-green-400);
	}
	
	/*-----------------------------------------------------------------------------
	| Imports
	|----------------------------------------------------------------------------*/
	
	/*-----------------------------------------------------------------------------
	| Notebook
	|----------------------------------------------------------------------------*/
	
	.jp-NotebookPanel {
	  display: block;
	  height: 100%;
	}
	
	.jp-NotebookPanel.jp-Document {
	  min-width: 240px;
	  min-height: 120px;
	}
	
	.jp-Notebook {
	  padding: var(--jp-notebook-padding);
	  outline: none;
	  overflow: auto;
	  background: var(--jp-layout-color0);
	}
	
	.jp-Notebook.jp-mod-scrollPastEnd::after {
	  display: block;
	  content: '';
	  min-height: var(--jp-notebook-scroll-padding);
	}
	
	.jp-Notebook .jp-Cell {
	  overflow: visible;
	}
	
	.jp-Notebook .jp-Cell .jp-InputPrompt {
	  cursor: move;
	}
	
	/*-----------------------------------------------------------------------------
	| Notebook state related styling
	|
	| The notebook and cells each have states, here are the possibilities:
	|
	| - Notebook
	|   - Command
	|   - Edit
	| - Cell
	|   - None
	|   - Active (only one can be active)
	|   - Selected (the cells actions are applied to)
	|   - Multiselected (when multiple selected, the cursor)
	|   - No outputs
	|----------------------------------------------------------------------------*/
	
	/* Command or edit modes */
	
	.jp-Notebook .jp-Cell:not(.jp-mod-active) .jp-InputPrompt {
	  opacity: var(--jp-cell-prompt-not-active-opacity);
	  color: var(--jp-cell-prompt-not-active-font-color);
	}
	
	.jp-Notebook .jp-Cell:not(.jp-mod-active) .jp-OutputPrompt {
	  opacity: var(--jp-cell-prompt-not-active-opacity);
	  color: var(--jp-cell-prompt-not-active-font-color);
	}
	
	/* cell is active */
	.jp-Notebook .jp-Cell.jp-mod-active .jp-Collapser {
	  background: var(--jp-brand-color1);
	}
	
	/* collapser is hovered */
	.jp-Notebook .jp-Cell .jp-Collapser:hover {
	  box-shadow: var(--jp-elevation-z2);
	  background: var(--jp-brand-color1);
	  opacity: var(--jp-cell-collapser-not-active-hover-opacity);
	}
	
	/* cell is active and collapser is hovered */
	.jp-Notebook .jp-Cell.jp-mod-active .jp-Collapser:hover {
	  background: var(--jp-brand-color0);
	  opacity: 1;
	}
	
	/* Command mode */
	
	.jp-Notebook.jp-mod-commandMode .jp-Cell.jp-mod-selected {
	  background: var(--jp-notebook-multiselected-color);
	}
	
	.jp-Notebook.jp-mod-commandMode
	  .jp-Cell.jp-mod-active.jp-mod-selected:not(.jp-mod-multiSelected) {
	  background: transparent;
	}
	
	/* Edit mode */
	
	.jp-Notebook.jp-mod-editMode .jp-Cell.jp-mod-active .jp-InputArea-editor {
	  border: var(--jp-border-width) solid var(--jp-cell-editor-active-border-color);
	  box-shadow: var(--jp-input-box-shadow);
	  background-color: var(--jp-cell-editor-active-background);
	}
	
	/*-----------------------------------------------------------------------------
	| Notebook drag and drop
	|----------------------------------------------------------------------------*/
	
	.jp-Notebook-cell.jp-mod-dropSource {
	  opacity: 0.5;
	}
	
	.jp-Notebook-cell.jp-mod-dropTarget,
	.jp-Notebook.jp-mod-commandMode
	  .jp-Notebook-cell.jp-mod-active.jp-mod-selected.jp-mod-dropTarget {
	  border-top-color: var(--jp-private-notebook-selected-color);
	  border-top-style: solid;
	  border-top-width: 2px;
	}
	
	.jp-dragImage {
	  display: flex;
	  flex-direction: row;
	  width: var(--jp-private-notebook-dragImage-width);
	  height: var(--jp-private-notebook-dragImage-height);
	  border: var(--jp-border-width) solid var(--jp-cell-editor-border-color);
	  background: var(--jp-cell-editor-background);
	  overflow: visible;
	}
	
	.jp-dragImage-singlePrompt {
	  box-shadow: 2px 2px 4px 0px rgba(0, 0, 0, 0.12);
	}
	
	.jp-dragImage .jp-dragImage-content {
	  flex: 1 1 auto;
	  z-index: 2;
	  font-size: var(--jp-code-font-size);
	  font-family: var(--jp-code-font-family);
	  line-height: var(--jp-code-line-height);
	  padding: var(--jp-code-padding);
	  border: var(--jp-border-width) solid var(--jp-cell-editor-border-color);
	  background: var(--jp-cell-editor-background-color);
	  color: var(--jp-content-font-color3);
	  text-align: left;
	  margin: 4px 4px 4px 0px;
	}
	
	.jp-dragImage .jp-dragImage-prompt {
	  flex: 0 0 auto;
	  min-width: 36px;
	  color: var(--jp-cell-inprompt-font-color);
	  padding: var(--jp-code-padding);
	  padding-left: 12px;
	  font-family: var(--jp-cell-prompt-font-family);
	  letter-spacing: var(--jp-cell-prompt-letter-spacing);
	  line-height: 1.9;
	  font-size: var(--jp-code-font-size);
	  border: var(--jp-border-width) solid transparent;
	}
	
	.jp-dragImage-multipleBack {
	  z-index: -1;
	  position: absolute;
	  height: 32px;
	  width: 300px;
	  top: 8px;
	  left: 8px;
	  background: var(--jp-layout-color2);
	  border: var(--jp-border-width) solid var(--jp-input-border-color);
	  box-shadow: 2px 2px 4px 0px rgba(0, 0, 0, 0.12);
	}
	
	/*-----------------------------------------------------------------------------
	| Cell toolbar
	|----------------------------------------------------------------------------*/
	
	.jp-NotebookTools {
	  display: block;
	  min-width: var(--jp-sidebar-min-width);
	  color: var(--jp-ui-font-color1);
	  background: var(--jp-layout-color1);
	  /* This is needed so that all font sizing of children done in ems is
		* relative to this base size */
	  font-size: var(--jp-ui-font-size1);
	  overflow: auto;
	}
	
	.jp-NotebookTools-tool {
	  padding: 0px 12px 0 12px;
	}
	
	.jp-ActiveCellTool {
	  padding: 12px;
	  background-color: var(--jp-layout-color1);
	  border-top: none !important;
	}
	
	.jp-ActiveCellTool .jp-InputArea-prompt {
	  flex: 0 0 auto;
	  padding-left: 0px;
	}
	
	.jp-ActiveCellTool .jp-InputArea-editor {
	  flex: 1 1 auto;
	  background: var(--jp-cell-editor-background);
	  border-color: var(--jp-cell-editor-border-color);
	}
	
	.jp-ActiveCellTool .jp-InputArea-editor .CodeMirror {
	  background: transparent;
	}
	
	.jp-MetadataEditorTool {
	  flex-direction: column;
	  padding: 12px 0px 12px 0px;
	}
	
	.jp-RankedPanel > :not(:first-child) {
	  margin-top: 12px;
	}
	
	.jp-KeySelector select.jp-mod-styled {
	  font-size: var(--jp-ui-font-size1);
	  color: var(--jp-ui-font-color0);
	  border: var(--jp-border-width) solid var(--jp-border-color1);
	}
	
	.jp-KeySelector label,
	.jp-MetadataEditorTool label {
	  line-height: 1.4;
	}
	
	.jp-NotebookTools .jp-select-wrapper {
	  margin-top: 4px;
	  margin-bottom: 0px;
	}
	
	.jp-NotebookTools .jp-Collapse {
	  margin-top: 16px;
	}
	
	/*-----------------------------------------------------------------------------
	| Presentation Mode (.jp-mod-presentationMode)
	|----------------------------------------------------------------------------*/
	
	.jp-mod-presentationMode .jp-Notebook {
	  --jp-content-font-size1: var(--jp-content-presentation-font-size1);
	  --jp-code-font-size: var(--jp-code-presentation-font-size);
	}
	
	.jp-mod-presentationMode .jp-Notebook .jp-Cell .jp-InputPrompt,
	.jp-mod-presentationMode .jp-Notebook .jp-Cell .jp-OutputPrompt {
	  flex: 0 0 110px;
	}
	
	</style>
	
		<style type="text/css">
	/*-----------------------------------------------------------------------------
	| Copyright (c) Jupyter Development Team.
	| Distributed under the terms of the Modified BSD License.
	|----------------------------------------------------------------------------*/

	
	:root {
	  /* Elevation
	   *
	   * We style box-shadows using Material Design's idea of elevation. These particular numbers are taken from here:
	   *
	   * https://github.com/material-components/material-components-web
	   * https://material-components-web.appspot.com/elevation.html
	   */
	
	  --jp-shadow-base-lightness: 0;
	  --jp-shadow-umbra-color: rgba(
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		0.2
	  );
	  --jp-shadow-penumbra-color: rgba(
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		0.14
	  );
	  --jp-shadow-ambient-color: rgba(
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		var(--jp-shadow-base-lightness),
		0.12
	  );
	  --jp-elevation-z0: none;
	  --jp-elevation-z1: 0px 2px 1px -1px var(--jp-shadow-umbra-color),
		0px 1px 1px 0px var(--jp-shadow-penumbra-color),
		0px 1px 3px 0px var(--jp-shadow-ambient-color);
	  --jp-elevation-z2: 0px 3px 1px -2px var(--jp-shadow-umbra-color),
		0px 2px 2px 0px var(--jp-shadow-penumbra-color),
		0px 1px 5px 0px var(--jp-shadow-ambient-color);
	  --jp-elevation-z4: 0px 2px 4px -1px var(--jp-shadow-umbra-color),
		0px 4px 5px 0px var(--jp-shadow-penumbra-color),
		0px 1px 10px 0px var(--jp-shadow-ambient-color);
	  --jp-elevation-z6: 0px 3px 5px -1px var(--jp-shadow-umbra-color),
		0px 6px 10px 0px var(--jp-shadow-penumbra-color),
		0px 1px 18px 0px var(--jp-shadow-ambient-color);
	  --jp-elevation-z8: 0px 5px 5px -3px var(--jp-shadow-umbra-color),
		0px 8px 10px 1px var(--jp-shadow-penumbra-color),
		0px 3px 14px 2px var(--jp-shadow-ambient-color);
	  --jp-elevation-z12: 0px 7px 8px -4px var(--jp-shadow-umbra-color),
		0px 12px 17px 2px var(--jp-shadow-penumbra-color),
		0px 5px 22px 4px var(--jp-shadow-ambient-color);
	  --jp-elevation-z16: 0px 8px 10px -5px var(--jp-shadow-umbra-color),
		0px 16px 24px 2px var(--jp-shadow-penumbra-color),
		0px 6px 30px 5px var(--jp-shadow-ambient-color);
	  --jp-elevation-z20: 0px 10px 13px -6px var(--jp-shadow-umbra-color),
		0px 20px 31px 3px var(--jp-shadow-penumbra-color),
		0px 8px 38px 7px var(--jp-shadow-ambient-color);
	  --jp-elevation-z24: 0px 11px 15px -7px var(--jp-shadow-umbra-color),
		0px 24px 38px 3px var(--jp-shadow-penumbra-color),
		0px 9px 46px 8px var(--jp-shadow-ambient-color);
	
	  /* Borders
	   *
	   * The following variables, specify the visual styling of borders in JupyterLab.
	   */
	
	  --jp-border-width: 1px;
	  --jp-border-color0: var(--md-grey-400);
	  --jp-border-color1: var(--md-grey-400);
	  --jp-border-color2: var(--md-grey-300);
	  --jp-border-color3: var(--md-grey-200);
	  --jp-border-radius: 2px;
	
	  /* UI Fonts
	   *
	   * The UI font CSS variables are used for the typography all of the JupyterLab
	   * user interface elements that are not directly user generated content.
	   *
	   * The font sizing here is done assuming that the body font size of --jp-ui-font-size1
	   * is applied to a parent element. When children elements, such as headings, are sized
	   * in em all things will be computed relative to that body size.
	   */
	
	  --jp-ui-font-scale-factor: 1.2;
	  --jp-ui-font-size0: 0.83333em;
	  --jp-ui-font-size1: 13px; /* Base font size */
	  --jp-ui-font-size2: 1.2em;
	  --jp-ui-font-size3: 1.44em;
	
	  --jp-ui-font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Helvetica,
		Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol';
	
	  /*
	   * Use these font colors against the corresponding main layout colors.
	   * In a light theme, these go from dark to light.
	   */
	
	  /* Defaults use Material Design specification */
	  --jp-ui-font-color0: rgba(0, 0, 0, 1);
	  --jp-ui-font-color1: rgba(0, 0, 0, 0.87);
	  --jp-ui-font-color2: rgba(0, 0, 0, 0.54);
	  --jp-ui-font-color3: rgba(0, 0, 0, 0.38);
	
	  /*
	   * Use these against the brand/accent/warn/error colors.
	   * These will typically go from light to darker, in both a dark and light theme.
	   */
	
	  --jp-ui-inverse-font-color0: rgba(255, 255, 255, 1);
	  --jp-ui-inverse-font-color1: rgba(255, 255, 255, 1);
	  --jp-ui-inverse-font-color2: rgba(255, 255, 255, 0.7);
	  --jp-ui-inverse-font-color3: rgba(255, 255, 255, 0.5);
	
	  /* Content Fonts
	   *
	   * Content font variables are used for typography of user generated content.
	   *
	   * The font sizing here is done assuming that the body font size of --jp-content-font-size1
	   * is applied to a parent element. When children elements, such as headings, are sized
	   * in em all things will be computed relative to that body size.
	   */
	
	  --jp-content-line-height: 1.6;
	  --jp-content-font-scale-factor: 1.2;
	  --jp-content-font-size0: 0.83333em;
	  --jp-content-font-size1: 14px; /* Base font size */
	  --jp-content-font-size2: 1.2em;
	  --jp-content-font-size3: 1.44em;
	  --jp-content-font-size4: 1.728em;
	  --jp-content-font-size5: 2.0736em;
	
	  /* This gives a magnification of about 125% in presentation mode over normal. */
	  --jp-content-presentation-font-size1: 17px;
	
	  --jp-content-heading-line-height: 1;
	  --jp-content-heading-margin-top: 1.2em;
	  --jp-content-heading-margin-bottom: 0.8em;
	  --jp-content-heading-font-weight: 500;
	
	  /* Defaults use Material Design specification */
	  --jp-content-font-color0: rgba(0, 0, 0, 1);
	  --jp-content-font-color1: rgba(0, 0, 0, 0.87);
	  --jp-content-font-color2: rgba(0, 0, 0, 0.54);
	  --jp-content-font-color3: rgba(0, 0, 0, 0.38);
	
	  --jp-content-link-color: var(--md-blue-700);
	
	  --jp-content-font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI',
		Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji',
		'Segoe UI Symbol';
	
	  /*
	   * Code Fonts
	   *
	   * Code font variables are used for typography of code and other monospaces content.
	   */
	
	  --jp-code-font-size: 13px;
	  --jp-code-line-height: 1.3077; /* 17px for 13px base */
	  --jp-code-padding: 5px; /* 5px for 13px base, codemirror highlighting needs integer px value */
	  --jp-code-font-family-default: Menlo, Consolas, 'DejaVu Sans Mono', monospace;
	  --jp-code-font-family: var(--jp-code-font-family-default);
	
	  /* This gives a magnification of about 125% in presentation mode over normal. */
	  --jp-code-presentation-font-size: 16px;
	
	  /* may need to tweak cursor width if you change font size */
	  --jp-code-cursor-width0: 1.4px;
	  --jp-code-cursor-width1: 2px;
	  --jp-code-cursor-width2: 4px;
	
	  /* Layout
	   *
	   * The following are the main layout colors use in JupyterLab. In a light
	   * theme these would go from light to dark.
	   */
	
	  --jp-layout-color0: white;
	  --jp-layout-color1: white;
	  --jp-layout-color2: var(--md-grey-200);
	  --jp-layout-color3: var(--md-grey-400);
	  --jp-layout-color4: var(--md-grey-600);
	
	  /* Inverse Layout
	   *
	   * The following are the inverse layout colors use in JupyterLab. In a light
	   * theme these would go from dark to light.
	   */
	
	  --jp-inverse-layout-color0: #111111;
	  --jp-inverse-layout-color1: var(--md-grey-900);
	  --jp-inverse-layout-color2: var(--md-grey-800);
	  --jp-inverse-layout-color3: var(--md-grey-700);
	  --jp-inverse-layout-color4: var(--md-grey-600);
	
	  /* Brand/accent */
	
	  --jp-brand-color0: var(--md-blue-700);
	  --jp-brand-color1: var(--md-blue-500);
	  --jp-brand-color2: var(--md-blue-300);
	  --jp-brand-color3: var(--md-blue-100);
	  --jp-brand-color4: var(--md-blue-50);
	
	  --jp-accent-color0: var(--md-green-700);
	  --jp-accent-color1: var(--md-green-500);
	  --jp-accent-color2: var(--md-green-300);
	  --jp-accent-color3: var(--md-green-100);
	
	  /* State colors (warn, error, success, info) */
	
	  --jp-warn-color0: var(--md-orange-700);
	  --jp-warn-color1: var(--md-orange-500);
	  --jp-warn-color2: var(--md-orange-300);
	  --jp-warn-color3: var(--md-orange-100);
	
	  --jp-error-color0: var(--md-red-700);
	  --jp-error-color1: var(--md-red-500);
	  --jp-error-color2: var(--md-red-300);
	  --jp-error-color3: var(--md-red-100);
	
	  --jp-success-color0: var(--md-green-700);
	  --jp-success-color1: var(--md-green-500);
	  --jp-success-color2: var(--md-green-300);
	  --jp-success-color3: var(--md-green-100);
	
	  --jp-info-color0: var(--md-cyan-700);
	  --jp-info-color1: var(--md-cyan-500);
	  --jp-info-color2: var(--md-cyan-300);
	  --jp-info-color3: var(--md-cyan-100);
	
	  /* Cell specific styles */
	
	  --jp-cell-padding: 5px;
	
	  --jp-cell-collapser-width: 8px;
	  --jp-cell-collapser-min-height: 20px;
	  --jp-cell-collapser-not-active-hover-opacity: 0.6;
	
	  --jp-cell-editor-background: var(--md-grey-100);
	  --jp-cell-editor-border-color: var(--md-grey-300);
	  --jp-cell-editor-box-shadow: inset 0 0 2px var(--md-blue-300);
	  --jp-cell-editor-active-background: var(--jp-layout-color0);
	  --jp-cell-editor-active-border-color: var(--jp-brand-color1);
	
	  --jp-cell-prompt-width: 64px;
	  --jp-cell-prompt-font-family: var(--jp-code-font-family-default);
	  --jp-cell-prompt-letter-spacing: 0px;
	  --jp-cell-prompt-opacity: 1;
	  --jp-cell-prompt-not-active-opacity: 0.5;
	  --jp-cell-prompt-not-active-font-color: var(--md-grey-700);
	  /* A custom blend of MD grey and blue 600
	   * See https://meyerweb.com/eric/tools/color-blend/#546E7A:1E88E5:5:hex */
	  --jp-cell-inprompt-font-color: #307fc1;
	  /* A custom blend of MD grey and orange 600
	   * https://meyerweb.com/eric/tools/color-blend/#546E7A:F4511E:5:hex */
	  --jp-cell-outprompt-font-color: #bf5b3d;
	
	  /* Notebook specific styles */
	
	  --jp-notebook-padding: 10px;
	  --jp-notebook-select-background: var(--jp-layout-color1);
	  --jp-notebook-multiselected-color: var(--md-blue-50);
	
	  /* The scroll padding is calculated to fill enough space at the bottom of the
	  notebook to show one single-line cell (with appropriate padding) at the top
	  when the notebook is scrolled all the way to the bottom. We also subtract one
	  pixel so that no scrollbar appears if we have just one single-line cell in the
	  notebook. This padding is to enable a 'scroll past end' feature in a notebook.
	  */
	  --jp-notebook-scroll-padding: calc(
		100% - var(--jp-code-font-size) * var(--jp-code-line-height) -
		  var(--jp-code-padding) - var(--jp-cell-padding) - 1px
	  );
	
	  /* Rendermime styles */
	
	  --jp-rendermime-error-background: #fdd;
	  --jp-rendermime-table-row-background: var(--md-grey-100);
	  --jp-rendermime-table-row-hover-background: var(--md-light-blue-50);
	
	  /* Dialog specific styles */
	
	  --jp-dialog-background: rgba(0, 0, 0, 0.25);
	
	  /* Console specific styles */
	
	  --jp-console-padding: 10px;
	
	  /* Toolbar specific styles */
	
	  --jp-toolbar-border-color: var(--jp-border-color1);
	  --jp-toolbar-micro-height: 8px;
	  --jp-toolbar-background: var(--jp-layout-color1);
	  --jp-toolbar-box-shadow: 0px 0px 2px 0px rgba(0, 0, 0, 0.24);
	  --jp-toolbar-header-margin: 4px 4px 0px 4px;
	  --jp-toolbar-active-background: var(--md-grey-300);
	
	  /* Input field styles */
	
	  --jp-input-box-shadow: inset 0 0 2px var(--md-blue-300);
	  --jp-input-active-background: var(--jp-layout-color1);
	  --jp-input-hover-background: var(--jp-layout-color1);
	  --jp-input-background: var(--md-grey-100);
	  --jp-input-border-color: var(--jp-border-color1);
	  --jp-input-active-border-color: var(--jp-brand-color1);
	  --jp-input-active-box-shadow-color: rgba(19, 124, 189, 0.3);
	
	  /* General editor styles */
	
	  --jp-editor-selected-background: #d9d9d9;
	  --jp-editor-selected-focused-background: #d7d4f0;
	  --jp-editor-cursor-color: var(--jp-ui-font-color0);
	
	  /* Code mirror specific styles */
	
	  --jp-mirror-editor-keyword-color: #008000;
	  --jp-mirror-editor-atom-color: #88f;
	  --jp-mirror-editor-number-color: #080;
	  --jp-mirror-editor-def-color: #00f;
	  --jp-mirror-editor-variable-color: var(--md-grey-900);
	  --jp-mirror-editor-variable-2-color: #05a;
	  --jp-mirror-editor-variable-3-color: #085;
	  --jp-mirror-editor-punctuation-color: #05a;
	  --jp-mirror-editor-property-color: #05a;
	  --jp-mirror-editor-operator-color: #aa22ff;
	  --jp-mirror-editor-comment-color: #408080;
	  --jp-mirror-editor-string-color: #ba2121;
	  --jp-mirror-editor-string-2-color: #708;
	  --jp-mirror-editor-meta-color: #aa22ff;
	  --jp-mirror-editor-qualifier-color: #555;
	  --jp-mirror-editor-builtin-color: #008000;
	  --jp-mirror-editor-bracket-color: #997;
	  --jp-mirror-editor-tag-color: #170;
	  --jp-mirror-editor-attribute-color: #00c;
	  --jp-mirror-editor-header-color: blue;
	  --jp-mirror-editor-quote-color: #090;
	  --jp-mirror-editor-link-color: #00c;
	  --jp-mirror-editor-error-color: #f00;
	  --jp-mirror-editor-hr-color: #999;
	
	  /* Vega extension styles */
	
	  --jp-vega-background: white;
	
	  /* Sidebar-related styles */
	
	  --jp-sidebar-min-width: 250px;
	
	  /* Search-related styles */
	
	  --jp-search-toggle-off-opacity: 0.5;
	  --jp-search-toggle-hover-opacity: 0.8;
	  --jp-search-toggle-on-opacity: 1;
	  --jp-search-selected-match-background-color: rgb(245, 200, 0);
	  --jp-search-selected-match-color: black;
	  --jp-search-unselected-match-background-color: var(
		--jp-inverse-layout-color0
	  );
	  --jp-search-unselected-match-color: var(--jp-ui-inverse-font-color0);
	
	  /* Icon colors that work well with light or dark backgrounds */
	  --jp-icon-contrast-color0: var(--md-purple-600);
	  --jp-icon-contrast-color1: var(--md-green-600);
	  --jp-icon-contrast-color2: var(--md-pink-600);
	  --jp-icon-contrast-color3: var(--md-blue-600);
	}
	</style>
	
	<style type="text/css">
	a.anchor-link {
	   display: none;
	}
	.highlight  {
		margin: 0.4em;
	}
	
	/* Input area styling */
	.jp-InputArea {
		overflow: hidden;
	}
	
	.jp-InputArea-editor {
		overflow: hidden;
	}
	
	@media print {
	  body {
		margin: 0;
	  }
	}
	</style>
	
	<!-- Load mathjax -->
		<script src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.7/latest.js?config=TeX-MML-AM_CHTML-full,Safe"> </script>
		<!-- MathJax configuration -->
		<script type="text/x-mathjax-config">
		init_mathjax = function() {
			if (window.MathJax) {
			// MathJax loaded
				MathJax.Hub.Config({
					TeX: {
						equationNumbers: {
						autoNumber: "AMS",
						useLabelIds: true
						}
					},
					tex2jax: {
						inlineMath: [ ['$','$'], ["\\(","\\)"] ],
						displayMath: [ ['$$','$$'], ["\\[","\\]"] ],
						processEscapes: true,
						processEnvironments: true
					},
					displayAlign: 'center',
					CommonHTML: {
						linebreaks: { 
						automatic: true 
						}
					},
					"HTML-CSS": {
						linebreaks: { 
						automatic: true 
						}
					}
				});
			
				MathJax.Hub.Queue(["Typeset", MathJax.Hub]);
			}
		}
		init_mathjax();
		</script>
		<!-- End of mathjax configuration --></head>
	<body class="jp-Notebook" data-jp-theme-light="true" data-jp-theme-name="JupyterLab Light">
	
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<h1 id="XKCD-plots-in-Matplotlib">XKCD plots in Matplotlib<a class="anchor-link" href="#XKCD-plots-in-Matplotlib">&#182;</a></h1>
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>This notebook originally appeared as a blog post at <a href="http://jakevdp.github.com/blog/2012/10/07/xkcd-style-plots-in-matplotlib/">Pythonic Perambulations</a> by Jake Vanderplas.</p>
	
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p><!-- PELICAN_BEGIN_SUMMARY -->
	<em>Update: the matplotlib pull request has been merged!  See</em>
	<a href="http://jakevdp.github.io/blog/2013/07/10/XKCD-plots-in-matplotlib/"><em>This post</em></a>
	<em>for a description of the XKCD functionality now built-in to matplotlib!</em></p>
	<p>One of the problems I've had with typical matplotlib figures is that everything in them is so precise, so perfect.  For an example of what I mean, take a look at this figure:</p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[1]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="kn">from</span> <span class="nn">IPython.display</span> <span class="kn">import</span> <span class="n">Image</span>
	<span class="n">Image</span><span class="p">(</span><span class="s1">&#39;http://jakevdp.github.com/figures/xkcd_version.png&#39;</span><span class="p">)</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt">Out[1]:</div>
	
	
	
	
	<div class="jp-RenderedImage jp-OutputArea-output jp-OutputArea-executeResult">
	<img src="data:image/png;base64,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
	RU5ErkJggg==
	"
	>
	</div>
	
	</div>
	
	</div>
	
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>Sometimes when showing schematic plots, this is the type of figure I want to display.  But drawing it by hand is a pain: I'd rather just use matplotlib.  The problem is, matplotlib is a bit too precise.  Attempting to duplicate this figure in matplotlib leads to something like this:
	<!-- PELICAN_END_SUMMARY --></p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[2]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="n">Image</span><span class="p">(</span><span class="s1">&#39;http://jakevdp.github.com/figures/mpl_version.png&#39;</span><span class="p">)</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt">Out[2]:</div>
	
	
	
	
	<div class="jp-RenderedImage jp-OutputArea-output jp-OutputArea-executeResult">
	<img src="data:image/png;base64,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
	SUVORK5CYII=
	"
	>
	</div>
	
	</div>
	
	</div>
	
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>It just doesn't have the same effect.  Matplotlib is great for scientific plots, but sometimes you don't want to be so precise.</p>
	<p>This subject has recently come up on the matplotlib mailing list, and started some interesting discussions.
	As near as I can tell, this started with a thread on a
	<a href="http://mathematica.stackexchange.com/questions/11350/xkcd-style-graphs">mathematica list</a>
	which prompted a thread on the <a href="http://matplotlib.1069221.n5.nabble.com/XKCD-style-graphs-td39226.html">matplotlib list</a>
	wondering if the same could be done in matplotlib.</p>
	<p>Damon McDougall offered a quick
	<a href="http://www.mail-archive.com/matplotlib-users@lists.sourceforge.net/msg25499.html">solution</a>
	which was improved by Fernando Perez in <a href="http://nbviewer.ipython.org/3835181/">this notebook</a>, and
	within a few days there was a <a href="https://github.com/matplotlib/matplotlib/pull/1329">matplotlib pull request</a> offering a very general
	way to create sketch-style plots in matplotlib.  Only a few days from a cool idea to a
	working implementation: this is one of the most incredible aspects of package development on github.</p>
	<p>The pull request looks really nice, but will likely not be included in a released version of
	matplotlib until at least version 1.3.  In the mean-time, I wanted a way to play around with
	these types of plots in a way that is compatible with the current release of matplotlib.  To do that,
	I created the following code:</p>
	
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<h2 id="The-Code:-XKCDify">The Code: XKCDify<a class="anchor-link" href="#The-Code:-XKCDify">&#182;</a></h2>
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>XKCDify will take a matplotlib <code>Axes</code> instance, and modify the plot elements in-place to make
	them look hand-drawn.
	First off, we'll need to make sure we have the Humor Sans font.
	It can be downloaded using the command below.</p>
	<p>Next we'll create a function <code>xkcd_line</code> to add jitter to lines.  We want this to be very general, so
	we'll normalize the size of the lines, and use a low-pass filter to add correlated noise, perpendicular
	to the direction of the line.  There are a few parameters for this filter that can be tweaked to
	customize the appearance of the jitter.</p>
	<p>Finally, we'll create a function which accepts a matplotlib axis, and calls <code>xkcd_line</code> on
	all lines in the axis.  Additionally, we'll switch the font of all text in the axes, and add
	some background lines for a nice effect where lines cross.  We'll also draw axes, and move the
	axes labels and titles to the appropriate location.</p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell jp-mod-noOutputs  ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[3]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="sd">&quot;&quot;&quot;</span>
	<span class="sd">XKCD plot generator</span>
	<span class="sd">-------------------</span>
	<span class="sd">Author: Jake Vanderplas</span>
	
	<span class="sd">This is a script that will take any matplotlib line diagram, and convert it</span>
	<span class="sd">to an XKCD-style plot.  It will work for plots with line &amp; text elements,</span>
	<span class="sd">including axes labels and titles (but not axes tick labels).</span>
	
	<span class="sd">The idea for this comes from work by Damon McDougall</span>
	<span class="sd">  http://www.mail-archive.com/matplotlib-users@lists.sourceforge.net/msg25499.html</span>
	<span class="sd">&quot;&quot;&quot;</span>
	<span class="kn">import</span> <span class="nn">numpy</span> <span class="k">as</span> <span class="nn">np</span>
	<span class="kn">import</span> <span class="nn">pylab</span> <span class="k">as</span> <span class="nn">pl</span>
	<span class="kn">from</span> <span class="nn">scipy</span> <span class="kn">import</span> <span class="n">interpolate</span><span class="p">,</span> <span class="n">signal</span>
	<span class="kn">import</span> <span class="nn">matplotlib.font_manager</span> <span class="k">as</span> <span class="nn">fm</span>
	
	
	<span class="c1"># We need a special font for the code below.  It can be downloaded this way:</span>
	<span class="kn">import</span> <span class="nn">os</span>
	<span class="kn">import</span> <span class="nn">urllib2</span>
	<span class="k">if</span> <span class="ow">not</span> <span class="n">os</span><span class="o">.</span><span class="n">path</span><span class="o">.</span><span class="n">exists</span><span class="p">(</span><span class="s1">&#39;Humor-Sans.ttf&#39;</span><span class="p">):</span>
		<span class="n">fhandle</span> <span class="o">=</span> <span class="n">urllib2</span><span class="o">.</span><span class="n">urlopen</span><span class="p">(</span><span class="s1">&#39;http://antiyawn.com/uploads/Humor-Sans-1.0.ttf&#39;</span><span class="p">)</span>
		<span class="nb">open</span><span class="p">(</span><span class="s1">&#39;Humor-Sans.ttf&#39;</span><span class="p">,</span> <span class="s1">&#39;wb&#39;</span><span class="p">)</span><span class="o">.</span><span class="n">write</span><span class="p">(</span><span class="n">fhandle</span><span class="o">.</span><span class="n">read</span><span class="p">())</span>
	
		
	<span class="k">def</span> <span class="nf">xkcd_line</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y</span><span class="p">,</span> <span class="n">xlim</span><span class="o">=</span><span class="kc">None</span><span class="p">,</span> <span class="n">ylim</span><span class="o">=</span><span class="kc">None</span><span class="p">,</span>
				  <span class="n">mag</span><span class="o">=</span><span class="mf">1.0</span><span class="p">,</span> <span class="n">f1</span><span class="o">=</span><span class="mi">30</span><span class="p">,</span> <span class="n">f2</span><span class="o">=</span><span class="mf">0.05</span><span class="p">,</span> <span class="n">f3</span><span class="o">=</span><span class="mi">15</span><span class="p">):</span>
		<span class="sd">&quot;&quot;&quot;</span>
	<span class="sd">    Mimic a hand-drawn line from (x, y) data</span>
	
	<span class="sd">    Parameters</span>
	<span class="sd">    ----------</span>
	<span class="sd">    x, y : array_like</span>
	<span class="sd">        arrays to be modified</span>
	<span class="sd">    xlim, ylim : data range</span>
	<span class="sd">        the assumed plot range for the modification.  If not specified,</span>
	<span class="sd">        they will be guessed from the  data</span>
	<span class="sd">    mag : float</span>
	<span class="sd">        magnitude of distortions</span>
	<span class="sd">    f1, f2, f3 : int, float, int</span>
	<span class="sd">        filtering parameters.  f1 gives the size of the window, f2 gives</span>
	<span class="sd">        the high-frequency cutoff, f3 gives the size of the filter</span>
	<span class="sd">    </span>
	<span class="sd">    Returns</span>
	<span class="sd">    -------</span>
	<span class="sd">    x, y : ndarrays</span>
	<span class="sd">        The modified lines</span>
	<span class="sd">    &quot;&quot;&quot;</span>
		<span class="n">x</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">asarray</span><span class="p">(</span><span class="n">x</span><span class="p">)</span>
		<span class="n">y</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">asarray</span><span class="p">(</span><span class="n">y</span><span class="p">)</span>
		
		<span class="c1"># get limits for rescaling</span>
		<span class="k">if</span> <span class="n">xlim</span> <span class="ow">is</span> <span class="kc">None</span><span class="p">:</span>
			<span class="n">xlim</span> <span class="o">=</span> <span class="p">(</span><span class="n">x</span><span class="o">.</span><span class="n">min</span><span class="p">(),</span> <span class="n">x</span><span class="o">.</span><span class="n">max</span><span class="p">())</span>
		<span class="k">if</span> <span class="n">ylim</span> <span class="ow">is</span> <span class="kc">None</span><span class="p">:</span>
			<span class="n">ylim</span> <span class="o">=</span> <span class="p">(</span><span class="n">y</span><span class="o">.</span><span class="n">min</span><span class="p">(),</span> <span class="n">y</span><span class="o">.</span><span class="n">max</span><span class="p">())</span>
	
		<span class="k">if</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">==</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]:</span>
			<span class="n">xlim</span> <span class="o">=</span> <span class="n">ylim</span>
			
		<span class="k">if</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">==</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">]:</span>
			<span class="n">ylim</span> <span class="o">=</span> <span class="n">xlim</span>
	
		<span class="c1"># scale the data</span>
		<span class="n">x_scaled</span> <span class="o">=</span> <span class="p">(</span><span class="n">x</span> <span class="o">-</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span> <span class="o">*</span> <span class="mf">1.</span> <span class="o">/</span> <span class="p">(</span><span class="n">xlim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span>
		<span class="n">y_scaled</span> <span class="o">=</span> <span class="p">(</span><span class="n">y</span> <span class="o">-</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span> <span class="o">*</span> <span class="mf">1.</span> <span class="o">/</span> <span class="p">(</span><span class="n">ylim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span>
	
		<span class="c1"># compute the total distance along the path</span>
		<span class="n">dx</span> <span class="o">=</span> <span class="n">x_scaled</span><span class="p">[</span><span class="mi">1</span><span class="p">:]</span> <span class="o">-</span> <span class="n">x_scaled</span><span class="p">[:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span>
		<span class="n">dy</span> <span class="o">=</span> <span class="n">y_scaled</span><span class="p">[</span><span class="mi">1</span><span class="p">:]</span> <span class="o">-</span> <span class="n">y_scaled</span><span class="p">[:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span>
		<span class="n">dist_tot</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">sum</span><span class="p">(</span><span class="n">np</span><span class="o">.</span><span class="n">sqrt</span><span class="p">(</span><span class="n">dx</span> <span class="o">*</span> <span class="n">dx</span> <span class="o">+</span> <span class="n">dy</span> <span class="o">*</span> <span class="n">dy</span><span class="p">))</span>
	
		<span class="c1"># number of interpolated points is proportional to the distance</span>
		<span class="n">Nu</span> <span class="o">=</span> <span class="nb">int</span><span class="p">(</span><span class="mi">200</span> <span class="o">*</span> <span class="n">dist_tot</span><span class="p">)</span>
		<span class="n">u</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">arange</span><span class="p">(</span><span class="o">-</span><span class="mi">1</span><span class="p">,</span> <span class="n">Nu</span> <span class="o">+</span> <span class="mi">1</span><span class="p">)</span> <span class="o">*</span> <span class="mf">1.</span> <span class="o">/</span> <span class="p">(</span><span class="n">Nu</span> <span class="o">-</span> <span class="mi">1</span><span class="p">)</span>
	
		<span class="c1"># interpolate curve at sampled points</span>
		<span class="n">k</span> <span class="o">=</span> <span class="nb">min</span><span class="p">(</span><span class="mi">3</span><span class="p">,</span> <span class="nb">len</span><span class="p">(</span><span class="n">x</span><span class="p">)</span> <span class="o">-</span> <span class="mi">1</span><span class="p">)</span>
		<span class="n">res</span> <span class="o">=</span> <span class="n">interpolate</span><span class="o">.</span><span class="n">splprep</span><span class="p">([</span><span class="n">x_scaled</span><span class="p">,</span> <span class="n">y_scaled</span><span class="p">],</span> <span class="n">s</span><span class="o">=</span><span class="mi">0</span><span class="p">,</span> <span class="n">k</span><span class="o">=</span><span class="n">k</span><span class="p">)</span>
		<span class="n">x_int</span><span class="p">,</span> <span class="n">y_int</span> <span class="o">=</span> <span class="n">interpolate</span><span class="o">.</span><span class="n">splev</span><span class="p">(</span><span class="n">u</span><span class="p">,</span> <span class="n">res</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span> 
	
		<span class="c1"># we&#39;ll perturb perpendicular to the drawn line</span>
		<span class="n">dx</span> <span class="o">=</span> <span class="n">x_int</span><span class="p">[</span><span class="mi">2</span><span class="p">:]</span> <span class="o">-</span> <span class="n">x_int</span><span class="p">[:</span><span class="o">-</span><span class="mi">2</span><span class="p">]</span>
		<span class="n">dy</span> <span class="o">=</span> <span class="n">y_int</span><span class="p">[</span><span class="mi">2</span><span class="p">:]</span> <span class="o">-</span> <span class="n">y_int</span><span class="p">[:</span><span class="o">-</span><span class="mi">2</span><span class="p">]</span>
		<span class="n">dist</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">sqrt</span><span class="p">(</span><span class="n">dx</span> <span class="o">*</span> <span class="n">dx</span> <span class="o">+</span> <span class="n">dy</span> <span class="o">*</span> <span class="n">dy</span><span class="p">)</span>
	
		<span class="c1"># create a filtered perturbation</span>
		<span class="n">coeffs</span> <span class="o">=</span> <span class="n">mag</span> <span class="o">*</span> <span class="n">np</span><span class="o">.</span><span class="n">random</span><span class="o">.</span><span class="n">normal</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mf">0.01</span><span class="p">,</span> <span class="nb">len</span><span class="p">(</span><span class="n">x_int</span><span class="p">)</span> <span class="o">-</span> <span class="mi">2</span><span class="p">)</span>
		<span class="n">b</span> <span class="o">=</span> <span class="n">signal</span><span class="o">.</span><span class="n">firwin</span><span class="p">(</span><span class="n">f1</span><span class="p">,</span> <span class="n">f2</span> <span class="o">*</span> <span class="n">dist_tot</span><span class="p">,</span> <span class="n">window</span><span class="o">=</span><span class="p">(</span><span class="s1">&#39;kaiser&#39;</span><span class="p">,</span> <span class="n">f3</span><span class="p">))</span>
		<span class="n">response</span> <span class="o">=</span> <span class="n">signal</span><span class="o">.</span><span class="n">lfilter</span><span class="p">(</span><span class="n">b</span><span class="p">,</span> <span class="mi">1</span><span class="p">,</span> <span class="n">coeffs</span><span class="p">)</span>
	
		<span class="n">x_int</span><span class="p">[</span><span class="mi">1</span><span class="p">:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+=</span> <span class="n">response</span> <span class="o">*</span> <span class="n">dy</span> <span class="o">/</span> <span class="n">dist</span>
		<span class="n">y_int</span><span class="p">[</span><span class="mi">1</span><span class="p">:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+=</span> <span class="n">response</span> <span class="o">*</span> <span class="n">dx</span> <span class="o">/</span> <span class="n">dist</span>
	
		<span class="c1"># un-scale data</span>
		<span class="n">x_int</span> <span class="o">=</span> <span class="n">x_int</span><span class="p">[</span><span class="mi">1</span><span class="p">:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">*</span> <span class="p">(</span><span class="n">xlim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span> <span class="o">+</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
		<span class="n">y_int</span> <span class="o">=</span> <span class="n">y_int</span><span class="p">[</span><span class="mi">1</span><span class="p">:</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">*</span> <span class="p">(</span><span class="n">ylim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span> <span class="o">+</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
		
		<span class="k">return</span> <span class="n">x_int</span><span class="p">,</span> <span class="n">y_int</span>
	
	
	<span class="k">def</span> <span class="nf">XKCDify</span><span class="p">(</span><span class="n">ax</span><span class="p">,</span> <span class="n">mag</span><span class="o">=</span><span class="mf">1.0</span><span class="p">,</span>
				<span class="n">f1</span><span class="o">=</span><span class="mi">50</span><span class="p">,</span> <span class="n">f2</span><span class="o">=</span><span class="mf">0.01</span><span class="p">,</span> <span class="n">f3</span><span class="o">=</span><span class="mi">15</span><span class="p">,</span>
				<span class="n">bgcolor</span><span class="o">=</span><span class="s1">&#39;w&#39;</span><span class="p">,</span>
				<span class="n">xaxis_loc</span><span class="o">=</span><span class="kc">None</span><span class="p">,</span>
				<span class="n">yaxis_loc</span><span class="o">=</span><span class="kc">None</span><span class="p">,</span>
				<span class="n">xaxis_arrow</span><span class="o">=</span><span class="s1">&#39;+&#39;</span><span class="p">,</span>
				<span class="n">yaxis_arrow</span><span class="o">=</span><span class="s1">&#39;+&#39;</span><span class="p">,</span>
				<span class="n">ax_extend</span><span class="o">=</span><span class="mf">0.1</span><span class="p">,</span>
				<span class="n">expand_axes</span><span class="o">=</span><span class="kc">False</span><span class="p">):</span>
		<span class="sd">&quot;&quot;&quot;Make axis look hand-drawn</span>
	
	<span class="sd">    This adjusts all lines, text, legends, and axes in the figure to look</span>
	<span class="sd">    like xkcd plots.  Other plot elements are not modified.</span>
	<span class="sd">    </span>
	<span class="sd">    Parameters</span>
	<span class="sd">    ----------</span>
	<span class="sd">    ax : Axes instance</span>
	<span class="sd">        the axes to be modified.</span>
	<span class="sd">    mag : float</span>
	<span class="sd">        the magnitude of the distortion</span>
	<span class="sd">    f1, f2, f3 : int, float, int</span>
	<span class="sd">        filtering parameters.  f1 gives the size of the window, f2 gives</span>
	<span class="sd">        the high-frequency cutoff, f3 gives the size of the filter</span>
	<span class="sd">    xaxis_loc, yaxis_log : float</span>
	<span class="sd">        The locations to draw the x and y axes.  If not specified, they</span>
	<span class="sd">        will be drawn from the bottom left of the plot</span>
	<span class="sd">    xaxis_arrow, yaxis_arrow : str</span>
	<span class="sd">        where to draw arrows on the x/y axes.  Options are &#39;+&#39;, &#39;-&#39;, &#39;+-&#39;, or &#39;&#39;</span>
	<span class="sd">    ax_extend : float</span>
	<span class="sd">        How far (fractionally) to extend the drawn axes beyond the original</span>
	<span class="sd">        axes limits</span>
	<span class="sd">    expand_axes : bool</span>
	<span class="sd">        if True, then expand axes to fill the figure (useful if there is only</span>
	<span class="sd">        a single axes in the figure)</span>
	<span class="sd">    &quot;&quot;&quot;</span>
		<span class="c1"># Get axes aspect</span>
		<span class="n">ext</span> <span class="o">=</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_window_extent</span><span class="p">()</span><span class="o">.</span><span class="n">extents</span>
		<span class="n">aspect</span> <span class="o">=</span> <span class="p">(</span><span class="n">ext</span><span class="p">[</span><span class="mi">3</span><span class="p">]</span> <span class="o">-</span> <span class="n">ext</span><span class="p">[</span><span class="mi">1</span><span class="p">])</span> <span class="o">/</span> <span class="p">(</span><span class="n">ext</span><span class="p">[</span><span class="mi">2</span><span class="p">]</span> <span class="o">-</span> <span class="n">ext</span><span class="p">[</span><span class="mi">0</span><span class="p">])</span>
	
		<span class="n">xlim</span> <span class="o">=</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_xlim</span><span class="p">()</span>
		<span class="n">ylim</span> <span class="o">=</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_ylim</span><span class="p">()</span>
	
		<span class="n">xspan</span> <span class="o">=</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
		<span class="n">yspan</span> <span class="o">=</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">-</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
	
		<span class="n">xax_lim</span> <span class="o">=</span> <span class="p">(</span><span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">ax_extend</span> <span class="o">*</span> <span class="n">xspan</span><span class="p">,</span>
				   <span class="n">xlim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">ax_extend</span> <span class="o">*</span> <span class="n">xspan</span><span class="p">)</span>
		<span class="n">yax_lim</span> <span class="o">=</span> <span class="p">(</span><span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">ax_extend</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
				   <span class="n">ylim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">ax_extend</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">)</span>
	
		<span class="k">if</span> <span class="n">xaxis_loc</span> <span class="ow">is</span> <span class="kc">None</span><span class="p">:</span>
			<span class="n">xaxis_loc</span> <span class="o">=</span> <span class="n">ylim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
	
		<span class="k">if</span> <span class="n">yaxis_loc</span> <span class="ow">is</span> <span class="kc">None</span><span class="p">:</span>
			<span class="n">yaxis_loc</span> <span class="o">=</span> <span class="n">xlim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
	
		<span class="c1"># Draw axes</span>
		<span class="n">xaxis</span> <span class="o">=</span> <span class="n">pl</span><span class="o">.</span><span class="n">Line2D</span><span class="p">([</span><span class="n">xax_lim</span><span class="p">[</span><span class="mi">0</span><span class="p">],</span> <span class="n">xax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">]],</span> <span class="p">[</span><span class="n">xaxis_loc</span><span class="p">,</span> <span class="n">xaxis_loc</span><span class="p">],</span>
						  <span class="n">linestyle</span><span class="o">=</span><span class="s1">&#39;-&#39;</span><span class="p">,</span> <span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">)</span>
		<span class="n">yaxis</span> <span class="o">=</span> <span class="n">pl</span><span class="o">.</span><span class="n">Line2D</span><span class="p">([</span><span class="n">yaxis_loc</span><span class="p">,</span> <span class="n">yaxis_loc</span><span class="p">],</span> <span class="p">[</span><span class="n">yax_lim</span><span class="p">[</span><span class="mi">0</span><span class="p">],</span> <span class="n">yax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">]],</span>
						  <span class="n">linestyle</span><span class="o">=</span><span class="s1">&#39;-&#39;</span><span class="p">,</span> <span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">)</span>
	
		<span class="c1"># Label axes3, 0.5, &#39;hello&#39;, fontsize=14)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="n">xax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">],</span> <span class="n">xaxis_loc</span> <span class="o">-</span> <span class="mf">0.02</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_xlabel</span><span class="p">(),</span>
				<span class="n">fontsize</span><span class="o">=</span><span class="mi">14</span><span class="p">,</span> <span class="n">ha</span><span class="o">=</span><span class="s1">&#39;right&#39;</span><span class="p">,</span> <span class="n">va</span><span class="o">=</span><span class="s1">&#39;top&#39;</span><span class="p">,</span> <span class="n">rotation</span><span class="o">=</span><span class="mi">12</span><span class="p">)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="n">yaxis_loc</span> <span class="o">-</span> <span class="mf">0.02</span> <span class="o">*</span> <span class="n">xspan</span><span class="p">,</span> <span class="n">yax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">],</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_ylabel</span><span class="p">(),</span>
				<span class="n">fontsize</span><span class="o">=</span><span class="mi">14</span><span class="p">,</span> <span class="n">ha</span><span class="o">=</span><span class="s1">&#39;right&#39;</span><span class="p">,</span> <span class="n">va</span><span class="o">=</span><span class="s1">&#39;top&#39;</span><span class="p">,</span> <span class="n">rotation</span><span class="o">=</span><span class="mi">78</span><span class="p">)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_xlabel</span><span class="p">(</span><span class="s1">&#39;&#39;</span><span class="p">)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_ylabel</span><span class="p">(</span><span class="s1">&#39;&#39;</span><span class="p">)</span>
	
		<span class="c1"># Add title</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.5</span> <span class="o">*</span> <span class="p">(</span><span class="n">xax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">xax_lim</span><span class="p">[</span><span class="mi">0</span><span class="p">]),</span> <span class="n">yax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">],</span>
				<span class="n">ax</span><span class="o">.</span><span class="n">get_title</span><span class="p">(),</span>
				<span class="n">ha</span><span class="o">=</span><span class="s1">&#39;center&#39;</span><span class="p">,</span> <span class="n">va</span><span class="o">=</span><span class="s1">&#39;bottom&#39;</span><span class="p">,</span> <span class="n">fontsize</span><span class="o">=</span><span class="mi">16</span><span class="p">)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_title</span><span class="p">(</span><span class="s1">&#39;&#39;</span><span class="p">)</span>
	
		<span class="n">Nlines</span> <span class="o">=</span> <span class="nb">len</span><span class="p">(</span><span class="n">ax</span><span class="o">.</span><span class="n">lines</span><span class="p">)</span>
		<span class="n">lines</span> <span class="o">=</span> <span class="p">[</span><span class="n">xaxis</span><span class="p">,</span> <span class="n">yaxis</span><span class="p">]</span> <span class="o">+</span> <span class="p">[</span><span class="n">ax</span><span class="o">.</span><span class="n">lines</span><span class="o">.</span><span class="n">pop</span><span class="p">(</span><span class="mi">0</span><span class="p">)</span> <span class="k">for</span> <span class="n">i</span> <span class="ow">in</span> <span class="nb">range</span><span class="p">(</span><span class="n">Nlines</span><span class="p">)]</span>
	
		<span class="k">for</span> <span class="n">line</span> <span class="ow">in</span> <span class="n">lines</span><span class="p">:</span>
			<span class="n">x</span><span class="p">,</span> <span class="n">y</span> <span class="o">=</span> <span class="n">line</span><span class="o">.</span><span class="n">get_data</span><span class="p">()</span>
	
			<span class="n">x_int</span><span class="p">,</span> <span class="n">y_int</span> <span class="o">=</span> <span class="n">xkcd_line</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y</span><span class="p">,</span> <span class="n">xlim</span><span class="p">,</span> <span class="n">ylim</span><span class="p">,</span>
									 <span class="n">mag</span><span class="p">,</span> <span class="n">f1</span><span class="p">,</span> <span class="n">f2</span><span class="p">,</span> <span class="n">f3</span><span class="p">)</span>
	
			<span class="c1"># create foreground and background line</span>
			<span class="n">lw</span> <span class="o">=</span> <span class="n">line</span><span class="o">.</span><span class="n">get_linewidth</span><span class="p">()</span>
			<span class="n">line</span><span class="o">.</span><span class="n">set_linewidth</span><span class="p">(</span><span class="mi">2</span> <span class="o">*</span> <span class="n">lw</span><span class="p">)</span>
			<span class="n">line</span><span class="o">.</span><span class="n">set_data</span><span class="p">(</span><span class="n">x_int</span><span class="p">,</span> <span class="n">y_int</span><span class="p">)</span>
	
			<span class="c1"># don&#39;t add background line for axes</span>
			<span class="k">if</span> <span class="p">(</span><span class="n">line</span> <span class="ow">is</span> <span class="ow">not</span> <span class="n">xaxis</span><span class="p">)</span> <span class="ow">and</span> <span class="p">(</span><span class="n">line</span> <span class="ow">is</span> <span class="ow">not</span> <span class="n">yaxis</span><span class="p">):</span>
				<span class="n">line_bg</span> <span class="o">=</span> <span class="n">pl</span><span class="o">.</span><span class="n">Line2D</span><span class="p">(</span><span class="n">x_int</span><span class="p">,</span> <span class="n">y_int</span><span class="p">,</span> <span class="n">color</span><span class="o">=</span><span class="n">bgcolor</span><span class="p">,</span>
									<span class="n">linewidth</span><span class="o">=</span><span class="mi">8</span> <span class="o">*</span> <span class="n">lw</span><span class="p">)</span>
	
				<span class="n">ax</span><span class="o">.</span><span class="n">add_line</span><span class="p">(</span><span class="n">line_bg</span><span class="p">)</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">add_line</span><span class="p">(</span><span class="n">line</span><span class="p">)</span>
	
		<span class="c1"># Draw arrow-heads at the end of axes lines</span>
		<span class="n">arr1</span> <span class="o">=</span> <span class="mf">0.03</span> <span class="o">*</span> <span class="n">np</span><span class="o">.</span><span class="n">array</span><span class="p">([</span><span class="o">-</span><span class="mi">1</span><span class="p">,</span> <span class="mi">0</span><span class="p">,</span> <span class="o">-</span><span class="mi">1</span><span class="p">])</span>
		<span class="n">arr2</span> <span class="o">=</span> <span class="mf">0.02</span> <span class="o">*</span> <span class="n">np</span><span class="o">.</span><span class="n">array</span><span class="p">([</span><span class="o">-</span><span class="mi">1</span><span class="p">,</span> <span class="mi">0</span><span class="p">,</span> <span class="mi">1</span><span class="p">])</span>
	
		<span class="n">arr1</span><span class="p">[::</span><span class="mi">2</span><span class="p">]</span> <span class="o">+=</span> <span class="n">np</span><span class="o">.</span><span class="n">random</span><span class="o">.</span><span class="n">normal</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mf">0.005</span><span class="p">,</span> <span class="mi">2</span><span class="p">)</span>
		<span class="n">arr2</span><span class="p">[::</span><span class="mi">2</span><span class="p">]</span> <span class="o">+=</span> <span class="n">np</span><span class="o">.</span><span class="n">random</span><span class="o">.</span><span class="n">normal</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mf">0.005</span><span class="p">,</span> <span class="mi">2</span><span class="p">)</span>
	
		<span class="n">x</span><span class="p">,</span> <span class="n">y</span> <span class="o">=</span> <span class="n">xaxis</span><span class="o">.</span><span class="n">get_data</span><span class="p">()</span>
		<span class="k">if</span> <span class="s1">&#39;+&#39;</span> <span class="ow">in</span> <span class="nb">str</span><span class="p">(</span><span class="n">xaxis_arrow</span><span class="p">):</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">[</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">arr1</span> <span class="o">*</span> <span class="n">xspan</span> <span class="o">*</span> <span class="n">aspect</span><span class="p">,</span>
					<span class="n">y</span><span class="p">[</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">arr2</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
					<span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">2</span><span class="p">)</span>
		<span class="k">if</span> <span class="s1">&#39;-&#39;</span> <span class="ow">in</span> <span class="nb">str</span><span class="p">(</span><span class="n">xaxis_arrow</span><span class="p">):</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">arr1</span> <span class="o">*</span> <span class="n">xspan</span> <span class="o">*</span> <span class="n">aspect</span><span class="p">,</span>
					<span class="n">y</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">arr2</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
					<span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">2</span><span class="p">)</span>
	
		<span class="n">x</span><span class="p">,</span> <span class="n">y</span> <span class="o">=</span> <span class="n">yaxis</span><span class="o">.</span><span class="n">get_data</span><span class="p">()</span>
		<span class="k">if</span> <span class="s1">&#39;+&#39;</span> <span class="ow">in</span> <span class="nb">str</span><span class="p">(</span><span class="n">yaxis_arrow</span><span class="p">):</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">[</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">arr2</span> <span class="o">*</span> <span class="n">xspan</span> <span class="o">*</span> <span class="n">aspect</span><span class="p">,</span>
					<span class="n">y</span><span class="p">[</span><span class="o">-</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="n">arr1</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
					<span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">2</span><span class="p">)</span>
		<span class="k">if</span> <span class="s1">&#39;-&#39;</span> <span class="ow">in</span> <span class="nb">str</span><span class="p">(</span><span class="n">yaxis_arrow</span><span class="p">):</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">arr2</span> <span class="o">*</span> <span class="n">xspan</span> <span class="o">*</span> <span class="n">aspect</span><span class="p">,</span>
					<span class="n">y</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="n">arr1</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
					<span class="n">color</span><span class="o">=</span><span class="s1">&#39;k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">2</span><span class="p">)</span>
	
		<span class="c1"># Change all the fonts to humor-sans.</span>
		<span class="n">prop</span> <span class="o">=</span> <span class="n">fm</span><span class="o">.</span><span class="n">FontProperties</span><span class="p">(</span><span class="n">fname</span><span class="o">=</span><span class="s1">&#39;Humor-Sans.ttf&#39;</span><span class="p">,</span> <span class="n">size</span><span class="o">=</span><span class="mi">16</span><span class="p">)</span>
		<span class="k">for</span> <span class="n">text</span> <span class="ow">in</span> <span class="n">ax</span><span class="o">.</span><span class="n">texts</span><span class="p">:</span>
			<span class="n">text</span><span class="o">.</span><span class="n">set_fontproperties</span><span class="p">(</span><span class="n">prop</span><span class="p">)</span>
		
		<span class="c1"># modify legend</span>
		<span class="n">leg</span> <span class="o">=</span> <span class="n">ax</span><span class="o">.</span><span class="n">get_legend</span><span class="p">()</span>
		<span class="k">if</span> <span class="n">leg</span> <span class="ow">is</span> <span class="ow">not</span> <span class="kc">None</span><span class="p">:</span>
			<span class="n">leg</span><span class="o">.</span><span class="n">set_frame_on</span><span class="p">(</span><span class="kc">False</span><span class="p">)</span>
			
			<span class="k">for</span> <span class="n">child</span> <span class="ow">in</span> <span class="n">leg</span><span class="o">.</span><span class="n">get_children</span><span class="p">():</span>
				<span class="k">if</span> <span class="nb">isinstance</span><span class="p">(</span><span class="n">child</span><span class="p">,</span> <span class="n">pl</span><span class="o">.</span><span class="n">Line2D</span><span class="p">):</span>
					<span class="n">x</span><span class="p">,</span> <span class="n">y</span> <span class="o">=</span> <span class="n">child</span><span class="o">.</span><span class="n">get_data</span><span class="p">()</span>
					<span class="n">child</span><span class="o">.</span><span class="n">set_data</span><span class="p">(</span><span class="n">xkcd_line</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y</span><span class="p">,</span> <span class="n">mag</span><span class="o">=</span><span class="mi">10</span><span class="p">,</span> <span class="n">f1</span><span class="o">=</span><span class="mi">100</span><span class="p">,</span> <span class="n">f2</span><span class="o">=</span><span class="mf">0.001</span><span class="p">))</span>
					<span class="n">child</span><span class="o">.</span><span class="n">set_linewidth</span><span class="p">(</span><span class="mi">2</span> <span class="o">*</span> <span class="n">child</span><span class="o">.</span><span class="n">get_linewidth</span><span class="p">())</span>
				<span class="k">if</span> <span class="nb">isinstance</span><span class="p">(</span><span class="n">child</span><span class="p">,</span> <span class="n">pl</span><span class="o">.</span><span class="n">Text</span><span class="p">):</span>
					<span class="n">child</span><span class="o">.</span><span class="n">set_fontproperties</span><span class="p">(</span><span class="n">prop</span><span class="p">)</span>
		
		<span class="c1"># Set the axis limits</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_xlim</span><span class="p">(</span><span class="n">xax_lim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="mf">0.1</span> <span class="o">*</span> <span class="n">xspan</span><span class="p">,</span>
					<span class="n">xax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="mf">0.1</span> <span class="o">*</span> <span class="n">xspan</span><span class="p">)</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_ylim</span><span class="p">(</span><span class="n">yax_lim</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span> <span class="o">-</span> <span class="mf">0.1</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">,</span>
					<span class="n">yax_lim</span><span class="p">[</span><span class="mi">1</span><span class="p">]</span> <span class="o">+</span> <span class="mf">0.1</span> <span class="o">*</span> <span class="n">yspan</span><span class="p">)</span>
	
		<span class="c1"># adjust the axes</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_xticks</span><span class="p">([])</span>
		<span class="n">ax</span><span class="o">.</span><span class="n">set_yticks</span><span class="p">([])</span>      
	
		<span class="k">if</span> <span class="n">expand_axes</span><span class="p">:</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">figure</span><span class="o">.</span><span class="n">set_facecolor</span><span class="p">(</span><span class="n">bgcolor</span><span class="p">)</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">set_axis_off</span><span class="p">()</span>
			<span class="n">ax</span><span class="o">.</span><span class="n">set_position</span><span class="p">([</span><span class="mi">0</span><span class="p">,</span> <span class="mi">0</span><span class="p">,</span> <span class="mi">1</span><span class="p">,</span> <span class="mi">1</span><span class="p">])</span>
		
		<span class="k">return</span> <span class="n">ax</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<h2 id="Testing-it-Out">Testing it Out<a class="anchor-link" href="#Testing-it-Out">&#182;</a></h2>
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>Let's test this out with a simple plot.  We'll plot two curves, add some labels,
	and then call <code>XKCDify</code> on the axis.  I think the results are pretty nice!</p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[4]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="o">%</span><span class="n">pylab</span> <span class="n">inline</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt"></div>
	
	
	<div class="jp-RenderedText jp-OutputArea-output" data-mime-type="text/plain">
	<pre>
	Welcome to pylab, a matplotlib-based Python environment [backend: module://IPython.zmq.pylab.backend_inline].
	For more information, type &#39;help(pylab)&#39;.
	</pre>
	</div>
	</div>
	
	</div>
	
	</div>
	
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[5]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="n">np</span><span class="o">.</span><span class="n">random</span><span class="o">.</span><span class="n">seed</span><span class="p">(</span><span class="mi">0</span><span class="p">)</span>
	
	<span class="n">ax</span> <span class="o">=</span> <span class="n">pylab</span><span class="o">.</span><span class="n">axes</span><span class="p">()</span>
	
	<span class="n">x</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">linspace</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mi">10</span><span class="p">,</span> <span class="mi">100</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">np</span><span class="o">.</span><span class="n">sin</span><span class="p">(</span><span class="n">x</span><span class="p">)</span> <span class="o">*</span> <span class="n">np</span><span class="o">.</span><span class="n">exp</span><span class="p">(</span><span class="o">-</span><span class="mf">0.1</span> <span class="o">*</span> <span class="p">(</span><span class="n">x</span> <span class="o">-</span> <span class="mi">5</span><span class="p">)</span> <span class="o">**</span> <span class="mi">2</span><span class="p">),</span> <span class="s1">&#39;b&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">1</span><span class="p">,</span> <span class="n">label</span><span class="o">=</span><span class="s1">&#39;damped sine&#39;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="o">-</span><span class="n">np</span><span class="o">.</span><span class="n">cos</span><span class="p">(</span><span class="n">x</span><span class="p">)</span> <span class="o">*</span> <span class="n">np</span><span class="o">.</span><span class="n">exp</span><span class="p">(</span><span class="o">-</span><span class="mf">0.1</span> <span class="o">*</span> <span class="p">(</span><span class="n">x</span> <span class="o">-</span> <span class="mi">5</span><span class="p">)</span> <span class="o">**</span> <span class="mi">2</span><span class="p">),</span> <span class="s1">&#39;r&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mi">1</span><span class="p">,</span> <span class="n">label</span><span class="o">=</span><span class="s1">&#39;damped cosine&#39;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">set_title</span><span class="p">(</span><span class="s1">&#39;check it out!&#39;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">set_xlabel</span><span class="p">(</span><span class="s1">&#39;x label&#39;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">set_ylabel</span><span class="p">(</span><span class="s1">&#39;y label&#39;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">legend</span><span class="p">(</span><span class="n">loc</span><span class="o">=</span><span class="s1">&#39;lower right&#39;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">set_xlim</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mi">10</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">set_ylim</span><span class="p">(</span><span class="o">-</span><span class="mf">1.0</span><span class="p">,</span> <span class="mf">1.0</span><span class="p">)</span>
	
	<span class="c1">#XKCDify the axes -- this operates in-place</span>
	<span class="n">XKCDify</span><span class="p">(</span><span class="n">ax</span><span class="p">,</span> <span class="n">xaxis_loc</span><span class="o">=</span><span class="mf">0.0</span><span class="p">,</span> <span class="n">yaxis_loc</span><span class="o">=</span><span class="mf">1.0</span><span class="p">,</span>
			<span class="n">xaxis_arrow</span><span class="o">=</span><span class="s1">&#39;+-&#39;</span><span class="p">,</span> <span class="n">yaxis_arrow</span><span class="o">=</span><span class="s1">&#39;+-&#39;</span><span class="p">,</span>
			<span class="n">expand_axes</span><span class="o">=</span><span class="kc">True</span><span class="p">)</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt">Out[5]:</div>
	
	
	
	
	<div class="jp-RenderedText jp-OutputArea-output jp-OutputArea-executeResult" data-mime-type="text/plain">
	<pre>&lt;matplotlib.axes.AxesSubplot at 0x2fecbd0&gt;</pre>
	</div>
	
	</div>
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt"></div>
	
	
	
	
	<div class="jp-RenderedImage jp-OutputArea-output ">
	<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAb4AAAEvCAYAAAA6t6QPAAAABHNCSVQICAgIfAhkiAAAAAlwSFlz
	AAALEgAACxIB0t1+/AAAIABJREFUeJzsnXd4FNXXx7+zm00PSSAhhIQivTcpAkpvAhEQBQWkSywg
	oij40lTqT0GsqFSlSe+R3luQXgKE0EtCT+/ZPe8fJ7ObQIBssrszm9zP8+wzk+zu3DPJ7nznnnuK
	REQEgUAgEAgKCRqlDRAIBAKBwJYI4RMIBAJBoUIIn0AgEAgKFUL4BAKBQFCoEMInEAgEgkKFED6B
	QGARbty4gYkTJ+LmzZtKmyIQPBchfAKBwCL8/fffmDBhAj799FOlTREInosQPoFdsGDBAvj7+0Oj
	0eDll1/G/v37jc9Vq1YNd+/ezfF9p06dAgDo9XqMHDkSjo6O0Gg0Tz0CAgKQkpJifJ/8emdnZzg4
	OKBLly64c+cOAGDx4sXw9vbG4cOH83w+I0aMgLu7O65du5bnYzyPkJAQlCtXznh+JUqUwOXLl1/4
	viNHjiAwMBBxcXFmjzlw4EA4OTlh+/btiI2NzYvZAoFNcFDaAIHgRcycOROff/45/Pz88O233yI1
	NRXLli3Da6+9BgBITk5GWlpaju9t1KgRUlNTsXr1avzwww/w8fHBN998Aycnp2yva9CgAZydnQEA
	6enp6NWrF1avXo2aNWsiODgY+/fvx/79+/HOO+/g8uXLiI2Nxdq1a9G4cWN88skn+Ouvv3DlyhX4
	+vrm6px27dqF5ORk3Lp1C8WKFUNSUtJzX1+0aFE4Ojrm6tiRkZHo3r070tPT8cMPP+DixYtYsGAB
	2rdvj0OHDsHPz8/42vj4eHh4eBh/liQJkZGRmDdvHkaMGAEAiIuLy5V9gYGB6NmzJxYuXIjTp0+j
	WbNmubJXILA5JBComMWLF5MkSVSxYkW6efPmU88/evSIPD09c3yOiEiSJCIiun//PkmSRPPnz3/h
	mIMHDyZJkuiNN96gtLS0p56fMGECSZJE27ZtowsXLpBGo6EyZcrk+Npn8eWXX5IkSbRnzx569913
	SZIk0mg05OTkRJIkPfWYPn16ro/dunVr8vLyop9++sn4u7Vr15KjoyO99tprxt+tW7eOtFotHTt2
	LNv7GzVqRKVLlzaezzvvvJNr++S/zd69e3Ntr0Bga4TwCVRLZGQkeXl5kbe3N929ezfH1+zevZuq
	Vq1KBoMhx+dl4UtOTiZJkmjkyJEUFRVlfMTGxmZ7fUhICEmSRK+99hplZGTkeEz54j516lSqXbs2
	BQQE0JUrV8w6t0qVKmUTiKNHj9LFixfp1q1bJEkSde3albZu3Upbt26l0NDQXB/34MGDJEkSbdy4
	8ann5JuI06dPExHRN998k+Nr9+7dS5Ik0axZs4y/y619M2bMII1GQydPnjTr7yEQ2BIhfALVMmXK
	FJIkif73v/898zW7d+/ONouJj4+nqKgounbtGo0YMYIkSaLbt29TbGwsSZJETk5O5OzsTJIkkU6n
	o9dffz2baLZr1440Gg0dOXLkmWPKwic/+vXrZ/a5lSlT5pkCIUkS/f3332Yfk4iobdu21KdPnxyf
	MxgMVL58eZowYQIRETVu3Jg8PT0pMTEx2+vS0tLIx8eHHBwc6Pr162bZl5aWRhs2bMiT7QKBrRBr
	fALVkpqaCo1Gg379+j33dadOncL58+fh5eWFVq1a4dKlS8bnJElCQEAAFi9eDADYsWMHatasif/+
	+w81atSAv7//U2NWqVIFDRs2zLWdS5cuRalSpTBx4kQzzg4oVaoU6tSpY9Z7nsft27exa9cuXLhw
	IcfnJUmCu7s7wsLCEBsbizNnzqB79+5wdXXN9jqdTodWrVph5cqVZgep6HQ6BAUF5fkcBAJbIKI6
	BaqlbNmyMBgM2L59+3Nfl5CQgBo1aiAwMBCXLl1C69atsWXLFvTp08f4moyMDADAq6++Ck9PT7Rt
	2/Yp0ZPHDA8Px40bN15o3+jRoxESEgInJydMnjwZX3zxRa7OKy0tDfHx8dBoLPv1W7FiBTp37oyK
	FSvm+Pz69etx4cIFDBs2DNHR0UhKSoJWq83xtWfPnkXx4sVRvnx5i9ooEKgBIXwC1dKrVy+UKVMG
	w4YNw71797I9d/ToUXz77beQJAlarRbDhw/Hli1bsGXLFixatAjt2rVDy5YtQZldtyIiIgAAjx8/
	RnJyMu7evWt8ZA3dHz16NIgIffr0gV6vzzbm/PnzsXbtWtStWxcAzw5ff/11bNmyBR4eHpgxYwbW
	rl37wvM6f/48oqOj8/W3yYmQkBB4e3s/9fuUlBRMnz4dPXr0wOTJk18YbXnmzBmEh4fjp59+gpub
	m1k2EBEuXbr01N9OIFATQvgEqsXR0RErV6405u5NnToVYWFhmDZtGpo2bWoM72/SpAlmzpyJdu3a
	oV27dihRogQAdu3JhIWFAQAqVKiAUqVKoWTJksbHBx98YHxdlSpV8Oeff+LQoUNo2LAh5syZg4sX
	L6Jfv34YPHgw/P398cYbb2Do0KEoU6YMAKBp06aYM2cOAGD16tUvPK9Dhw4BYLegJalUqRLWrVuH
	L774AnPnzsXcuXMxZcoUlCtXDtOmTcPMmTMxcuRIADDO9AwGw1PHCQ4ORt26ddGzZ0+zbfjxxx9R
	pUoVVKpUyWo5igJBfhFrfAJVU79+fRw9ehRTpkzBuHHjMGbMGOh0OgwdOhRffvkl9u3b90yXYYMG
	DVC7dm0AQMWKFVG2bFm0aNECAF/4O3bsCA8Pj6dmQIMHD0aVKlUwbtw4BAcHAwB8fHwwZ84cvPLK
	KwCAn3/+Odt7evTogdTUVFSuXPmF53T//n0AnGP4JLdu3Xrh+5/FtGnT4Ovriz///BMPHjyATqdD
	cHAwZs2aha5du2Z7balSpVCjRg1s3LgRycnJcHFxAcDCdfToUZw7dy7HMV5kn7e3NyRJgqOjI9zd
	3fN8LgKBVVE4uEaQybVr12jTpk107949pU2xK/bv30/Dhw9X2gyzkKNCBw4c+NRzBw8eJHd3d7p2
	7ZrV7Vi4cCFJkkTBwcE0Z84cGj9+PDk4ONDChQuf+R5b2icQWAuJKHMRRKAYFy5cQNu2bXHnzh00
	b94ce/bsUdokgRWZNm0apk2bhn///RdNmjRR1JZu3bph/fr1ANhlPGXKFFFxRVDgEcKnIGlpaZg2
	bRomT56creTWr7/+io8//lhBywQCgaDgIoRPIUJDQzF48GBj0EXlypURHh4OgIMy1qxZ89S6jEAg
	EAjyj4jqtAEPHjxAeno6AM45Gz58OJo0aYKwsDBUqFABu3btMgY6dOvWDUSEd999N1/V/wUCgUCQ
	M0L4rMzjx48xfPhwY3uc3r174+eff4ZGo8Ho0aNx5swZtGzZ0ujq7N69O95//32kpKQgKCgIDx8+
	VNJ8gUAgKHCIdAYrs2vXLixbtgyDBw8GAEyYMAH37t3DH3/8ka1clSx8Tk5OmDVrFu7du4f27dvD
	x8dHEbsFAoGgoCKELx+kpaUhJSUFRYoUAQBjtQo5r0ySJBw9ehRVq1ZFjRo1AAD16tXD4cOHsyVX
	AzC6QnU6HRwcHLBu3bqnXiMQCASC/CNcnflg9erVCA4OxtatW5GRkQGtVgutVgtJkoylsv777z9o
	tVrEx8cb35eToMlVPOSZnxA9gUAgsA5C+PKIwWDA1atXsXz5cnTr1g3NmjXDzJkzcejQIcTGxhpn
	feHh4TAYDDnWUMyKXBMxMTHR6rYLBAJBYUa4OvOIRqPB+++/j+vXr2PJkiUIDQ1FaGgoAKBGjRoY
	PHgwjhw5YiyEXLRo0eceTy7vJIRPIBAIrIsQvnxQvHhxzJkzB9999x1OnjyJjRs3YteuXbh48SJG
	jhxpbIXTsWNHALwG+Kw2MB4eHgBgdv8zgUAgEJiHEL48YDAYoNFosGzZMjRt2hSlSpVCq1at0KpV
	K0RGRmLv3r1Yu3Yt7t27h/r162PQoEEA8Nz+ayVLlgTAzUQFAoFAYD1E5ZZ8oNFoMGvWLAQHB8Ng
	MGSbzUVHR+Phw4eoUKFCrgJV1q5dizfffBOdO3fGxo0brWm2QCAQFGrEjC+PREVFAQCKFStmbIYK
	cCNOSZLg7e39woCWrJQqVQpA/trSCAQCgeDFiKhOM5EnyAcOHICrq6uxGenJkyexbNky4+wuLCwM
	EydOzPVxhfAJBAKBbRDCZyay8O3atQsVKlRA8eLFAQCzZ8/GsmXLjK/btGkT5s+fj9x6kn19feHo
	6IjHjx8jKSnJ8oYLBAKBAIAQvjxz6NAh1KxZ0+jOPHr0KPz8/IxCd+TIEZQvXz7XUZoajQaBgYEA
	RICLQCAQWBMhfGYiR2ZevHgROp3OWK7s+vXr6NSpk9HVeeLECTg6Opp1bH9/fwCm9UOBQCAQWB4R
	3JIHHj9+jMDAQPz1119wd3dHbGws0tLS0KZNGwCc7qDX6+Hq6mpMTM8NQvgEAoHA+ogZXx7w8vLC
	iBEj4OPjg927d2Pbtm1IS0vD4MGDsW3bNvz11194/PgxGjRoAAeH3N9blChRAoAQPoFAILAmYsaX
	BzQaDYYOHYqhQ4ciISEBFy5cwPr167F48WKsWLECAM/6UlNTAZhSHF6EXNZMVG8RCAQC6yFmfGYi
	B6+cOXMGU6dORUZGBho0aIBJkybh+vXrOHXqFEaNGoXOnTvj5ZdfNuvYcocGuUWRQCAQCCyPmPHl
	gRUrVmDYsGF48OABWrRogcaNG+Pq1auIjY1FnTp18PXXXyMuLg5eXl4Act9iSA6GkVsTCQQCgcDy
	COEzk+PHj2Pw4MHo1KkTKleujNmzZyMuLg5fffUVrl+/Dh8fH2zZsgXlypUz+9jyeqBc3FogEAgE
	lke4OnOJwWAAAISEhMDX1xe//fYbunXrhnXr1uHrr79G06ZNcfXqVbRs2RLBwcEAkOvkdZmEhAQA
	MCsSVCAQCATmIYTPTHbu3In69evDyckJtWvXxksvvYTAwEB8//338PLyQsWKFbFz505s3rzZ7C7q
	0dHRAGBWjU+BQCAQmIcQvlwiJ667urriwoULSE5OBgAkJSWhWrVqRpHz9fUFYIrMNGfW9/jxYwBC
	+AQCgcCaCOEzkyFDhkCn08HHxwdEhEWLFqFfv37G9TnZTdm4cWMAuQ9sAUwzvhd1axcIBAJB3hHB
	LblEbj5boUIFVKpUCbGxsfD09ESDBg2yve7ixYvZujaYg3B1CgQCgfURwpdL5JlbVFQUDh8+jKJF
	i0Kn06FixYpo3rw5goKCULlyZZw7dw41a9YEAOj1+mzNaV+EcHUKBAKB9RHCl0tk4atbty727duH
	0qVL4/bt29iyZQt2796Njz76CJGRkUhNTcWnn36a7T25Rcz4BAKBwPpIZG7MfSEla9mxlJQUSJIE
	Jyenp1538+ZN+Pn55fjci/Dy8kJsbCyio6ONye8CgUAgsCxixpdLJEnC7t27ceDAAYSHh+PmzZtI
	T09HlSpV4OTkhLS0NHh5eaF69eooVaoUWrdubZabEzAlrptT2FogEAgE5iGusLlg9+7d+Oqrr3Dp
	0iXExMSgTJkyeOmllyBJEv7++294eHjAyckJsbGxqFGjBgYNGoR27dqZPY5co1Ou2SkQCAQCyyOE
	LxesWLEC//33H4YNG4a+ffuiYsWKKFKkCIYOHYrbt29j/vz5qFOnDlJTU5GQkJDnyitixicQCATW
	R1xhc4HcZaFSpUrZOi60adMG+/fvR4kSJeDp6QkAKF68eJ7GMBgMMBgMkCTJbBepQCAQCHKPCG7J
	JR999BEWLVqEDz/8EO+++y7q1q0LgCu67N69G82bNzcKl7nRnACQmpoKZ2dn6HQ60Z1BIBAIrIiY
	8eWSsWPHQqfTYdGiRZg+fTpatmwJJycnODs749atW8YEd1n8zEV2c4r1PYFAILAuYsZnJv/99x8W
	LFiAlStXGhPOq1Wrhi+++AK9e/fO8/pcTEwMvL29UaRIEdGBXSAQCKyIEL5ckjWPDwAePXqEI0eO
	YO/evdi7dy/Onj2L5ORkTJ06FaNGjTL7+A8fPoSvry+KFSuGhw8fWtJ0gUAgEGRBCJ+ZPGsdLy4u
	DnPnzkW9evXQokULs48bFRWFkiVLws/PD3fv3rWQtQKBQCB4ErHGZyZyeyIigl6vB8Drf2+++SaG
	Dx9ufN5cRA6fGRABjx7x1ssLEH8zgUBgBqItUS7IaVIsSRIcHBxw7tw5fPfdd7h8+TK0Wm2eAlsA
	kcP3QpKSgAULgA4dAG9vwNcXKF4ccHcH6tcHpkwBIiOVtlIgENgBQvieQ9bcOvlnWQTl7bFjx+Dr
	64smTZrkaywx43sGBgMwezZQvjwwcCCwdSsQGwsUKQIUKwakpQHHjwNjxgAVKgDjxgGpqUpbLRAI
	VIwQvudw+PBhfPXVV9i7dy8AdnPKIii7OUNDQ6HVao2dFfJKaubF2tHRMV/HKVA8eAC0bQsEBwN3
	7wL16gFz5gB37gAxMcDDh0BcHLBpE9ClC5CcDEyaBLzyCnDjhtLWCwQClSKE7zkcOHAAP/74I/r0
	6YM+ffpg/fr1uHPnDgCTS/Ls2bNwdnbOdzeFxMREAICbm1v+jC4onD4NNGgA7NrFLs1ly4Bjx4DB
	g4GSJQHZpezhAXTqBKxbBxw4wDPDU6eAJk2Ac+eUPQeBQKBKxILScxg4cCCOHz+OVatWYenSpVi6
	dCl0Oh0aNWqEYcOG4fLlyzh//jwSExMtJnx5rfNZoDh0CGjfHkhIABo2ZFHz93/x+5o2ZXHs0gXY
	tw947TVg926gTh3r2ywQCOwGkc7wAvR6PcLCwrB3715s374dhw8fxqNHj1CkSBHExcUBAMqUKYNr
	167la5z169eja9euCAoKwoYNGyxhun1y9CjQpg27MHv2BP76C3B2Nu8YKSnAu++aBDM0FChd2irm
	CgQC+0PM+F6AVqtFrVq1UKtWLQwbNgznz5/H3r17sWLFCty9exelS5fG8OHDAcBYtiwv3Lt3DwDg
	6+trMdvtjitXOGozLg7o0QNYvBjIS5SrszO7Rjt0APbsYVdoaCgg3MgCgQBC+HKFHN3p4OCAatWq
	oVKlSmjXrh0ePHiAKlWqGN2ceRU9AMa1w5IlS1rEZrsjLg4ICgIePwY6dsy76Mk4OQFr1pjW+oYO
	5XQIgUBQ6BHBLblAo9Fky69zcHBA+fLl8corr+R7bU9GFr6AgACLHM+uIAIGDAAuXACqVQP++ccy
	Sene3sCqVYCLC7tMFy/O/zEFAoHdI4QvjxBRjonteSUyM/m6UArfvHk8O/PwADZs4Bw9S1G9OvDz
	z7z/ySecFiEQCAo1QvjySF777j2LQuvqDA8HMtdI8fvvnI5gaQYN4ijR6GgWP4FAUKgRwqcSCqWr
	U68H+vThcmS9e/PDGkgS8OefHNyyciWwfr11xhEIBHaBEL4XkLVMmbVISUnBo0eP4ODggOLFi1t1
	LFXxyy+cd1eqFPDbb9Ydq0wZrucJ8AwzJcW64wkEAtUihO8FyGXKDAaD1caIiooCAPj7++crMtSu
	uHEDGDuW93/7DfD0tP6YH38M1KzJY//0k/XHEwgEqqSQXGXNJyEhAT/++CNmzJiBxMREqwpSVuEr
	FBCxCCUmAm+/zWkMtkCrBWbM4P0pU4D7920zrkAgUBVC+J6Bm5sbXF1dMWnSJLz//vu4efMmAFM0
	pyXdn/czL8B+fn4WO6aqWb0aCAnhWZ6tZ15t2wKvv855g5Mm2XZsgUCgCoTwPYchQ4bg+++/x4ED
	BzB27FjcvXvXGM1pyYhOWfgKxfpeUhLw2We8P21a7mpwWpr//Y+3s2eL9AaBoBAihC8HiMgobIMH
	D8b8+fOxePFi1K9fH7Nnz0ZYWBju3r2LpKQki4wnC1+hKFf23XfArVtA3brA++8rY0PNmkC3bty3
	b/p0ZWwQCASKIYQvByRJMjaGPXPmjFGYoqKi8OGHH6J9+/YYMGAAxo8fbyxUnR8SEhIAAEUsmbit
	Rm7eNM22fvqJ19yUYswY3v7+O/f1EwgEhQZRq/MJHj9+jM2bN2P58uUICwtDXFwc/Pz80Lp1a3Ts
	2BGlS5fGwYMHERISgtDQUAQEBGDEiBH5GlNuQuvk5GSJU1AvX3zBaQQ9e3LLICV5+WUuYr1lCzB/
	PvDll8raIxAIbIYQvicIDQ3FxIkTIUkSWrdujerVq+O1115DvXr1jK954403MGDAAKSmpqJy5cr5
	HjMtLQ1AARe+ffuAFSu4buZ33yltDTN0KAvfn38CI0cChSWVRCAo5Ajhe4LmzZvj8OHDKFKkCAwG
	A3SZxZL1ej20ma45nU6HmjVrWmzMjIwMADAev8BBZApoGTVKPb3xOnTgxParV4EdO4B27ZS2SCAQ
	2ABxi/sEbm5u8Pb2hlarNYoekF2ULJ3O4OHhAQAWWS9UJevWAcePAyVKsLtTLWi1wJAhvD9vnrK2
	CAQCmyGELw9YOp3B29sbABAdHW2xY6oGvR4YN473x4wBXF2VtedJ5PqgmzZxqoVAICjwCOFTAQVa
	+JYvB8LC2L2pVPrC8yhTBmjYkEVv82alrREIBDZACJ8KKLDCp9cDX3/N++PHc1d0NfL227xduVJZ
	OwQCgU0QwqcC5C7uBU741q4FIiKAcuWAfv2UtubZvPkmbzdvBjLzNwUCQcFFCJ8KkGd8MTExClti
	YX74gbeffQY4qDiAuFw5oFo1rt+5f7/S1ggEAisjhE8FFEhX5+HD/PD2Bvr3V9qaFyN3iNi0SVk7
	BAKB1RHCpwIKpPDJs70PPuDO52qnc2febtzIeYcCxQgL46XhNm24R7GHB+Dnx8V2PviA/0WZNR8E
	gjwhkbXbiwteSHp6OhwdHaHVapGenm7RVAlFuHYNqFCB8+SuXwdKllTaohej1wPFiwOPHwMXLwIW
	qMgjMI8dO1jwDh588Wv9/IBPPgE+/VR9GTIC9SNmfCpAp9PBzc0Ner0e8fHxSpuTf37+GTAYgHfe
	sQ/RA1ikO3bk/Y0blbWlkHH3LscXtW3LoufpyZkva9YAV64A0dHAnTvA3r3AxIlA9erAvXucFlq1
	KtdHELfvAnMQMz6VUKpUKdy+fRs3btxAabWU9MoLcXFAYCAQHw+cOMHth+yFFSu4gHbz5sCePUpb
	UyjYv5//5FFR7BH/v/8Dhg1j9+azIAJ27uTyqqdP8+969ADmzAEKeoMTgWUQMz6VUGBSGhYsYNFr
	3ty+RA8A2rfn6NMDB3iaIbAqv/4KtGzJovfaa8CFCyx8zxM9AJAkXv87doydCx4efM9Svz5w9qxt
	bBfYN0L4VILci8+uXZ16PffZA4B8tmpSBE9PoFkzPo8tW5S2pkAzeTLP7PR67gi1axcHspiDgwMf
	49gx7i0cEQE0aQJs324dmwUFByF8KkEuVG3XwrdxIwe2lCtnipK0N7JGdwqswtdfA2PH8sxt3jzu
	TZyfNM9KlYDQUF5STkjgpdrFiy1mrqAAIoRPJRQI4Zs5k7effKJsd/X8IAvf5s1AZrsogeX44w/g
	m2/447F0KTBwoGWO6+oKLFnCzT8yMoC+fYG//7bMsQUFDyF8KsHuhe/AAW426+kJDBigtDV5p2JF
	oEoVICaG/W8Ci/Hvv8DHH/P+7Nk8Q7MkGg33OJ46lQNgBgwQMz9BzgjhUwl2L3wTJ/J2+HD7D63r
	2ZO3//yjrB0FiOvXgT59OMtl/HjLzfRyYvRoYNIkFr9+/bhBiECQFSF8KsHd3R2AnQrfkSPAtm2A
	uzsLn73z7ru8Xb0aSE5W1pYCQHo630tER7MnecIE6485Zgy7VA0G4L33gN27rT+mwH4QwqcS5G7v
	er1eYUvygDzbGzYMKFpUWVssQeXKXB8rPp79c4J8MX068N9/3JLx77/ZJWkLxo/nyi7p6UC3bsC5
	c7YZV6B+hPCpBE3m1cDuhO/4cSAkhKML7DGF4Vn06sXbpUuVtcPOiYjgmRcAzJ1r+/ui6dO5Kkxs
	LEd7RkXZdnyBOhHCpxK0mVGQdid88mzvo48AX19lbbEkPXtyvH1ICF81BWZDBAQHA6mpHGXZtq3t
	bdBqOcClSRPg1i3uOSwKXAuE8KkEWfgMBoPClpjBmTPA+vWAszPw+edKW2NZAgKAFi34qr12rdLW
	2CWrVvHamo8PMGOGcna4uHDdz4AArgX62WfK2SJQB0L4VIZSpVOJ+I740CGukv/ff1w8+LlMncrb
	IUOAEiWsbqPNEe7OPJOezuXHAI6w9PFR1h4/PxY/R0fgt99Ejl9hRwifSkhKSgIAuNqwx0p6Ot+V
	v/02UKwYBx80bcouqUaNAH9/oEwZzr06cuSJN0dEcIFEnY6rBRdEunfn89u5Mxd3AYKszJkDXL7M
	VVUGDVLaGqZhQxY9gPv6nTihrD0C5RDCpxISExMB2Eb4Hj7kcO9SpVj0Vq3iUHMfHxa8Vq04qNHD
	A7h5E5g1i71+jx5lOcj//sex4n37ml9k0V7w9uaICIOBRV6QK5KSgG+/5f2pU/NXjszSDB7MLY9S
	UoC33uI6BYLCh4o+koUbecbnZsVu5UlJwI8/smbFxfHvqlVjT2XnzkD58tlfr9cDJ09yArDBwLNC
	AOwTXbiQ49JHjbKavaqgVy9ex1y6lEuxCV7I7NncL69+fU4jUBs//8zByCdOAP378xKuvfd+FpiH
	ED6VIM/4rCV827axwN24wT+3b8+zvldfffaXXqvli1f9+ix8RqZPZz/pO+9wia+CTOfOnJh/5Ah3
	RX3y7kCQjeRkvrECOI9OjYLi7AysXAnUq8f3NDNmFFxvvSBnhKtTJVjL1Rkfz2ss7duz6NWuzcEr
	W7ZwD7TcXpiMScf37/MCDgB89ZVFbVUlrq6maYsoYfZC5s3j5dA6ddTdoKNcOVOAy+jRXGpWUHgQ
	wqcSrOHqPHcOaNAAmD8fcHLi9ZZjx4DWrfNx0B9/5Nv6oCCgVi0AvF5SoBsZyNGdS5Zw+KsgR1JT
	gWnTeF+ts72sdOnC3Rz0ek7bvH9faYsEtkIIn0qw9Ixv4UKOYgsPB6pX5/WM0aPzGWgQE2MKi5Nj
	1QH88gsfPcdIAAAgAElEQVQHQBbYspatW3Ny/sWLwOnTSlujWv76C7hzB6hRg0XFHpg8md39kZFA
	794sgoKCjxA+lWCpGZ9ezzEY/fqxEPXty8tT1apZwMjffuOomFatgFdeAcBrf/PnAxs2sPgVyKoY
	Oh3Qowfvi5y+HElLA6ZM4f1x42xXjzO/6HTAsmV8X7Njh6kQkaBgYycfz4KPJfL4EhKArl15Bubo
	yEtxf/0FWMR7mpjIbk4g22xPo+FAAR8f7t3at28BvWuW3Z3//PNEpI8AABYt4tSXqlX5BsieCAjg
	+xlJ4jSMbduUtkhgbYTwqYT8RnVGRgLNmgGbNnEh4B07OGfJYussc+ZwAqCc6JeFGjU4WMbDg1Mf
	CmSGQ+PGQNmywO3bIhLiCfR609re2LEcDWxvtGkDfP01L+H27s3/ZkHBRQifSsjPjO/MGdajkyeB
	ChWAw4c5YtNipKZyCgPAs70c1PTll9nd6eDA4eFLllhwfDUgSaY+fQXu5PLHhg1cpeWll0weYXtk
	7FigXTu+v+vZkzN2BAUTIXwqIa/Ct3cvi9zt21xu7PBhLhNlURYu5KiFmjWfG6PeogUnBwM82wwL
	s7AdSiML35o1BTyM1TzkAtSffqquKi3motFwJ4fAQK5ZO3q00hYJrIVESlVFFmRDypxF6fV6Y2++
	F7F+Pd+ZpqZy+aVFizg516JkZABVqnDy9tKlpov/MyACBg7ktcVatbjYtZOThW1SCiJexAoPZ19y
	vvJCCgZHjnCck5cXF/Rxd1faovxz+DAvG2RkAKtXcz8/QcFCzPhUhpTLRbkFC/gLmZrKBXeXLbOC
	6AFco/LKFfah5sKPJUkcXFO+PLtg5SakBQJJ4jsMgAucCoyzveDggiF6AC/nfvcd7w8YwG5cQcFC
	zPhUglarhcFgQEZGhrE337P4/nvgyy95f9w4FherJAsbDFzq5dw5Dm4ZPDjXbz10iF2vOh0LYJUq
	VrBPCU6f5rIkxYtzRJE9RnJYiLt32S0IcFWggABl7bEkRHyPs2YN/7sPHeK+foKCgZjxqQTZvfm8
	RrRELHiy6P30E4dfW61CxoYNLHqBgZynYAZNmrBOpqcDQ4cWoIIntWrx7Pf+fWD/fqWtUZSFCzmi
	s3PngiV6AH+n5s9nz8WpU8Dw4UpbJLAkQvhUwouELyODa25+/z0HECxZYuVmAQYDMGEC73/5JScG
	msnUqZxasXNnAfIMCncnAL6RmTeP99XSb8/SeHryv9jJiR0eCxcqbZHAUgjhUwnPE760NI4pWbCA
	3S0bNpjyqa3GmjXsowwM5AZmecDHh0tCAdwJosAEQsrCt3p1oU1mP3IEuHQJKFECeP11pa2xHnXq
	AL/+yvsffMApQwL7RwifSpCF78klV7lh5qpVQJEiHExo9QuNXm+a7Y0Zk6+omUGD2DMYEWGqhm/3
	1KvHyex37/LiTyFkzRre9uxp3ykMuWHQIO7bl5zMbl2R3G7/COFTCRqNBtWrV4dzFpFJSuJivxs3
	sstw1y5eO7M6y5cD588DpUtzbkI+0OlM3bi/+YaF3O7J6u5cuVJZWxSAiJu3AupsNGtpJAn44w9O
	cYiMBDp14nZfAvtFCJ9KqFOnDvbs2WOc+SUk8Bds2zYuoLt7N1dHsTqJiaaaY+PH52lt70l69uSY
	kFu3+AJSIHj7bd4WQndnWBiH+Pv6cmeDwoCTE5cD/P13Lsadmqq0RYL8IIRPJWzcuBE+Pj4AgNhY
	Lp20Zw/g78/VWTJb31mfSZPYl1OvHvt3LIBGY1rrmzaNZ7J2T4MGQKlSXNHmyBGlrbEpspvzjTcK
	VzaHhwev83XqxOvXAvtFCJ9K8PLyAgA8fswFcw8fZk/jvn1cLMQmhIebMpJnzbLoVa1TJ6B+feDe
	PVMDd7umEEd3FiY353NJSWE3xtmz/EXduZO/uKdOAVevFtAeXQUDkcCuIu7fB9q25WDKcuV4Ta9M
	GRsNTgR06MC+1UGDgLlzLT7Exo08S/D35+uCVSrN2BI5S790aeD6dfW3HLcA167xZ9PdHXjwoAD8
	D80hIgJYtw44dowLGUREPN/NrdFwgmP58rxO0bgxP0qWtJ3NghwRwqcSIiO59OPFi0DlynzzaNOk
	4DVruJGalxfHqfv6WnwIIvagnjrFIeIff2zxIWyLwcDuzshIdnc2bKi0RVbnhx+Azz/nddtly5S2
	xgbExwN//snVq0+fzv6cVgv4+QHe3vzQ6Tj0MykJiI5mN3hOwlijBhAUxHeBjRoVihsm1UECxbl+
	nah8eSKAqEYNort3bWxAQgJRqVJswKxZVh1qzRoeJjCQKCXFqkPZhmHD+IS++EJpS2zCq6/y6S5b
	prQlViYpiWjaNKJixfiEAaIiRYjee49o/nyiEyde/AFOTSW6fJlo0yai8eOJ2rUj8vAwHQ8gqlCB
	aPJkotu3bXNeAiIiEsKnMBERJs15+WWihw8VMOL//o8NqFePKCPDqkPp9UQ1a/Jwf/xh1aFsw759
	fDLlyhEZDEpbY1Xu3iWSJCJHR6LYWKWtsSL//UdUpYpJnBo3Jlq3zjJ3aqmpRNu3E33yCVHJkqYx
	NBqijh2JQkL4SyKwKkL4FCQsjMjfnz/3TZoQxcQoYER4OF/JAKJDh2wy5IoVPFzp0nwdsGsyMoh8
	fPiEzp9X2hqr8ueffJqdOiltiZUwGIimTiXSavlEq1Qh2rLlhTc0BgPRuXNEu3bxRDAtLZfjZWQQ
	/fsv0VtvEel0JhGsVo1o3rwC4hJRJ0L4FCI01HS9bNGCKD5eASMMBqL27dmIAQNsNqxez99tgGjO
	HJsNaz369OGT+f57pS2xKh068GnOnau0JVYgJYWob18+QUki+uwzdnfmkvh4ogYN+O116uRhRvzg
	AdF33xEFBJgEsEQJokmTiB49MvNgghchhE8BNm4kcnHhz3bHjtm/XxlWdjVmQ15w8/IiunfPduMS
	0dKlPPRLL5lxh6xWli0z3cEUUGJieFKi0RDdv6+0NRYmKYmoTRv+H7q6slszD9y/z0t2AFHr1nmc
	sKWmEi1aRFS7tkkA3dyIRowgunUrT3YJnkYIn43580++eABEAwdmv+jfunWLbt68aRtDEhPZ1wgQ
	/fqrbcbMQkYGUeXKPPzixTYf3rJER7N7TKvl/QKIfKPSrJnSlliYpCSitm355Pz82FeZD65c4cMA
	RO+8k4/lOoOB1wLbtTMJoE7HF42LF/Nlo0AIn83Q64nGjjV9hseNy750cPPmTSpbtixFRETYxiDZ
	mDp1rB7Q8izmzmUT6tYtAHEhzZvzySxfrrQlVuGtt/j0fvxRaUssSFKSSViKF7fYGu2JE6bgzc8+
	s9ABe/Y03TFLEtGbb3IQjiBPCOGzAfHxRN26mYK3copmbNy4MQGg8PBw6xt06ZIpoOXgQeuP9wyS
	k/l6AxDt3KmYGZbhu+/4RPr2VdoSi5OUxN42gFNvCgR6PVH37nxSvr4caWZBtm83xatMn26hg0ZE
	EAUHm767sk91+/YCcOdoW4TwWZkrVzg3DyDy9OQgrpyoXLkyAaALFy5Y1yCDgej119mg/v2tO1Yu
	+PZb01qnXXP+PJ+Ij49iM2hrsWGDKdulwDBypOlLeeaMVYZYssSkTxZ150dGEn35ZfacwPr1iVat
	KnCfPWshhM+KrF9P5O3Nn8vKlTlz4FlUrVqVAFCYhe88n2LdOtMX3sYBLTnx4IEp0Mfap25VDAaO
	1AGIDh9W2hqL0r8/n9akSUpbYiF+/51PyMGBaMcOqw41Y4ZpqK1bLXzw6GiiKVNMbhOAqFIlXkMQ
	qRDPRQifFUhJIRo+3PRZDAp6cY5e9erVCQCdPXvWeoYlJhKVKcNG/fKL9cYxkw8+YJOGDlXaknwi
	V3EZM0ZpSyxGerqpeIld35jI7NljytObP98mQ37+OQ/n7k507JgVBkhKIvrtN6KyZU0XnYAAVt24
	OCsMaP8I4bMw589zBRY5CGvGjNy532vWrEkA6PTp09Yz7quv2LDatfmKphJOnTJNQhMSlLYmH2ze
	bCrBU0DYtcs0kbD7ZaSoKM6NA9hVaCP0eqJevUwxNFaLX0tPZ/+qXBpJTlUaNUqURHsCIXwWIj2d
	iz44OZny044cyf37a9euTQDo5MmT1jHw+HG+05Ukm1VoMYfGjfnvNm+e0pbkg8RE/gBIUoFJdpNn
	46NHK21JPsnIIGrZkk+meXOb3/ilpppSBUuUILLm/S0ZDFwfVC6sCnCRBYER0Y/PApw9y91GvvqK
	OzMPGgScPGlesX6587rBGt2809PZKL0e+OQTNlZlfPABb+26Q7urK9CsGV9qduxQ2pp8k54OrFzJ
	++++q6wt+WbSJGD3bu6m8M8/gIODTYd3dOQGKK1aAXfv8sdk/34rDSZJ3ABz/34gNBTo0QP49FMr
	DWafCOHLB48ecWudOnW4RVfp0sDWrdzKztPTvGNJma1JiKzQJWrSJO4F9NJLplboKuPtt7mzy9Gj
	bKrd0q4db7dtU9YOC7B9O3/Gq1UDatZU2pp8cPIkfwcAYOlSbgipAB4ewL//cvev2Fj+qPzzj5UH
	bdQIWL6c+wEKjNid8BG7ZxW1IT0d+PlnoGJFblQOAEOH8sxPvu6Zi9VmfDt2ABMn8l3gvHmAm5tl
	j28hXFyAXr14f9EiZW3JF1mFT+HPaX5ZupS3vXrZccu4tDSgf38gIwMYNoynXAri5MQ6FBzMDdx7
	9QI++4zNE9gQhV2tZnHw4EFq1qwZLVaoxlVKCkdCy4GRAPvtLRGI2bBhQwJAoaGh+T+YzJ07nJwL
	EH39teWOayVCQ02Vo1QUe2MeBoMpgMKaEbpWJjHRlLR+5YrS1uSDCRP4JMqXV1XklMHAgZgODqYy
	r1FRSltVeLAL4Tt79iy98cYbBIAAUP369clgwxCz+Hgu1ZS1fVaVKpynZykzXn75ZQJAR48etcwB
	U1KImjY1qbMdJLYaDEQVK7LJmzcrbU0+kKv8W6xkh+2R6243aqS0Jfng6lVTtNmePUpbkyMHDpju
	k4oVI1q5UmmLzEOv19OSJUvowYMHSptiFrZd4TWT69evY8KECVi0aBGICK6urhgxYgRGjhxpXBOz
	JmfPcrDFokVAfDz/rlYtYOxY4M03Aa3W9NrU1FTs378fx48fh0ajgU6nMz6cnZ3h4eGBIkWKPLV1
	c3ODRqNBRqavQ5v1oHnFYGD3zsGDQEAAsGRJdmNViiQB770HjB/Pf/MOHZS2KI+0bQssXAjs3Qt8
	/jkAIC4uDidOnAARwWAw5LjN63N6vR5paWlIT0/P8aHRaODo6AidTpdt6+TkBFdXV7i6usLFxQVN
	mzaFs7MzgOxuTrvliy842qx3b6B5c6WtyZGmTYHjx/nrun07r3X36gX88gtQtKjS1r2Yo0ePonfv
	3tBoNGjcuDGCgoIQFBSEqlWr2uQanVckIvUtROzatQuzZ8/G6tWrkZGRAQcHB3Tp0gV9+/ZF0aJF
	n7ogPHlhyOnnjIwMJCUlISkpCcnJycb9Pn36oGLFisaxb9zgSLZ//gFOnDDZ9Oqr/D3q3BnQZFkZ
	NRgMWLBgAcaMGYN79+7l6Xzd3d2RkpKCjIwMVK5cGb6+vnBzczM+3N3d4eLiAq1WC61WC41GY9zX
	arUICgpC7dq1TQccPRr43/94NX3/fiDzuZSUFJw4cQJOTk7w8PCAu7s7PDw8jOKrBq5dA8qV4zW/
	e/f4FGxFeno64uPjER0dbXzExMRk+1n+XUxMDGJjYxEXF4fY2FiUL18e+/bt4wNdvQqULw/4+vJJ
	SBKSk5Ph6emJ9PR0252QGTRq1AihoaEAgMePgRIlOAj4zh3eB4Bu3brh+PHjcHFxySaYT+4/+Ttn
	Z2dIkvTcx/OE/kmxlyQJTk5OcHZ2zvaoU6cO/Pz82Njdu3k9z9UVCA8HAgMB8P+YiODg4ACDwYCM
	jAzo9fqntvJ+TtcZ+QEAzs7OcHFxMT50Ol2e/v5EwO+/8zUmKYn/5tOnm7++mpqaiqioKERGRhq3
	8n5MTAwSEhKQmJiIhIQEJCUlgYiM/wOAb7ydnZ2Nf98nt0/+7sGDB9i7dy+uXr0KvV5vtMPPzw+v
	vPIKWrVqhZ49e5r+LypBlcKn1WqtE9afA8HBwfgjM4b++HGgfn3Tc56efLP4wQc5R7X9999/CAkJ
	waNHj+Dh4QGdTpfjHXdKSgri4+MRFxf31DYxMTFf9s+aNQsffvhh1l9wqKmDAxASYgy2SE9PR5s2
	bUwX5ydwdXU1iuGTD1l8sz6cnJzg4OAABwcHaLVa4778M2CKUH3WRUPelyQJXbt2hZOTEwBTqPfC
	hTwDBIDTp0/j+PHjzz1eenp6thubZ22f9VxGPiIMdDodYmNj4eLiwlcxPz/gwQPgyhVWcgBDhgxB
	REQEJEmCRqPJts3pd7l9TqvVGmdyWR+Ojo7GC3x6erpxVpiWloa0tDSkpqYaz33gwIF48803AXBU
	8vvvA23a8CwEAKKiohAYGGiz76W5VKlSBWFhYXwDR8QpO0eOcGDX2LEA+BwqVaqEhIQEq9nh4OAA
	FxcXuLu7Z/Ps5OTtqVq1Ktq3b8+fmUwuXzY5awD+LsycCdSrl32cjIwMnD59Gjt27MDu3btx584d
	REZG4vHjx1Y7t7zy7bffYty4cUqbkQ1VCl/ZsmVx69atbF8ynU4Hb29vFC1aFF5eXtDpdNnuGLNe
	CHL62cHBwXgXmvWu1MfHBx999BE0Gg30eqBKFY78ffttToXJ9PxYDb1ej4SEBNStWxfXrl3D6tWr
	4ePjY7wrS0xMRGJiIpKTk413onq9HgaDAR06dECjRo1MB5s/n/P1AI7gHDjQ+NRPP/2ErVu3IjEx
	EfHx8UhISDBurXkhyC2//fYbPvroIwDAr79yAF6XLsC6dfx8eHg4qlSpYlUbtFot3N3d4e3t/dTD
	y8sr276Xlxc8PT2zPdzd3U3unTfeADZuBBYv5rsnO6JVK54wzZ8PDBjAv0tPT0dMTMwLbyRy2k9J
	SXnmrEl+aDSaHIU+J7EnIqSlpSElJcX4GDFiBFq0aMHG/vsvf3l9fdmFkBnJHBwcjEWLFiE9PR0Z
	GRlGj4l8s5Z1X94+eR3J+gDYi5KcnGx8ZJ315AZfX1/88ccfxpsOgFcq/v4bGDWK750Avh59+y1f
	n57kzp07WLduHbZt24b9+/fD3d0d/v7+KFmyJEqWLAl/f3/4+/ujWLFixptYNzc3uLq6Gv+eAIye
	sdTUVKSmpiIlJSXbNqffZd0mJyfjwoULOHnypPGGvkiRIliyZAk6d+5s1t/F2qhS+ACesu/Zswcb
	N27Ehg0bcOvWLeNzzZs3x549e6wyrsGQ3ZVpK8qVK4dr167hypUrKJc5QzCLRYuAfv34bveHH4AR
	I3L9VoPBgKSkJKMYyqIrP578OSEhAWlpacjIyHjmA0C2C8SzLhzyfrVq1TB+/HgAQGQkL006OfEX
	X3Z3jh07FpGRkc88Xtabmyfdbs/aZt3Pq5sqR6ZMAcaM4TyXX36x3HGtTGQkewUdHdlLa24+quIQ
	AQ0asPtmxgzOFbAhstchISEhRw/Ps/YbNWqEYcOGwSOLbz86mtNuf/2VlyolCejWjU+pSZOcXaCy
	B8WWEBFCQkLwf//3fzh79iwAoGrVqpg8eTK6du2qzrW+fATG2AyDwUCnTp2iiRMnUsOGDWnatGlK
	m2RxSpUqRQDoel4ani1dampSOXWq5Y1TADkgddkypS3JI1u28Am89prSlpjFDz+w2d26KW1JHpG7
	j5QowTkZBYBbt7gNn9zfT24R9eOPRHfvKmvbsWPHqEmTJsaI+1KlStGCBQsoQ+VR5HYhfE9iy1QG
	W+Hv708A6La5xWRnzeLakADRN99YxzgFmDmTT+mtt5S2JI/cusUnULSoXVV3rl+fzba3sHoi4mrQ
	tWrxCfz0k9LWWJzISG78UbSoSQC1Ws4BnDKFOz/YOv91z549BIB8fHxo5syZlJycbFsD8ohqXZ2F
	DV9fXzx8+BD37t1D8eLFX/wGIvaDyIvGU6ZwNKca3Qp54OZNoEwZDsq7f1+1BWeeDRHXYIuNBaKi
	TKGRKiYiAqhUiV3L9+5xZK1d8c8/HAYZGMgnY+0FeoVITjYtH2/enL3qi7MzUKMGB3K/9BIvGQQG
	8v+1dGnzx6IswWNZI7/j4uIwdepUvPvuuzh58iS6deuGIkWK5PfUbIeyuiuQ8fT0JAD0+PHjF784
	I8PU8E+SiP780/oGKkCjRnyKq1YpbUkeadKET8DKzU4txTffsLl9+yptSR5IS+PqLADRnDlKW2Mz
	Hj3i2fn77xOVK2eaCT75GDQod8czGAzPdFPq9XqKy+zvN3PmTJIkiUaNGkUpdtj0Vh3JWwJjftcL
	AyxiYjiZ8KefAJ0OWLECGDLEBhbanrfe4u2qVcrakWeqV+dtWJiyduQCIjtPWl+wgFNHKlXifIBC
	QtGi/D2ZPZtPPzqa6yb8+itHhfbuDbRoYUzlzYaco5gVOT0GACIiIjB//nz06dMHlSpVgoODA5Zm
	fkjkIDMXFxc4OTnlKxVICVRduaUwkSvhO30a6NmTE3KLFQNWr1ZtRQpL0L07J/Ru3MjuHbtzvdmR
	8J0+zR8rX1+gdWulrTGT5GTgm294f+JEm7ccUhNeXpz716zZi1+bU9GKhw8fYsKECZgzZw4yMjLg
	6emJypUro3nz5pgwYQK6d+8OACiR6bq/cuXKM4+lZgrvJ+RZ/PorZ7E3bGjTvIaMjAxUqVIFjo6O
	Tz9pMHA7iFGjuNp8zZrA+vXsxC/AvPQS/yuOHeO1jCypTvaBLHznzilrRy6QZ9Xdu9uhbsyaxXkY
	deqY3ASCbMg50RqNBmlpaTh8+DB27tyJx48fo3fv3mic2aNz48aN+P3331G5cmXMnDkTvr6+8PPz
	g5eXF9zc3IypCRUqVAAAY9EJexM+scaXlchIk1O8eHGigQM5PNrKVd31ej2dO3cu5yePHSN65RWT
	XR98UGDCtHPDd9/xaffsqbQleeDuXTbew0PVkZ0GA1Hlymzq9u1KW2MmsbFc3Rkg+vdfpa1RDDnS
	/fbt27R7926Kymz1oNfrs73u6NGjFBAQQFqtlqpWrUqVKlWigIAAGjp0KKWkpNDy5ctJkiRasGDB
	c8fbsGEDubi4ULt27ejRo0fZbLAHhPBl5dYtok8+ISpbNvvKsLMzUadOHERibrpBXgkLI+rXz5Sq
	UKIE0dq1thlbRVy7xqfv6mqnel+8OJ/A1atKW/JMzp0zdQdIS1PaGjMZP56Nf/VVVd9cWBNZ3CIi
	IsjJyYkkSaL33nsv2/MnTpyglStXUrVq1ah69eq0ceNGio2NpejoaGrTpg1ptVrav38/EREVLVqU
	goOD6cSJE7Ro0SLq27cvtW/fnsLDw43H3LJlCwUEBFC9evXo8uXLRESUkZFhN+InhC8nDAaiM2eI
	Jk0yhRZmfVSqxBmly5ZZNoM0IYFDtDp1Mo3l4EA0ciTf2RZSGja049yytm3Z+HXrlLbkmcjRnAMH
	Km2Jmdy/T+Tuzsbv26e0NYrTo0cPKlKkCHXs2JEkSaKPP/6Y0jMT+z7//HNq1aoVbX9iSr9w4ULq
	2rUr7dq1y/i7zp07kyRJ5OXlRSVKlKCWLVvSsmXLKDU11RjxuW/fPvLw8KD27dvTw4cPsx0zOjqa
	wsLCVC2CQvhyQ1QU0dy5RG+8YfqiZX289BJnWk+dyrOyc+eIXpTIaTAQPXjAoe5TpxJ16ULk4pJ9
	lvnhh0SZd1OFmenT+U/So4fSluSBkSNVX1ygXj02cdMmpS0xkxEj2PDXX1faEqsSERFBgwYNMs6s
	ciI8PJzc3Nzojz/+ICKiIUOGkCRJtHDhQiIiOnToELVo0YIOHjxIRESXLl2icePGkSRJNGTIEEpP
	TzcK1TvvvEN16tSh5ORkSkpKovj4+KfGO3fuHHl4eJCHhwcFBwfTqlWrqH///lStWjWSJIk8PDzo
	5s2blv5TWAwhfOaSlkZ0+DCXSmjTJrtYPfnw8uLcogYNiBo35tljgwYslHKDzCcfr7xC9P33fDcr
	ICKi69dN7k4VNdHOHQsXsvHduyttSY7cv8/mOTkRJSUpbY0Z3Lxp+g6dOKG0NValb9++JEkS/fHH
	H0/NouSft2/fTr6+vrRkyRIiIgoLC6MGDRpQvXr16Pz580RENHToUJoyZQoREU2ePJm8vb2pS5cu
	1Lp1a6pTpw6dPXuW9Ho9DRw4kCpWrPiUHQaDwdhwNiUlhRo1akSSJBkfxYoVoy5dutA///xDa9as
	oZiYGKv9TfKLEL78kp7ObtEFC3h9sEMHogoVuJbQswRRfri7sxB+9BG/X8V3SEojx/esWKG0JWZy
	+jQbXqGC0pbkyNKlbF7btkpbYiaDB9utG8BgMGSbYT0L2a0YFBREkiTRykxff9YEc3l9b926deTg
	4EDTp083PrdlyxYqW7YszZo1i4iIZsyYYVz7i46Opps3b1JcXBxFRETQ4MGD6a3M+oBTpkwhBwcH
	iomJodOnT9OUKVOobdu2VLRoUerQoYMxif3SpUvUpUsX8vb2ptatW9PmzZvp/v37qnZxythb4LL6
	cHDg9IKaNbMnzur1nGz+6BF39jQYuJyYRsNZp/7+gLu7YmbbGz16AKGhnK//9ttKW2MGVapwoYEr
	V4CEBNX9z7dt421m20b74NIlTljXajlvz86QO4m8CK1Wi/T0dJQqVQoAJ5Q/Cy8vLwAwlg3bv38/
	xo4di/v372Pfvn348MMPUbVqVWzatMn4evk9Hh4e8Pf3N7aPSklJAQCUKVMG6enpSE5ORosWLTBy
	5Ej0798fHh4eyMjIQMWKFbF06VK4urrm/Y+hEEL4rIVWy0nmxYqZ9TZSoK2IPfDWW9yOJSRElfrx
	bBwdgapVgTNnOJE9a/9EhSGyU+EbP55vLAcN4kotdkZMTAyWL18OvV6PwYMH55y7m4lWqzX2oZR7
	/e1UwOMAACAASURBVGXNmZP3HR0dodfrsXbtWgQGBuKrr75C5cqVjWNNmzYN5cuXhyRJuHHjBsqU
	KYP4+HicPXsWf/zxBxYvXozVq1dDkiSUKFECRYoUQaVKldC5c2eULFkSLVu2ROnSpaHRaEDEHeyJ
	yCh6er3e2LvQLlB2wimQGT16NDVu3Dh3tToLKY0b22mrovfeY8Nnz1bakmzIaQx+fnaUCXD2LBvt
	6GgXSwNPujTT09Pp66+/NgaAREdHP/f9BoOBfvnlF2OKwpN5efKxz507RwEBASRJErm6ulJQUBDF
	xsbS8uXLycXFhcaNG0c3btygt956i7744guaN28eDRkyhEqXLk2SJNHIkSONx7x//z6dPn2aEu0y
	fyh3iBmfSvj7778RFRWFpKQkeHt7K22OKnn7beDwYWDlSq7cZjfUqsXbM2eUteMJss727OVGHd9+
	y9shQ4BMF6BaiYmJwS+//IIyZcqgb9++0Ov1OHr0KL7JLK/Wq1cvuD/HdUGZ3h9PT0+4uLggMTER
	cXFx8PLygsFgyDbDKlGiBMqVKwdvb2/8+eefaNKkCQCgSZMmeOWVV7B161YMGDAAPj4+mD59Ohwd
	HVGuXDm0bdsWHTt2RLssU35fX1/4+vo+ZUdBQgifSpBdBklJSQpbol7s1t1pB8JnF5w7x3c9jo7c
	gkvlpKenY8KECQAAT09PdOnSBZ9//jkAoEWLFvjyyy+fu9aXVXCSk5Ph4+NjfL1Go0FqairCw8OR
	lJSEBg0aoEKFCjh//jyaNGlidIsGBgZi2LBhGDJkCFJSUjBq1Ch88MEHqJ1T1epnUNBEDxDCpxrc
	MhvOJSYmKmyJeilVCmjcmGd9ISF2NOvLKnxEqphepaRwFX8AaNNGWVtyzeTJvB0yhBvNqRxfX18E
	Bwdj7ty56N+/P+rUqYOTJ09Cp9NhzJgxKFeu3HPfT5mdE4oWLQpJkrBy5UrodDoEBgYiJCQE58+f
	R3R0NOrWrYsdO3agfPnyOH/+vPH9cpeFjh07YsiQIXBzc0PpLE355PqdkiQVSHF7HkL4VIIsfGLG
	93x69LBDd6efH7c9ePAAuHUrbx1BLczBg9zUoHZtu+iRC9y+zf90rRb48kulrck1v/76K1xdXTFz
	5kzszbzTeOedd9C6dWuju/JZyM/VrFkT1apVQ1hYGGbNmgUAqFWrFvr27Ytu3bpBo9HAzc0NMTEx
	uH//PlJSUuCcpQmvk5MTJss3DTkcvzBSeM9cZciuTjHjez5y8X3Z3WkXSJLq3J125+acNYsjObt3
	V/3aXla0Wi1GjRqFYcOGGX938uRJhIaGQqPRGF2SOSHPwkqXLo0NGzagXbt2qFGjBj755BMsWbIE
	U6ZMQbNmzfDqq6/C0dERH330EQ4fPpxN9AQ5I4RPJXh4eAAA4uLiFLZE3QQGAk2asKsuJERpa8xA
	CF/eSU7mTqsAMHy4srbkAU9PTyQnJxt/Dg8PR58+fbBkyRKjO/J5EBFeeukl/Pvvvzhz5gx+/PFH
	VK9eHS5ZGlTKr/Hz87PKORQ0hPCpBB8fHwDcCFLwfOQE9pUrlbXDLFQkfPfuAadOAc7OwKuvKm1N
	LggJ4UIQdevyIq+dsWvXLsydOxcBAQE4ffo0goKCcPXqVbz33nuYPn268WZXXtN7EkmSQETGHDp5
	be7J1whyjxA+lSCHDz948EBhS9SP3JB261buy2sXyMJ38qSydgDYsYO3zZuz+KmeFSt427u3KgKD
	zGXs2LHGbc2aNbF+/XpjtOcPP/zw3IosMrKwSZJUqNfmLIUIblEJQvhyT+nSQI0aHN1+4ADQqpXS
	FuWCGjU4DP/SJSAuDsgsLaUEduXmTEgAMstsoUcPZW3JA8nJyahWrRqqV6+OPn36GH8/ZswY1KtX
	DykpKahatSoAMWuzJUL4VIIQPvPo2JGF799/7UT4HB151nfsGHDiBNCihSJm2F2Zsk2beI2vSRO7
	CmqRcXFxwbx582AwGODs7GzMzXNwcEBQUJDS5hVaxJxZJQjhM4+OHXn777/K2mEWL7/M2+PHFTPh
	3Dng7l2ukV69umJm5J7ly3lrN7krT+Po6GiMtBSzOnUghE8lCOEzjyZN2Ft44QJw/brS1uQSWfiO
	HlXMBDlpvVUrO1gui4sDNm9mQ+U8FoHAAgjhUwlC+MxDpwPatuX9zZuVtSXXNG3K2z172OeoAPv2
	8bZ5c0WGN48NG4DUVOC114CSJZW2RlCAEMKnErIK37PCmgXZkd2ddpPPV7UqX8Dv3WOfo40hMglf
	s2Y2H958CoCbU6BOhPCpBGdnZxQpUgTp6emIjY1V2hy74PXXebtzJ2AXld4kyVQYc/t2mw8fEcGa
	W7y4HbSxi47mfBWNhqu1CAQWRAifipBnfffv31fYEvvA3x9o0ICruMi5aapH9s/KIfo25MAB3r72
	mh2s761bB6Snc/SrqEYisDBC+FRE8eLFAQjhM4c33uDthg3K2pFrOnXiBcq9ewEb/5/l3PmGDW06
	bN5YtIi377yjrB2CAokQPhUhhM98ZOHbtAnIoZKT+vD25gQ6gwFYs8amQ8vCV6eOTYc1nytXgN27
	ARcXu0xaF6gfIXwqQgif+dSsyZVc7t1TNEvAPOSL+bJlNhvSYABOn+Z91Qvf/Pm8ffttwNNTWVsE
	BRIhfCpCCJ/5SJJp1rdunbK25JouXbhI5t69wLVrNhnyyhWu/lWyJAe3qJaMDOCvv3h/8GBFTREU
	XITwqQjPzLtb0ZrIPOSgv2XLFEuPMw9PT5PRCxbYZMhTp3ir+tneli1AZCSHndpF6wiBPSKET0XI
	Pfni4+MVtsS+aNaM+/Rdvw4cOqS0Nblk0CDeLljADVatjCx8detafaj8MW8ebwcOtIPQU4G9IoRP
	Rbi7uwMAEuymtbg60GiAd9/l/SVLlLUl1zRvDpQrB9y+bZOcvhMneFu7ttWHyjt373KUklYL9Oun
	tDWCAowQPhUhz/iE8JlP7968XbGC079Uj0bDsxrANMuxEkSmutj161t1qPyxcCGv8XXuDJQoobQ1
	ggKMED4V4ebmBkAIX16oVYu7DTx6xAU/7IL+/VkA168HrFij9fZtPry3N1C2rNWGyR9EphsA2Q0s
	EFgJIXwqQqvVAgAMdpGQpi4kyTTrW7xYWVtyTUAA0KEDT1H/+cdqw8izvZdfVvGy2cGD3KTX399U
	i04gsBJC+AQFhl69eLt+PXe0sQsGDOCtHMJvBY4d462q3Zxz5/K2f3/AQfTHFlgXIXyCAkOZMhzh
	mZICrF6ttDW5JCiIfZAnT5oyzC3M/v28VW2psthYYOVK3pfXPQUCKyKET1CgeO893tqNu9PJyTRV
	tcKsLyEBOHyYlxJbtLD44S3DsmXcXqNFC6BCBaWtERQChPCpiPTMcESNRvxb8spbb7GW7N7NQR12
	Qf/+vF2yhKMaLcj+/byEWL8+TyxViQhqEdgYcYVVEY8fPwYAFCtWTGFL7BcvL/YeEgFLlyptTS55
	+WWuVPLggalTrIWQu9PLbQBVx5kzXGQ1azUbgcDKCOFTEQ8fPgQA+Pj4KGyJfSO7OxctspMSZpJk
	uuhbcHEyI8PUxLxbN4sd1rLIs73evbkbg0BgA4TwqQhZ+OSGtIK80aEDUKwYcO4cTyjsAln41q61
	WH+lnTu55V+lSjypVB2pqabFWOHmFNgQIXwqQsz4LIOjI9CzJ+/L/UxVT716nF0eFcXRKBZALt/W
	p49K8/fWrQMeP+bK2fXqKW2NoBAhhE9FCOGzHHKg5OrVdujutECD2sRE02Hkv4XqkHP3RPshgY0R
	wqciHmSWrRLCl38aNwb8/Lhjg5XS4yxPly683bAh32q9YQOL3yuvAOXLW8A2S3P9OrBjR/Z0DoHA
	RgjhUxFixmc5NBqga1fet8AEyjY0bgwULQpcvgyEh+frUPLSWZ8+FrDLGsh9CLt3V3GehaCgIoRP
	RQjhsyxyJOPatcrakWscHIBOnXh/48Y8H+bBAy7UrdUCPXpYyDZLotcD8+fzvnBzChRACJ9KICKj
	8Ik8PsvQsiWnh507B0REKG1NLgkK4m0+hG/5ctaWDh0AVQYIb9/O1QXKleO+hAKBjRHCpxISEhKQ
	lpYGV1dXuLq6Km1OgcDRkVu7AXY062vfHtDpuFvBo0d5OoQczSl3q1AdclDLwIHskxYIbIz41KkE
	4ea0Dnbn7ixShGtWGgzAv/+a/fYbN4DQUMDVFXjjDcubl28ePODIG43GVKpNILAxQvhUghA+69Ch
	A+DszGIQFaW0NblEdndu2GD2W1et4m3nzkBmX2N1sXAhFw99/XXuRygQKIAQPpUghM86uLkB7drx
	/vr1ytqSa2Th27oVSEsz661yd5+337awTZaACJgzh/dFUItAQYTwqQQhfNZDTmuwG3dn2bJAzZpA
	fDywd2+u33bz/9u787CoyvYP4N9nhl12RZTQwhRN3HkVks0URRRNvbLMxOWXS1lu1ZtbKWYueIWG
	W5bhlhaauaQgLkG+uaQV+VouYYL6goooKC4swty/Px5nYGTVgHOQ+3Ndc81h5uGc+4zIzbNfAo4d
	k82cffpUX3iP7dAhOU2jceOi0auMKYATn0pw4qs+/frJLqX4eODmTaWjqaTHaO7Uz1fs00cmP9X5
	4gv5PGqUHMDDmEI48akEJ77q06CB3Jm9oACIiVE6mkrSj0zZtavSq7jExcln/QIwqpKVVdQByQtS
	M4Vx4lMJ/XJlvDND9ah1ozs7d5Zrrl28KCciViA3t2grv549qzm2x7FxowwyMFDO32NMQZz4VIJr
	fNVLXwuKiwNycpSNpVI0mqJJiJVo7jxyRN5Xu3YyX6pK8UEtY8cqGwtj4MSnGpz4qtfTT8udb+7e
	lWsj1wqPsIqL/p5UWds7fhz44w+5jIwq22FZXcOJTyV4ubLqN2iQfNYP+Ve9wEC5e8GxY8DVq+UW
	3b9fPqsy8ekHtYwYIZfTYUxhnPhU4saD5am4xld99As279ghu5tUr149mfyAckflZGYCv/0mc4qf
	Xw3FVlk3bwLR0fKY5+4xleDEpwJEhMzMTACAA2/RUm1atJDNnbdvA3v2KB1NJVWiuTMhQXaj+fio
	cBrDl18C9+4BPXoALVsqHQ1jADjxqcKdO3dQUFAAKysrWFhYKB3OE+2VV+Tz5s3KxlFp+gEu+/aV
	OSpH37/Xo0cNxVRZBQXAsmXyeMoUZWNhrBhOfCqgr+1x/1710zd37tolB7qo3lNPAZ6eMunFx5da
	5Icf5LPqEt+2bXI5GXd3uTYnYyrBiU8F9InP0dFR4UiefM88A3h7y9a33buVjqaSymnuvHRJ7jVo
	awv86181HFd5iIAlS+TxpEm8/RBTFf5pVAFOfDWr1jV3lrOKi762162b3MBdNRIS5JYYDg5yNCdj
	KsKJTwXuPmhzs7a2VjiSumHwYEAIud1ddrbS0VRChw6Aqytw+TKQmGj0ln6ZMlU1cxIBs2bJ43ff
	Ven+SKwu48SnAvkPtp4x4zlONeKpp+Sw/7y8WrJVkRClruJSWFi0V62+NVQV9u+XO8g7OgITJyod
	DWMlcOJTgby8PACc+GpSrW7ufKCwUI556dgRcHNTKK6HFa/tvf8+YGOjbDyMlYITnwpwja/mvfSS
	HG+xd6+cAK56L7wgmwx//x1ITQUgJ6w//3zRijSqEBsrV5pxcgLeflvpaBgrFSc+FdAnPlPeo6zG
	NGwIdO8up5rVih0bLCyK1iMrVuvr318mcVUgAmbPlsfTpnHfHlMtTnwqoE949+/fVziSumXIEPlc
	m5s7X34ZaNVKoXgetn27XDutUSPgjTeUjoaxMnHiUwGrB+tM5dSK/XKeHAMHyikAP/wAXLumdDSV
	0LevHOgSHw/cuQNA7jqhCoWFwIcfyuMPPlDh2mmMFeHEpwKWlpYAgHv37ikcSd3i6Aj06gXodMB3
	3ykdTSU0bChn3+flqW/2/TffAKdPy0w8ZozS0TBWLk58KsA1PuXoR3fWin4+AHj1Vfm8caOycRR3
	/35R397s2bz1EFM9TnwqoE98XOOrecHBsvXwP/+Ry5ip3iuvAFqtnLmekaF0NNKaNUBystx9ITRU
	6WgYqxAnPhXQN3Vyja/mOTnJNaDz8oCDB5WOphIaNgR695Z9amqo9eXmAnPnyuOPPlLZummMlY4T
	nwpwjU9ZvXvLZ/3yX6qn39B1xQrZQamkzz4D0tKA9u1VNK+CsfJx4lMBrvEpq9Ylvn795CCS8+eL
	1ixTwt27wMKF8njuXN6BgdUa/JOqApz4lOXlBdjZAUlJsqtKjXJzi23Hp9UWrYry6aeKxYTPPpPz
	QDp3LlpLlLFagBOfCpibmwMoWrOT1SwTk6JFUfbuVTaWsmzcCHzySbEXXn8dsLaWkxAPHar5gO7c
	AcLD5fFHH8kRQozVEpz4VECf+HJzcxWOpO5Sc3MnEbB8uYzNsI2SgwMwZYo8/uCDEvv0VbsVK4Dr
	1+W8wqCgmr02Y/+QIKrp/zHsYTqdDlqtFgBQWFgIDfeV1LjUVKBJE1mJunFDXVPRDh8GfH3lCNS0
	NMCwpOvNm0CzZkBWltwKKDCwZgK6fVtuZZ+ZCezbV1RdZqyW4N+wKqDRaAzrdeoXrGY1y9UV8PCQ
	LXhHjigdjbEVK+Tz6NHFkh4A2NsD//63PJ4yRU4krwkLFsik5+NTc8mWsSrEiU8lLCwsAHA/n5L0
	zZ07digbR3FXrwJbt8oBk6Wu+zx5sqz1/flnzQx0OX8eiIiQxxER3LfHaiVOfCrBA1yUp9+tYdMm
	QC0V79WrZUWuf3+gadNSClhaFlUJw8KAS5eqN6B335UfzvDhcjgsY7UQJz6V4MSnPE9P2dx5/ToQ
	E6N0NHKvwM8/l8dvvVVOwd695f5E9+7JBaKra1L7jh3Azp2yI1Q/f4+xWogTn0pw4lOeEMCoUfJ4
	zRplYwFkjklLk0tg9uhRQeFPPwXq15eDTaqjyfPaNWDsWHk8fz7QuHHVX4OxGsKJTyX0iY8Htygr
	NFSO6IyJkd1ZSlq+XD6/9VYlutIaNy7K1tOmVe0IncJC2bSZkSG3rS+3+smY+nHiUwmzB+Pnucan
	rIYN5c4/RMCyZcrFceoU8OOPQL16MudUSv/+wKRJslNwwADg4sWqCebDD+XM/gYNgHXreGkyVuvx
	T7BK6BMf1/iUN2mSfF6zptiE8Rq2cqV8Dg2Vy6lV2iefyN11MzJk39+VK/8skM8/l9MXtFpgyxY5
	2ZGxWo4Tn0pwU6d6dOwIBATIedpr19b89bOzgQ0b5PEjtyqamACbNwNt2gBnz8obSU19vEA2bADe
	fFMeL1sGvPDC452HMZXhxKcSXONTl8mT5fOSJTU3L1xvwwY5kT4gQOavR2ZvDyQkAB06AOfOAc8/
	Dxw/XvnvLyyUzZsjRsg23/DwogTI2BOAE59K8KhOdenXT46mvHhRzuurKURFzZz/aAxJgwZyAWsf
	H1nj69pVrvJy40b535ecDPTtC3z8sezLi4gA3n//HwRSu+h0OgwdOhQajQaj9fseliM9PR09e/bE
	uXPnaiC6qpGYmIg5c+ZgrRLNGWpBTBVefPFFAkDbt29XOhT2wPr1RACRuztRQUHNXPPAAXnNxo2J
	8vOr4IR5eURTphAJIU9sZUU0bBjR118TnThBlJJCdPIk0caNRC+/TGRqKss5OhLt318FAShv27Zt
	VL9+fRJCGB7t2rWjv//+u0TZ8ePHkxCC7OzsSAhB8fHxhve2bNlC7733nuFrnU5Hvr6+JISg+fPn
	ExFRixYtjK6jf7z22muUkpJCRERJSUnUtWvXUssJIWjy5MlERJSWlkZ+fn5G75mbm9PGjRsf63PI
	zc2lV199lTQaDQUHB9OJEyeIiCgxMZE0Gg0tXLiQ7t+/T6ampjRjxowS3z979mwSQtCGDRuIiGjN
	mjXUsGHDUu9Bq9VSYmLiY8VZEzjxqcTgwYMJAG3evFnpUNgD+flEbm4yD0RH18w1Q0Lk9ebOreIT
	Hz9OFBwsT17eQ6MhCg0lSk2t4gCU8csvv5C5uTm5urrSxIkTKSoqiqZPn052dnZka2tLZ86cMZT9
	+eefSQhBkZGRlJGRQW5ubtS6dWvS6XRERDRp0iQSQtDUqVOJiGjZsmVkZmZGWq2WPD09iYhICEEa
	jYaGDBlCUVFRFBUVRWPGjCEhBHl7exMRUXBwMAkhqHfv3oYy+sfatWvp7t27VFBQQF5eXmRiYkL9
	+/en1atX04oVK6ht27YkhKB58+Y90ueQk5NDHh4eZGNjQ6tXrzbcExHR4sWLSQhBo0aNotzcXBJC
	kJWVFR06dMjoHPrEd/DgQbp69SoJIcjExITCwsJK3MfevXsf/R+rBnHiU4nQ0FACQOvXr1c6FFbM
	F1/IfNCmDVFhYfVeKylJXsvcnOjatWq6yF9/EYWHE/XrR9SqFVGTJvL5xRfl6//7X5Vf8o8/iH79
	leiXX4iOHiWKjSVatozo/feJwsKIdu8munevyi9LRESnT58mc3NzmjNnjtHrmZmZ5OTkRCEhIYbX
	lixZQubm5pSVlUVEMrEJISg2NpaIihKfs7MzJSYmkouLCy1evJgWLlxIQggikokvODi4RBytW7cm
	Z2dnIiJatGgRmZqaUl5eXrmx+/v7k5ubm9FrOp2OxowZQ/b29qXWWEuTm5tLwcHBZGZmRsnJySXe
	P336dInEJ4SgNm3aGJXTJ77jx48TEVGbNm1o+PDhlYpBbTjxqcSYMWMIAK1atUrpUFgxublErq4y
	IVV3K/Tbb8vrvP569V6nprVoUXFF09JStshWh/fee4/CwsJKvB4bG0tCCLp58yYREQUGBlL9+vUN
	76ekpJCjoyO9+eabREQUERFh1JzXqVMnunfvHq1bt440Gg0RycS3cuXKUq+zaNEiIiL67LPPSKPR
	0LFjx+jKlSuGx/37942+788//6RnnnmmRNx5eXnUpEkTWrJkSaXuXx/3tGnTSn1//vz5hsR3584d
	o3tcs2aNodzMmTPJ1dXV8LWXlxf5+flRWlqa4R4yMjIqFZPSeHCLSlhaWgIAcnJyFI6EFWduDkyd
	Ko9nzpQDHqtDVpacGw4UzSN8Unh4AJ06yUeXLnIno9GjgXnz5Gfr6Qnk5ABWVtVzfWtr61JfDw4O
	hrOzM+Li4pCTk4PTp0+jR7G14ezs7ODj44MzZ84AADp16gQAmDRpEszNzbFq1SpYWlrCx8cHVGxb
	08jISBw7dgxXrlzB0qVLMWjQIPj6+uLfD7aQys/PBxEhMDAQTz31FFxcXODu7o6NGzcaxVevXr1S
	4zYzM8Nrr72GuErumuzr6wtXV1eEh4ejb9++JQbi6AfU2draYteuXQCAo0ePIigoCDNmzMDly5cB
	ANu2bTNsn6a/jyNHjqBly5ZwcXGBi4sLfH19DeXVzETpAJikT3y8C7v6jBkDLF4MnD4NrF8P/N//
	Vf01IiPlFIbAQKBt26o/v5K2b6+4zNWr1Xd9nU5X5ubOXl5eaNq0KdLT03HlyhWcPXsWaWlp2LRp
	E2bMmAGdTgcPDw+j75kyZQpCQ0MNifDhcyclJeH55583fO3n54eYYquef//992jevDmSkpJw6tQp
	pKeno1u3biXOUzyZPqxLly64UdEI3WJlL126hISEBEyfPh0eHh6IiopCaGioUbkBAwbgwoULhs9l
	2bJl8PT0RFBQEL755hvcuHHD8EdEamoqTpw4gY8++ggzZ85EQkICnJ2dS3xWasU1PpXQ78fHNT71
	MTeXo/sBYNYsuQlCVbp5s2hd6VmzqvbctUWjRvJRHfbt2wdRymKn9+7dw549e9Cm2GTJP/74A02a
	NMHHH3+MLVu2wMfHB6dOncK1a9eMvlef9ADgp59+MnpPo9FgwoQJiIuLw759+7Bv3z6jWmdBQQF8
	fX0BAB4eHujevXupiXnfvn1l3tPWrVvR9hH/QnrhhRdw5MgRLF++HOPHj0diYiIAGBLowzE0b94c
	u3btQnJyMtq1a4eMjAx0797dcA8A4O/vDyEEunfvXmuSHsCJTzW4qVPdhgyRK7qkpckVvKrSkiXA
	rVtyBwY/v6o9NwMaNmxY6us7duyAl5cXbGxsDK8NHz4c8fHxuHXrFgYNGoRx48YBAI6Us+j3wzWz
	oKAgREZGolevXggMDDTM0QWAwsJCJCcnIysrC0SEzMxMXL161fC4X2y1BCcnp1Kvl5+fj927dyM4
	OLjim3+IRqPBsGHDYGpqii+++AJAycRdnL+/P6ZNmwYAEEJg8ODBAGBoLr1+/ToKCwuN7iEjI+OR
	46pp3NSpEpz41E2jAZYulYkpPBx47TWgVat/ft7UVLm8JiD3kWVVr1OnTiWSU05ODiIiIrDgob9i
	Ro4ciW7duhm+dnd3BwBcunQJbm5uZV6jeG1p4MCBZZa7ffs2UlNTkZqaihYtWiA5OdnwnhACsbGx
	CAoKMsRdmsWLF6Nv375o3rx5mdfRS0tLw549exAUFISffvoJd+7cwcqVK5GdnY1+/fohMzMTZ8+e
	hRDCqP+uuJkzZyI1NRX169dHr169AACnTp0CALzyyito0qSJoYkUAJ5++mmcO3cOJibqTS/qjayO
	4cSnfr6+wOuvA1FRcseEQ4fkFkb/xL//LZtOX3pJnp9VPSIyauq8ffs2xo4dC3t7e8Mv8qSkJFhY
	WBj1zQGyf0xfY2zfvj2aN29u+L+q5+3tja+++srwtVk5PxRWVlZwcXFB69at4erqCn9/fzRu3BgB
	AQFo0KBBmckOkINQ4uLisHDhQhw+fLhS956QkICxY8dCo9FA92CDYm9vbxw8eBA+Pj64cOEC8vPz
	IYSAl5cX/vrrrxLn0Gg0+Fy/I/IDzz77LKytrTFw4EBotVoAQNeuXdG0aVO0b99e1UkP4MSnGvo+
	Ph7com6ffALs3w/88ovsj/snG5F/9x0QHQ1YWBTV+ljVys3Nxffffw9nZ2d8+eWXiImJwc6dpHU4
	PwAADptJREFUO9G5c2fEx8cbyp0/fx6WlpZGzZJ6b7zxBpo2bQpAJsiHtWrVCq0eVP9HjhwJLy+v
	MuMxMzNDaiUXDd+8eTOysrLw5ZdfIjk5GatWrUJeXh5iY2Mr3Z/m6+sLc3Nz5OfnIzg42GiQzcM0
	Gk2pfaGl6devH7KV2rqkKig4lYIV8+233xIAGjRokNKhsAocOiQXOAGI1q17vHNcvEjk4CDPsWxZ
	1cbHipw4ccJoXpq1tTWFhYVRdna20qFVqEOHDoa49cuMnTx5skqvkZaWRhqNht555x0iIkpOTqZO
	nTpV6TXUiGt8KqFvLiisrolirMr4+MhRmBMnyqZPR0e5qHVl3bght8rLygKCg3lD8+rk5uaGOXPm
	wN/fHwEBAUqH80jeeecd3Lp1C2+//Xa1XcPFxcXod46bmxt+++23arueWnDiUwl9EwOVM3eHqceE
	CcDly7Kpc9AgYMUKOd+vopaijAwgJAQ4c0ZuObRpU8Xfwx6fra0tPvzwQ6XDeCwPz7NjVYenM6iE
	flQYJ77aY/58OTiloAAYNw7o319Oci8NEbBnD/Cvf8mt8Z55Bti7F3BwqNGQGWPgGp9q6Gt8+pFX
	TP2EABYtAtq1A8aPB3bvlg8fH6B7d+DZZ+USZ0lJMumdPCm/r3NnYMcOwMVF2fgZq6s48akEN3XW
	XsOGycnnc+fKJc0OH5aPhzk5yRri5MlAGVOmGGM1gBOfSnDiq90aN5Y7py9aBMTGAomJwP/+B5iY
	AE2aAN7eQM+ecvkzxpiyuI9PJbiP78lgbQ28/LIc9LJpk6wBfvyxHNDCSU/9dDodhg4dCo1Gg9Gj
	R1dYPj09HT179iyx40Fdk5CQgLCwMGyvzIrkKsCJTyW4j4+x6rF9+3Y0aNAAGo3G8Gjfvj3Onz9f
	ouyECRMQHR0NW1tbrFmzBgkJCYb3vv32W8PWQoD8I/Wll17CDz/8gK1btwKQS5wVv47+MWzYMMOy
	XufOnYOPj0+p5TQaDaZMmQIAuHz5Mvz9/Y3es7CwwKZNmx7rc9i9ezeeffZZaDQatGrVCtu2bStR
	RqfT4b333oOZmZnhmoMHDy7zD/KsrCz06tULPXr0wLlz59ChQwcAwM6dO6HRaBAdHY2UlBRotVrD
	2qDFjRw5EhqNxrBeaHh4OGxtbUv9XKytratuyyMF5xCyYuLi4ggA9ezZU+lQGHti/PLLL2Rubk6u
	rq40ceJEioqKounTp5OdnR3Z2trSmTNnDGV//vlnEkJQZGQkZWRkkJubG7Vu3Zp0Oh0RFe3APnXq
	VCKSO7SbmZmRVqslT09PIiLDZPMhQ4ZQVFQURUVF0ZgxY0gIQd7e3kREFBwcTEII6t27t6GM/rF2
	7Vq6e/cuFRQUkJeXF5mYmFD//v1p9erVtGLFCmrbti0JIWjevHmP9Dls2bKFTExMyN7enqZNm0bh
	4eE0YMCAEuX0m9a2a9eOvvjiC+rSpQsJIWj8+PElyl6+fJlcXV3JycmJdu7cafTexIkTSQhBc+bM
	obNnz5IQgpycnOjs2bNG5UaMGEFCCLp48SIdO3aMhBBUr149ioiIKPHZHD58+JHuuTyc+FRi7969
	BIACAwOVDoWxJ8bp06fJ3Nyc5syZY/R6ZmYmOTk5UUhIiOG1JUuWkLm5OWVlZRGRTGxCCIqNjSWi
	osTn7OxMiYmJ5OLiQosXL6aFCxeSEIKIZOILDg4uEUfr1q3J2dmZiIgWLVpEpqamlJeXV27s/v7+
	5ObmZvSaTqejMWPGkL29Pf3999+V+gwOHjxIWq2WnJ2d6b///W+Z5ZKTk8nExITatm1Lf/75JxER
	FRYW0pAhQ0gIQVFRUYaymZmZ5OnpSfXr16dbt26VOJd+1/niiU8IQX369DEqN2LECNJqtZSenk75
	+flkb29Ps2bNqtR9/RPc1KkS3MfHnljjxsmlbUJCgD595KNfPznno3Nnufvu0KHA2rVVfunnnnsO
	EyZMKPH/ysHBAevXr0dMTAxu3boFAIiJiYG1tTXs7e0BACEhIXBwcDDsSq5fr/PatWvw9PREo0aN
	8MYbb6BRo0ZGa1z2e2gZnz179uDMmTN49913AQA2NjYoLCzEiRMnjLbz0e9xp7dy5coScQshsHz5
	ctjY2BjiKk9eXh5CQ0MhhEB8fDzatWtXZtkFCxagYcOG+O233wxrgWo0GmzatAm9evXC8uXLDWVn
	z56NxMREzJs3D7a2tiXOtWPHDsNx8a2W4uLi8MMPPxi+LigoQNeuXdGwYUOYmprCwsICSUlJRp9L
	VlZWhff5yKo9tbJKOXDgAAGgF154QelQGKtaLVrIRUkrerzxRrVcPiwsjMLCwkp9r1GjRhQdHU33
	7t0jFxcXevnllw3vZWZmUr9+/ahbt25ERJSQkEBCCJo8eTJZWFjQ8ePHiYjo3LlzRjW+li1b0s8/
	/0yXL1+myMhIsrCwID8/P8N5IyMjSQhBNjY2pNFoDMdr1641ii0lJYWeeeaZUuOeNm0aBQUFVXjv
	X3/9NQkh6M033yy3XEpKCpmZmdGBAwdKfV9fazx//jwREe3atYvq169PWq2Whg0bRmlpaUbl9U2Y
	S5YsoQULFpCtrS2dPHmSOnbsSB4eHnT79m0iIrK0tDT8ztPpdOTo6EimpqZkZWVFQgjSarXUuXNn
	unPnToX3+ih4OoOCiAi//vorOnfuXO50hkuXLsHc3BzOzs41HSJj/9yqVcDdu/JYv29dYSFgaQnY
	2QGZmXIB00rsL/c4dDpdqTucA4CXlxeaNm2K9PR0XLlyBWfPnkVaWho2bdqEGTNmQKfTldgJYcqU
	KQgNDTVsIfTwuZOSkoy2N/Lz8zPaFeH7779H8+bNkZSUhFOnTiE9PR3dunUrcZ7SfhfodenSxbBz
	enny8vIAAKNGjSq33MaNG+Hl5YUePXqU+r6trS10Oh3OnDmDZs2aISQkBNevX8e2bdswc+ZMw2CZ
	wMBAo+8bMGAA1q1bB0dHR7Rt2xZLly5F9+7d8eKLL2Lp0qXIzc01/O47fPgwsrKysGHDBgwePBg/
	/fQTmjVrhmbNmlV4n4+sStMoqzSdTkejR48mIQRt27aN4uPjCQAFBAQYyhQUFNDSpUupXr16NHjw
	YOWCZawW8/b2LtHHR0R09+5dMjMzo+zsbEpJSTHaxcHGxoa+++478vX1JSEEpaenG2p8Fy9eNDrP
	unXrjGp8Wq2WJk6cSHv37qX9+/dTbm6uUfmAgAAaNWpUhXGvWrWqzBrf0KFDaenSpRWeQx/z3Llz
	yy3Xrl072r59e6nvFRQU0Kuvvkru7u5UUFBQ4v38/Hz64IMPyMXFhVJSUoiIKCQkhDQaDV24cIFm
	z55tdB/R0dGk1WoNn/XHH39sFOvDn2914D4+hQgh8PTTT4OIMHToUJx+sMgjPfgr79SpU/D19cXE
	iRNx9+5dEJHhrzfGWOXpN5J92I4dO+Dl5QUbGxvDa8OHD0d8fDxu3bqFQYMGYdy4cQCAI0eOlHl+
	eqhmFhQUhMjISPTq1QuBgYFGe/wVFhYiOTkZWVlZICJkZmYa9WcV7w9zcnIq9Xr5+fnYvXs3goOD
	K7z3bt26wdvbG2FhYYbfMXoXLlzAhAkTkJqaij/++AMOpSwce+7cOQwcOBDx8fHYsWOHYReZ4kxN
	TfHWW2/hypUriI6OBgDD9ITSvPLKKxgxYgQA+XvwpZdeMlwLAK5fv478/HyjzyUzM7PCe30U3NSp
	oJkzZ+LSpUtYvXo1ZsyYAUB29oaFhWH+/Pm4f/8+XFxcsGLFCgwYMEDhaBmrnTp16lQiOeXk5CAi
	IgILFiwwen3kyJHo1q2b4Wt3d3cAsrvBzc2tzGsUb6YcOHBgmeVu376N1NRUpKamokWLFkhOTja8
	J4RAbGwsgoKCDHGXZvHixejbty+aV7JpeOPGjejRowf8/Pwwfvx4DBw4EImJiZg8eTKGDRsGe3t7
	ODk54Z133sHQoUNhZ2cHADh69Cg2bNiAjh074sCBA3juuecAAGfPnsWRI0fQp08fxMXFIScnB+Hh
	4bCwsEBQUBBOnz6N7OxsCCFgWsbafMuXL0d2djb8/f3RsmVLAPKPfUA2DTs4OBjN2fP29i73j49H
	xYlPQUIIrFy5EpcvXzb0Afz++++Gf+Bx48YhPDzc8IPIGHt0RGQ06vL27dsYO3Ys7O3t0atXLwCy
	X87CwsKobw6QfWn6GmP79u3RvHlzWFpaGpXx9vbGV199ZfjazMyszFisrKzg4uKC1q1bw9XVFf7+
	/mjcuDECAgLQoEGDMpMdIPvr4uLisHDhQhwubTHYMjRr1gxHjx5FREQEPvnkE8ybNw9CCAwdOhQR
	ERGoV68eYmJisGjRIkydOhU6nQ7u7u4YPHgwkpKSSiT8rVu3YtasWdBoNIYFN3r37o2YmBh4eHjg
	xx9/BAA89dRTcCljJXZLS0t8++23Rq+1aNECTk5O6Nu3LwD5+7FHjx5wcnIqd1f7x1LtjamsQnfu
	3KHnnnuOABAAcnd3p4MHDyodFmO1Xk5ODnXs2JF69+5Nq1evpgEDBpAQgrp06WI0UnDlypXk6OhY
	6jlmz55dZv/Xw0aNGlVikvbjWrhwIdnZ2dHq1atp+vTp5ODgQFZWVvTjjz9Wyfkf16FDh0ir1ZJG
	oyl1tKi+r65Zs2ZEJEfVltVXqRROfCoRHx9P9vb21KFDB8rJyVE6HMaeCCdOnDAatGJtbU1hYWGU
	nZ2tdGgV6tChgyFujUZDwcHBdPLkSaXDqtCRI0dIq9XSp59+SkREx44dq9TUi5okiHjGNGPsyZSd
	nY3IyEj4+/sjICBA6XAeyVdffYVbt27h7bffVjqUJw4nPsYYY3UKT2dgjDFWp3DiY4wxVqdw4mOM
	MVancOJjjDFWp3DiY4wxVqdw4mOMMVancOJjjDFWp3DiY4wxVqdw4mOMMVancOJjjDFWp3DiY4wx
	Vqdw4mOMMVancOJjjDFWp3DiY4wxVqdw4mOMMVancOJjjDFWp/w/8uJgOFS3f4YAAAAASUVORK5C
	YII=
	"
	>
	</div>
	
	</div>
	
	</div>
	
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<h2 id="Duplicating-an-XKCD-Comic">Duplicating an XKCD Comic<a class="anchor-link" href="#Duplicating-an-XKCD-Comic">&#182;</a></h2>
	</div>
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>Now let's see if we can use this to replicated an XKCD comic in matplotlib.
	This is a good one:</p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[6]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="n">Image</span><span class="p">(</span><span class="s1">&#39;http://imgs.xkcd.com/comics/front_door.png&#39;</span><span class="p">)</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt">Out[6]:</div>
	
	
	
	
	<div class="jp-RenderedImage jp-OutputArea-output jp-OutputArea-executeResult">
	<img src="data:image/png;base64,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
	Q/8f50+l/lmW+7QAAAAASUVORK5CYII=
	"
	>
	</div>
	
	</div>
	
	</div>
	
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>With the new <code>XKCDify</code> function, this is relatively easy to replicate.  The results
	are not exactly identical, but I think it definitely gets the point across!</p>
	
	</div>
	</div><div class="jp-Cell jp-CodeCell jp-Notebook-cell   ">
	<div class="jp-Cell-inputWrapper">
	<div class="jp-InputArea jp-Cell-inputArea">
	<div class="jp-InputPrompt jp-InputArea-prompt">In&nbsp;[7]:</div>
	<div class="jp-CodeMirrorEditor jp-Editor jp-InputArea-editor" data-type="inline">
		 <div class="CodeMirror cm-s-jupyter">
	<div class=" highlight hl-ipython3"><pre><span></span><span class="c1"># Some helper functions</span>
	<span class="k">def</span> <span class="nf">norm</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">x0</span><span class="p">,</span> <span class="n">sigma</span><span class="p">):</span>
		<span class="k">return</span> <span class="n">np</span><span class="o">.</span><span class="n">exp</span><span class="p">(</span><span class="o">-</span><span class="mf">0.5</span> <span class="o">*</span> <span class="p">(</span><span class="n">x</span> <span class="o">-</span> <span class="n">x0</span><span class="p">)</span> <span class="o">**</span> <span class="mi">2</span> <span class="o">/</span> <span class="n">sigma</span> <span class="o">**</span> <span class="mi">2</span><span class="p">)</span>
	
	<span class="k">def</span> <span class="nf">sigmoid</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">x0</span><span class="p">,</span> <span class="n">alpha</span><span class="p">):</span>
		<span class="k">return</span> <span class="mf">1.</span> <span class="o">/</span> <span class="p">(</span><span class="mf">1.</span> <span class="o">+</span> <span class="n">np</span><span class="o">.</span><span class="n">exp</span><span class="p">(</span><span class="o">-</span> <span class="p">(</span><span class="n">x</span> <span class="o">-</span> <span class="n">x0</span><span class="p">)</span> <span class="o">/</span> <span class="n">alpha</span><span class="p">))</span>
		
	<span class="c1"># define the curves</span>
	<span class="n">x</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">linspace</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mi">1</span><span class="p">,</span> <span class="mi">100</span><span class="p">)</span>
	<span class="n">y1</span> <span class="o">=</span> <span class="n">np</span><span class="o">.</span><span class="n">sqrt</span><span class="p">(</span><span class="n">norm</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.7</span><span class="p">,</span> <span class="mf">0.05</span><span class="p">))</span> <span class="o">+</span> <span class="mf">0.2</span> <span class="o">*</span> <span class="p">(</span><span class="mf">1.5</span> <span class="o">-</span> <span class="n">sigmoid</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.8</span><span class="p">,</span> <span class="mf">0.05</span><span class="p">))</span>
	
	<span class="n">y2</span> <span class="o">=</span> <span class="mf">0.2</span> <span class="o">*</span> <span class="n">norm</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.5</span><span class="p">,</span> <span class="mf">0.2</span><span class="p">)</span> <span class="o">+</span> <span class="n">np</span><span class="o">.</span><span class="n">sqrt</span><span class="p">(</span><span class="n">norm</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.6</span><span class="p">,</span> <span class="mf">0.05</span><span class="p">))</span> <span class="o">+</span> <span class="mf">0.1</span> <span class="o">*</span> <span class="p">(</span><span class="mi">1</span> <span class="o">-</span> <span class="n">sigmoid</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.75</span><span class="p">,</span> <span class="mf">0.05</span><span class="p">))</span>
	
	<span class="n">y3</span> <span class="o">=</span> <span class="mf">0.05</span> <span class="o">+</span> <span class="mf">1.4</span> <span class="o">*</span> <span class="n">norm</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="mf">0.85</span><span class="p">,</span> <span class="mf">0.08</span><span class="p">)</span>
	<span class="n">y3</span><span class="p">[</span><span class="n">x</span> <span class="o">&gt;</span> <span class="mf">0.85</span><span class="p">]</span> <span class="o">=</span> <span class="mf">0.05</span> <span class="o">+</span> <span class="mf">1.4</span> <span class="o">*</span> <span class="n">norm</span><span class="p">(</span><span class="n">x</span><span class="p">[</span><span class="n">x</span> <span class="o">&gt;</span> <span class="mf">0.85</span><span class="p">],</span> <span class="mf">0.85</span><span class="p">,</span> <span class="mf">0.3</span><span class="p">)</span>
	
	<span class="c1"># draw the curves</span>
	<span class="n">ax</span> <span class="o">=</span> <span class="n">pl</span><span class="o">.</span><span class="n">axes</span><span class="p">()</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y1</span><span class="p">,</span> <span class="n">c</span><span class="o">=</span><span class="s1">&#39;gray&#39;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y2</span><span class="p">,</span> <span class="n">c</span><span class="o">=</span><span class="s1">&#39;blue&#39;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">(</span><span class="n">x</span><span class="p">,</span> <span class="n">y3</span><span class="p">,</span> <span class="n">c</span><span class="o">=</span><span class="s1">&#39;red&#39;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.3</span><span class="p">,</span> <span class="o">-</span><span class="mf">0.1</span><span class="p">,</span> <span class="s2">&quot;Yard&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.5</span><span class="p">,</span> <span class="o">-</span><span class="mf">0.1</span><span class="p">,</span> <span class="s2">&quot;Steps&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.7</span><span class="p">,</span> <span class="o">-</span><span class="mf">0.1</span><span class="p">,</span> <span class="s2">&quot;Door&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.9</span><span class="p">,</span> <span class="o">-</span><span class="mf">0.1</span><span class="p">,</span> <span class="s2">&quot;Inside&quot;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.05</span><span class="p">,</span> <span class="mf">1.1</span><span class="p">,</span> <span class="s2">&quot;fear that</span><span class="se">\n</span><span class="s2">there&#39;s</span><span class="se">\n</span><span class="s2">something</span><span class="se">\n</span><span class="s2">behind me&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">([</span><span class="mf">0.15</span><span class="p">,</span> <span class="mf">0.2</span><span class="p">],</span> <span class="p">[</span><span class="mf">1.0</span><span class="p">,</span> <span class="mf">0.2</span><span class="p">],</span> <span class="s1">&#39;-k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mf">0.5</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.25</span><span class="p">,</span> <span class="mf">0.8</span><span class="p">,</span> <span class="s2">&quot;forward</span><span class="se">\n</span><span class="s2">speed&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">([</span><span class="mf">0.32</span><span class="p">,</span> <span class="mf">0.35</span><span class="p">],</span> <span class="p">[</span><span class="mf">0.75</span><span class="p">,</span> <span class="mf">0.35</span><span class="p">],</span> <span class="s1">&#39;-k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mf">0.5</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">text</span><span class="p">(</span><span class="mf">0.9</span><span class="p">,</span> <span class="mf">0.4</span><span class="p">,</span> <span class="s2">&quot;embarrassment&quot;</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">plot</span><span class="p">([</span><span class="mf">1.0</span><span class="p">,</span> <span class="mf">0.8</span><span class="p">],</span> <span class="p">[</span><span class="mf">0.55</span><span class="p">,</span> <span class="mf">1.05</span><span class="p">],</span> <span class="s1">&#39;-k&#39;</span><span class="p">,</span> <span class="n">lw</span><span class="o">=</span><span class="mf">0.5</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">set_title</span><span class="p">(</span><span class="s2">&quot;Walking back to my</span><span class="se">\n</span><span class="s2">front door at night:&quot;</span><span class="p">)</span>
	
	<span class="n">ax</span><span class="o">.</span><span class="n">set_xlim</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mi">1</span><span class="p">)</span>
	<span class="n">ax</span><span class="o">.</span><span class="n">set_ylim</span><span class="p">(</span><span class="mi">0</span><span class="p">,</span> <span class="mf">1.5</span><span class="p">)</span>
	
	<span class="c1"># modify all the axes elements in-place</span>
	<span class="n">XKCDify</span><span class="p">(</span><span class="n">ax</span><span class="p">,</span> <span class="n">expand_axes</span><span class="o">=</span><span class="kc">True</span><span class="p">)</span>
	</pre></div>
	
		 </div>
	</div>
	</div>
	</div>
	
	<div class="jp-Cell-outputWrapper">
	
	
	<div class="jp-OutputArea jp-Cell-outputArea">
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt">Out[7]:</div>
	
	
	
	
	<div class="jp-RenderedText jp-OutputArea-output jp-OutputArea-executeResult" data-mime-type="text/plain">
	<pre>&lt;matplotlib.axes.AxesSubplot at 0x2fef210&gt;</pre>
	</div>
	
	</div>
	
	<div class="jp-OutputArea-child">
	
		
		<div class="jp-OutputPrompt jp-OutputArea-prompt"></div>
	
	
	
	
	<div class="jp-RenderedImage jp-OutputArea-output ">
	<img src="data:image/png;base64,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
	WFgYOI5DYmKiwUvQKRQKKJVKJCUlIT09XWjZ2dnCMfxa+/nz59G9e3f06dMHI0eOhLu7OziOg0ql
	Qnh4uHC8Wq3GrVu3MH/+fDz33HPYuHEjfv31VwDAzp07AVRWfHJ2dsbQoUOxb98+FBUVITg4GB4e
	HlrrqeXl5UhPT4dCoYCXlxdycnLw119/YeDAgbh9+zbUarVwbE5ODnJychAYGAgA2Lx5c536BYBT
	p04hOjoajz76KJ6uphxWSUkJnnnmGSgUCnz66adVXq/tOvG0aNECubm5uHTpknCtMzIyqtQy5teI
	7969q9UX//yTTz6J0NBQ3Lp1C3v27EFubi6mT58OKysrrFu3DjExMdi9ezeuX79exc4qPFSOmxEH
	Dx4kAPTYY48Z2xSDsm/fPurTp4+xzWgUr7/+Or0jwnRjXfjlFzYga9mSSMzB9t69eyk6OppWrVol
	TIuZLGfPsotYh9GIPlm3bh1FR0fT8ePHDdqPWq0mJycnUigUpFKpiOM4srCwoK5duwojKn40NW/e
	PHJxcaFr164REdGePXuEUVrfvn2Fke2sWbOEkdWoUaOooqJC6G/+/Pkkk8mooKCAiNgo9KeffiKO
	4+jGjRtERPTss89S69atKTs7m8aOHUv+/v5ao7Vhw4ZRaWkp3bp1iziOo3Xr1tFPP/1EvXr1IgsL
	C+E4Nzc3+v333+vcb05ODnXr1o04jqOQkBCKjY2tcr1mzpxJMpmMWrZsSf/73/+0Zm3qcp3mz59P
	HMeRnZ2dYKezszPt3r1bOM/Ro0eJ4zjasmULzZw5k1QqFZ0/f56IiIYMGUIymazav+WiRYvqNGrW
	RRrZasB7IpuzcxQAWFlZoaSkxNhmNIr+/fsjLi7O2GYgLw949VW2/9FHgFhlQ7Ozs3HkyBEALNSH
	E7syvb45cYJtRRrV8vAjRUNnVTt8+DCys7Px/fffIysrC7GxsUhKSsLx48er1Jr94osvEBsbC39/
	fwCoMYxr/vz5iImJwc6dO3H9+nU88cQTwv/1qVOnQEQ4cuQI4uPjERISgkmTJgGA4Bw5ZswYXL16
	FUFBQdi8eTOuX7+OyZMnIyYmBufOncP27duhUCjg6emJRx99FHPnzsWUKVNw6NAh2NnZYePGjcKo
	buTIkXXu197eHseOHcPVq1cRGhqKQYMG4YUXXtD6bKtWrUJ+fj5mzJiBV199FX379hVmHutynbZt
	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
	"
	>
	</div>
	
	</div>
	
	</div>
	
	</div>
	
	</div>
	<div class="jp-Cell-inputWrapper"><div class="jp-InputPrompt jp-InputArea-prompt">
	</div><div class="jp-RenderedHTMLCommon jp-RenderedMarkdown jp-MarkdownOutput " data-mime-type="text/markdown">
	<p>Pretty good for a couple hours's work!</p>
	<p>I think the possibilities here are pretty limitless: this is going to be a hugely
	useful and popular feature in matplotlib, especially when the sketch artist PR is mature
	and part of the main package.  I imagine using this style of plot for schematic figures
	in presentations where the normal crisp matplotlib lines look a bit too "scientific".
	I'm giving a few talks at the end of the month... maybe I'll even use some of
	this code there.</p>
	<p>This post was written entirely in an IPython Notebook: the notebook file is available for
	download <a href="http://jakevdp.github.com/downloads/notebooks/XKCD_plots.ipynb">here</a>.
	For more information on blogging with notebooks in octopress, see my
	<a href="http://jakevdp.github.com/blog/2012/10/04/blogging-with-ipython/">previous post</a>
	on the subject.</p>
	
	</div>
	</div>
	</body>
	
	
	
	
	
	
	
	</html>`
}

// RegisterSettingsBundles pushes the settings bundle definitions for this extension to the ocis-settings service.
func RegisterSettingsBundles(l *olog.Logger) {
	request := &settings.SaveBundleRequest{
		Bundle: &settings.Bundle{
			Id:          bundleIDGreeting,
			Name:        "greeting",
			DisplayName: "Greeting",
			Extension:   "ocis-jupyter",
			Type:        settings.Bundle_TYPE_DEFAULT,
			Resource: &settings.Resource{
				Type: settings.Resource_TYPE_SYSTEM,
			},
			Settings: []*settings.Setting{
				{
					Id:          settingIDGreeterPhrase,
					Name:        "phrase",
					DisplayName: "Phrase",
					Description: "Phrase for replies on the greet request",
					Resource: &settings.Resource{
						Type: settings.Resource_TYPE_SYSTEM,
					},
					Value: &settings.Setting_StringValue{
						StringValue: &settings.String{
							Required:  true,
							Default:   "Hello",
							MaxLength: 15,
						},
					},
				},
			},
		},
	}

	// TODO this won't work with a registry other than mdns. Look into Micro's client initialization.
	// https://github.com/owncloud/ocis-proxy/issues/38
	bundleService := settings.NewBundleService("com.owncloud.api.settings", mclient.DefaultClient)
	response, err := bundleService.SaveBundle(context.Background(), request)
	if err != nil {
		l.Warn().Msg("error registering settings bundle at first try. retrying")
		for i := 1; i <= maxRetries; i++ {
			if _, err := bundleService.SaveBundle(context.Background(), request); err != nil {
				l.Warn().
					Str("bundle_name", request.Bundle.Name).
					Str("attempt", fmt.Sprintf("%v/%v", strconv.Itoa(i), strconv.Itoa(maxRetries))).
					Msgf("error creating bundle")
				continue
			} else {
				l.Info().
					Str("bundle_name", request.Bundle.Name).
					Str("after", fmt.Sprintf("%v retries", strconv.Itoa(i))).
					Str("bundleName", request.Bundle.Name).
					Str("bundleId", request.Bundle.Id).
					Msg("default settings bundle registered")
				goto OUT
			}
		}
		l.Err(err).Str("setting_name", request.Bundle.Name).Msg("bundle could not be registered. max number of retries reached")
	} else {
		l.Info().
			Str("bundleName", response.Bundle.Name).
			Str("bundleId", response.Bundle.Id).
			Msg("default settings bundle registered")
	}

OUT:
	permissionRequests := []*settings.AddSettingToBundleRequest{
		{
			BundleId: ssvc.BundleUUIDRoleAdmin,
			Setting: &settings.Setting{
				Id: "d5f42c4b-e1b6-4b59-8eca-fc4b9e9f2320",
				Resource: &settings.Resource{
					Type: settings.Resource_TYPE_SETTING,
					Id:   settingIDGreeterPhrase,
				},
				Name: "phrase-admin-read",
				Value: &settings.Setting_PermissionValue{
					PermissionValue: &settings.Permission{
						Operation:  settings.Permission_OPERATION_READWRITE,
						Constraint: settings.Permission_CONSTRAINT_OWN,
					},
				},
			},
		},
	}

	for i := range permissionRequests {
		l.Debug().Str("setting_name", permissionRequests[i].Setting.Name).Str("bundle_id", permissionRequests[i].BundleId).Msg("registering setting to bundle")
		if res, err := bundleService.AddSettingToBundle(context.Background(), permissionRequests[i]); err != nil {
			go retryPermissionRequests(context.Background(), bundleService, permissionRequests[i], maxRetries, l)
		} else {
			l.Info().Str("setting_name", res.Setting.Name).Msg("permission registered")
		}
	}
}

// proposal: the retry logic should live in the settings service.
func retryPermissionRequests(ctx context.Context, bs settings.BundleService, setting *settings.AddSettingToBundleRequest, maxRetries int, l *olog.Logger) {
	for i := 1; i < maxRetries; i++ {
		if _, err := bs.AddSettingToBundle(ctx, setting); err != nil {
			l.Warn().Str("setting_name", setting.Setting.Name).Str("attempt", fmt.Sprintf("%v/%v", strconv.Itoa(i), strconv.Itoa(maxRetries))).Msgf("error on add setting to bundle")
			continue
		}
		l.Info().Str("setting_name", setting.Setting.Name).Str("after", fmt.Sprintf("%v retries", strconv.Itoa(i))).Msg("permission registered")
		return
	}

	l.Error().Str("setting_name", setting.Setting.Name).Msg("setting could not be registered. max number of retries reached")
}
