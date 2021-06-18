import * as vscode from "vscode";

export function activate(context: vscode.ExtensionContext) {
  let disposable = vscode.commands.registerCommand(
    "jest-diff-parser.parseSelection",
    () => {
      // Get the active text editor
      const editor = vscode.window.activeTextEditor;

      if (editor) {
        const document = editor.document;
        const selection = editor.selection;

        // Get the the selection
        const text = document.getText(selection);
        if (text.length > 0) {
          // TODO! validate if text is valid jest diff

          editor.edit((editBuilder) => {
            // TODO! replace with better parsing
            editBuilder.replace(
              selection,
              text.replace(/Array\s|Object\s|\+\s|\-\s|/g, "")
            );
          });
        }
      }
    }
  );

  context.subscriptions.push(disposable);
}

export function deactivate() {}
