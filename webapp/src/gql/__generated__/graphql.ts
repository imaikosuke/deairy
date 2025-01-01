/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core";
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T,
> = { [_ in K]?: never };
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends " $fragmentName" | "__typename" ? T[P] : never;
    };
export type DateString = string & { readonly brand: unique symbol };
export type TimeString = string & { readonly brand: unique symbol };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
  Date: { input: DateString; output: DateString };
  DateTime: { input: string; output: string };
  Time: { input: TimeString; output: TimeString };
  Uint: { input: number; output: number };
  Uint32: { input: number; output: number };
};

export type CreateDiaryInput = {
  /** 本文 */
  content: Scalars["String"]["input"];
  /** タグ */
  tags?: InputMaybe<Array<Scalars["String"]["input"]>>;
  /** タイトル */
  title: Scalars["String"]["input"];
  /** 公開設定 */
  visibility: DiaryVisibility;
};

export type CreateUserInput = {
  /** メールアドレス */
  email: Scalars["String"]["input"];
  /** ユーザ名 */
  name: Scalars["String"]["input"];
};

/** 日記を表す型 */
export type Diary = {
  /** 作成者 */
  author: User;
  /** 本文 */
  content: Scalars["String"]["output"];
  /** 作成日時 */
  createdAt: Scalars["DateTime"]["output"];
  /** 日記ID */
  id: Scalars["ID"]["output"];
  /** タグ一覧 */
  tags: Array<Scalars["String"]["output"]>;
  /** タイトル */
  title: Scalars["String"]["output"];
  /** 更新日時 */
  updatedAt: Scalars["DateTime"]["output"];
  /** 公開設定 */
  visibility: DiaryVisibility;
};

/** 日記の公開設定 */
export type DiaryVisibility =
  /** 非公開 */
  | "PRIVATE"
  /** 公開 */
  | "PUBLIC";

export type Mutation = {
  /** 日記作成 */
  createDiary: Diary;
  /** ユーザ作成 */
  createUser: User;
  /** 日記削除 */
  deleteDiary: Scalars["Boolean"]["output"];
  /** 日記更新 */
  updateDiary: Diary;
};

export type MutationCreateDiaryArgs = {
  input: CreateDiaryInput;
};

export type MutationCreateUserArgs = {
  input: CreateUserInput;
};

export type MutationDeleteDiaryArgs = {
  id: Scalars["ID"]["input"];
};

export type MutationUpdateDiaryArgs = {
  id: Scalars["ID"]["input"];
  input: UpdateDiaryInput;
};

export type Query = {
  /** 日記一覧取得 */
  diaries: Array<Diary>;
  /** 日記取得 */
  diary: Diary;
  /** ユーザ取得 */
  user: User;
  /** ユーザ一覧取得 */
  users: Array<User>;
};

export type QueryDiariesArgs = {
  userId: InputMaybe<Scalars["ID"]["input"]>;
  visibility: InputMaybe<DiaryVisibility>;
};

export type QueryDiaryArgs = {
  id: Scalars["ID"]["input"];
};

export type QueryUserArgs = {
  id: Scalars["ID"]["input"];
};

export type UpdateDiaryInput = {
  /** 本文 */
  content: Scalars["String"]["input"];
  /** タグ */
  tags?: InputMaybe<Array<Scalars["String"]["input"]>>;
  /** タイトル */
  title: Scalars["String"]["input"];
  /** 公開設定 */
  visibility: DiaryVisibility;
};

export type UpdateUserInput = {
  email: Scalars["String"]["input"];
  name: Scalars["String"]["input"];
};

/** ユーザを表す型 */
export type User = {
  /** 作成日時 */
  createdAt: Scalars["DateTime"]["output"];
  /** ユーザの日記一覧 */
  diaries: Array<Diary>;
  /** メールアドレス */
  email: Scalars["String"]["output"];
  /** ユーザID */
  id: Scalars["ID"]["output"];
  /** ユーザ名 */
  name: Scalars["String"]["output"];
  /** 更新日時 */
  updatedAt: Scalars["DateTime"]["output"];
};

export type GetUserListPageQueryVariables = Exact<{ [key: string]: never }>;

export type GetUserListPageQuery = {
  users: Array<{ id: string; name: string; email: string }>;
};

export type UserListTableUserFragment = {
  id: string;
  name: string;
  email: string;
};

export const UserListTableUserFragmentDoc = {
  kind: "Document",
  definitions: [
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "UserListTableUser" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "User" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "id" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "email" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<UserListTableUserFragment, unknown>;
export const GetUserListPageDocument = {
  kind: "Document",
  definitions: [
    {
      kind: "OperationDefinition",
      operation: "query",
      name: { kind: "Name", value: "GetUserListPage" },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          {
            kind: "Field",
            name: { kind: "Name", value: "users" },
            selectionSet: {
              kind: "SelectionSet",
              selections: [
                {
                  kind: "FragmentSpread",
                  name: { kind: "Name", value: "UserListTableUser" },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: "FragmentDefinition",
      name: { kind: "Name", value: "UserListTableUser" },
      typeCondition: {
        kind: "NamedType",
        name: { kind: "Name", value: "User" },
      },
      selectionSet: {
        kind: "SelectionSet",
        selections: [
          { kind: "Field", name: { kind: "Name", value: "id" } },
          { kind: "Field", name: { kind: "Name", value: "name" } },
          { kind: "Field", name: { kind: "Name", value: "email" } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  GetUserListPageQuery,
  GetUserListPageQueryVariables
>;
