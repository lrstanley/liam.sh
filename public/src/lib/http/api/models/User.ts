/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Post } from './Post';

export type User = {
    id: number;
    create_time: string;
    update_time: string;
    user_id: number;
    login: string;
    name?: string;
    avatar_url?: string;
    email?: string;
    location?: string;
    bio?: string;
    posts?: Array<Post>;
};
