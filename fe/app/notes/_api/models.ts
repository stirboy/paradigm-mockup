export interface Note {
  id: string;
  title: string;
  userId: string;
  content?: string;
  isArchived: boolean;
  parentId?: string;
  coverImage?: string;
  icon?: string;
}
