/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { PostRead_Author } from './PostRead_Author';
import type { PostRead_Labels } from './PostRead_Labels';

export type PostRead = {
    id: number;
    create_time: string;
    update_time: string;
    slug: string;
    title: string;
    content: string;
    published_at: string;
    author: PostRead_Author;
    labels?: Array<PostRead_Labels>;
};
