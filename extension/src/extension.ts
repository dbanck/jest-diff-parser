import * as path from "path";
import * as vscode from "vscode";

import {
  Executable,
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
} from "vscode-languageclient/node";

let client: LanguageClient;

function newClient(command: string) {
  const serverArgs = ["serve"];
  const executable: Executable = {
    command,
    args: serverArgs,
    options: {},
  };
  const serverOptions: ServerOptions = {
    run: executable,
    debug: executable,
  };

  const clientOptions: LanguageClientOptions = {
    // TODO! define file type
    documentSelector: [{ scheme: "file", language: "plaintext" }],
  };

  return new LanguageClient(
    `languageServerJestDiff`,
    `Language Server Jest Diff`,
    serverOptions,
    clientOptions
  );
}

export function activate(context: vscode.ExtensionContext) {
  const binary = context.asAbsolutePath(
    path.join("..", "server", "jest-diff-parser")
  );
  let client = newClient(binary);
  client.start();

  const disposable = vscode.commands.registerCommand(
    "jest-diff-parser.parseSelection",
    async () => {
      // Get the active text editor
      const editor = vscode.window.activeTextEditor;

      if (editor) {
        const document = editor.document;
        const selection = editor.selection;

        // Get the the selection
        const text = document.getText(selection);
        if (text.length > 0) {
          try {
            // Send selection to ls for replacement
            const result = await client.sendRequest<string>(
              "$/formatJestDiff",
              text
            );

            editor.edit((editBuilder) => {
              editBuilder.replace(selection, result);
            });
          } catch (error) {
            console.error("Got bad result from ls", error);
          }
        }
      }
    }
  );

  context.subscriptions.push(disposable);
}

export function deactivate(): Thenable<void> | undefined {
  if (!client) {
    return undefined;
  }
  return client.stop();
}
