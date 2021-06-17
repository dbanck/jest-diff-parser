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

        // Get the word within the selection
        const text = document.getText(selection);
        if (text.length > 0) {
          console.log(
            "🚀 ~ file: extension.ts ~ line 29 ~ activate ~ word",
            text
          );
        }

        // editor.edit(editBuilder => {
        //   editBuilder.replace(selection, reversed);
        // });
      }
    }
  );

  context.subscriptions.push(disposable);
}

export function deactivate() {}
