/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Post_AuthorRead } from '../models/Post_AuthorRead';
import type { Post_LabelsList } from '../models/Post_LabelsList';
import type { PostCreate } from '../models/PostCreate';
import type { PostList } from '../models/PostList';
import type { PostRead } from '../models/PostRead';
import type { PostUpdate } from '../models/PostUpdate';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class PostService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * List Posts
     * List Posts.
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns PostList result Post list
     * @throws ApiError
     */
    public listPost(
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<PostList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/posts',
            query: {
                'page': page,
                'itemsPerPage': itemsPerPage,
            },
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * Create a new Post
     * Creates a new Post and persists it to storage.
     * @param requestBody Post to create
     * @returns PostCreate Post created
     * @throws ApiError
     */
    public createPost(
        requestBody: {
            create_time: string;
            update_time: string;
            slug: string;
            title: string;
            content: string;
            published_at: string;
            author: number;
            labels?: Array<number>;
        },
    ): CancelablePromise<PostCreate> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/posts',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `invalid input, data invalid`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * Find a Post by ID
     * Finds the Post with the requested ID and returns it.
     * @param id ID of the Post
     * @returns PostRead Post with requested ID was found
     * @throws ApiError
     */
    public readPost(
        id: number,
    ): CancelablePromise<PostRead> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/posts/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * Deletes a Post by ID
     * Deletes the Post with the requested ID.
     * @param id ID of the Post
     * @returns void
     * @throws ApiError
     */
    public deletePost(
        id: number,
    ): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/posts/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * Updates a Post
     * Updates a Post and persists changes to storage.
     * @param id ID of the Post
     * @param requestBody Post properties to update
     * @returns PostUpdate Post updated
     * @throws ApiError
     */
    public updatePost(
        id: number,
        requestBody: {
            update_time?: string;
            slug?: string;
            title?: string;
            content?: string;
            published_at?: string;
            author?: number;
            labels?: Array<number>;
        },
    ): CancelablePromise<PostUpdate> {
        return this.httpRequest.request({
            method: 'PATCH',
            url: '/posts/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * Find the attached User
     * Find the attached User of the Post with the given ID
     * @param id ID of the Post
     * @returns Post_AuthorRead User attached to Post with requested ID was found
     * @throws ApiError
     */
    public readPostAuthor(
        id: number,
    ): CancelablePromise<Post_AuthorRead> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/posts/{id}/author',
            path: {
                'id': id,
            },
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

    /**
     * List attached Labels
     * List attached Labels.
     * @param id ID of the Post
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns Post_LabelsList result Posts list
     * @throws ApiError
     */
    public listPostLabels(
        id: number,
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<Post_LabelsList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/posts/{id}/labels',
            path: {
                'id': id,
            },
            query: {
                'page': page,
                'itemsPerPage': itemsPerPage,
            },
            errors: {
                400: `invalid input, data invalid`,
                404: `resource not found`,
                409: `conflicting resources`,
                500: `unexpected error`,
            },
        });
    }

}