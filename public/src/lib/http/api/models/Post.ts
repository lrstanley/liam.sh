/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Label } from './Label';
import type { User } from './User';

export type Post = {
    id: number;
    create_time: string;
    update_time: string;
    slug: string;
    title: string;
    content: string;
    published_at: string;
    author: User;
    labels?: Array<Label>;
};
