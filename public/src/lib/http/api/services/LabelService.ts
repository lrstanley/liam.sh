/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Label_PostsList } from '../models/Label_PostsList';
import type { LabelCreate } from '../models/LabelCreate';
import type { LabelList } from '../models/LabelList';
import type { LabelRead } from '../models/LabelRead';
import type { LabelUpdate } from '../models/LabelUpdate';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class LabelService {

    constructor(public readonly httpRequest: BaseHttpRequest) {}

    /**
     * List Labels
     * List Labels.
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns LabelList result Label list
     * @throws ApiError
     */
    public listLabel(
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<LabelList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/labels',
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
     * Create a new Label
     * Creates a new Label and persists it to storage.
     * @param requestBody Label to create
     * @returns LabelCreate Label created
     * @throws ApiError
     */
    public createLabel(
        requestBody: {
            create_time: string;
            update_time: string;
            name: string;
            posts?: Array<number>;
        },
    ): CancelablePromise<LabelCreate> {
        return this.httpRequest.request({
            method: 'POST',
            url: '/labels',
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
     * Find a Label by ID
     * Finds the Label with the requested ID and returns it.
     * @param id ID of the Label
     * @returns LabelRead Label with requested ID was found
     * @throws ApiError
     */
    public readLabel(
        id: number,
    ): CancelablePromise<LabelRead> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/labels/{id}',
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
     * Deletes a Label by ID
     * Deletes the Label with the requested ID.
     * @param id ID of the Label
     * @returns void
     * @throws ApiError
     */
    public deleteLabel(
        id: number,
    ): CancelablePromise<void> {
        return this.httpRequest.request({
            method: 'DELETE',
            url: '/labels/{id}',
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
     * Updates a Label
     * Updates a Label and persists changes to storage.
     * @param id ID of the Label
     * @param requestBody Label properties to update
     * @returns LabelUpdate Label updated
     * @throws ApiError
     */
    public updateLabel(
        id: number,
        requestBody: {
            update_time?: string;
            name?: string;
            posts?: Array<number>;
        },
    ): CancelablePromise<LabelUpdate> {
        return this.httpRequest.request({
            method: 'PATCH',
            url: '/labels/{id}',
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
     * @param id ID of the Label
     * @param page what page to render
     * @param itemsPerPage item count to render per page
     * @returns Label_PostsList result Labels list
     * @throws ApiError
     */
    public listLabelPosts(
        id: number,
        page?: number,
        itemsPerPage?: number,
    ): CancelablePromise<Array<Label_PostsList>> {
        return this.httpRequest.request({
            method: 'GET',
            url: '/labels/{id}/posts',
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