/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Post } from './Post';

export type Label = {
    id: number;
    create_time: string;
    update_time: string;
    name: string;
    posts?: Array<Post>;
};
