import type Record from 'pocketbase';

export enum ProfileType {
  Work = "work",
  Education = "education",
  Project = "project"
}

export interface TagRecord extends Record {
  name: string;
}

export interface ResponsibilityRecord extends Record {
  description: string;
  order: number;
}

export interface PostRecord extends Record {
  body: string;
  title: string;
  created: Date;
  slug: string;
  subTitle: string;
  expand: { tags: TagRecord[] };
}

export interface ProfileRecord extends Record {
  institutionName?: string;
  location?: string;
  title: string;
  description?: string;
  from?: Date;
  to?: Date;
  current: boolean;
  type: ProfileType;
  expand: { responsibilities: ResponsibilityRecord[] }
}
