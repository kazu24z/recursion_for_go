#!/bin/bash

# 引数チェック
if [ $# -ne 1 ]; then
  echo "使用方法: $0 <フォルダ名>"
  exit 1
fi

FOLDER_NAME=$1

# フォルダが既に存在するかチェック
if [ -d "$FOLDER_NAME" ]; then
  echo "エラー: フォルダ '$FOLDER_NAME' は既に存在します"
  exit 1
fi

# template/ディレクトリをコピー
cp -r template "$FOLDER_NAME"

echo "フォルダ '$FOLDER_NAME' を作成しました"
