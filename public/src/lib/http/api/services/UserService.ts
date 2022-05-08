/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { User_PostsList } from '../models/User_PostsList';
import type { UserCreate } from '../models/UserCreate';
import type { UserList } from '../models/UserList';
import type { UserRead } from '../models/UserRead';
import type { UserUpdate } from '../models/UserUpdate';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class UserService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * List Users
     * List Users.
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns UserList result User list
     * @throws ApiError
     */
    public listUser(
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<UserList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users',
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
     * Create a new User
     * Creates a new User and persists it to storage.
     * @param requestBody User to create
     * @returns UserCreate User created
     * @throws ApiError
     */
    public createUser(
        requestBody: {
            create_time: string;
            update_time: string;
            user_id: number;
            login: string;
            name?: string;
            avatar_url?: string;
            email?: string;
            location?: string;
            bio?: string;
            posts?: Array<number>;
        },
    ): CancelablePromise<UserCreate> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/users',
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
     * Find a User by ID
     * Finds the User with the requested ID and returns it.
     * @param id ID of the User
     * @returns UserRead User with requested ID was found
     * @throws ApiError
     */
    public readUser(
        id: number,
    ): CancelablePromise<UserRead> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/{id}',
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
     * Deletes a User by ID
     * Deletes the User with the requested ID.
     * @param id ID of the User
     * @returns void
     * @throws ApiError
     */
    public deleteUser(
        id: number,
    ): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/users/{id}',
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
     * Updates a User
     * Updates a User and persists changes to storage.
     * @param id ID of the User
     * @param requestBody User properties to update
     * @returns UserUpdate User updated
     * @throws ApiError
     */
    public updateUser(
        id: number,
        requestBody: {
            update_time?: string;
            login?: string;
            name?: string;
            avatar_url?: string;
            email?: string;
            location?: string;
            bio?: string;
            posts?: Array<number>;
        },
    ): CancelablePromise<UserUpdate> {
        return this.httpRequest.request({
            method: 'PATCH',
            url: '/users/{id}',
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
     * List attached Posts
     * List attached Posts.
     * @param id ID of the User
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns User_PostsList result Users list
     * @throws ApiError
     */
    public listUserPosts(
        id: number,
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<User_PostsList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/users/{id}/posts',
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