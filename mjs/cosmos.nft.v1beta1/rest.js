/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */
import axios from "axios";
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    instance;
    securityData = null;
    securityWorker;
    secure;
    format;
    constructor({ securityWorker, secure, format, ...axiosConfig } = {}) {
        this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
        this.secure = secure;
        this.format = format;
        this.securityWorker = securityWorker;
    }
    setSecurityData = (data) => {
        this.securityData = data;
    };
    mergeRequestParams(params1, params2) {
        return {
            ...this.instance.defaults,
            ...params1,
            ...(params2 || {}),
            headers: {
                ...(this.instance.defaults.headers || {}),
                ...(params1.headers || {}),
                ...((params2 && params2.headers) || {}),
            },
        };
    }
    createFormData(input) {
        return Object.keys(input || {}).reduce((formData, key) => {
            const property = input[key];
            formData.append(key, property instanceof Blob
                ? property
                : typeof property === "object" && property !== null
                    ? JSON.stringify(property)
                    : `${property}`);
            return formData;
        }, new FormData());
    }
    request = async ({ secure, path, type, query, format, body, ...params }) => {
        const secureParams = ((typeof secure === "boolean" ? secure : this.secure) &&
            this.securityWorker &&
            (await this.securityWorker(this.securityData))) ||
            {};
        const requestParams = this.mergeRequestParams(params, secureParams);
        const responseFormat = (format && this.format) || void 0;
        if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
            requestParams.headers.common = { Accept: "*/*" };
            requestParams.headers.post = {};
            requestParams.headers.put = {};
            body = this.createFormData(body);
        }
        return this.instance.request({
            ...requestParams,
            headers: {
                ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
                ...(requestParams.headers || {}),
            },
            params: query,
            responseType: responseFormat,
            data: body,
            url: path,
        });
    };
}
/**
 * @title HTTP API Console
 */
export class Api extends HttpClient {
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Balance
     * @summary Balance queries the number of NFTs of a given class owned by the owner, same as balanceOf in ERC721
     * @request GET:/cosmos/nft/v1beta1/balance/{owner}/{class_id}
     */
    cosmosNftV1Beta1Balance = (owner, classId, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/balance/${owner}/${classId}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Classes
     * @summary Classes queries all NFT classes
     * @request GET:/cosmos/nft/v1beta1/classes
     */
    cosmosNftV1Beta1Classes = (query, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/classes`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Class
     * @summary Class queries an NFT class based on its id
     * @request GET:/cosmos/nft/v1beta1/classes/{class_id}
     */
    cosmosNftV1Beta1Class = (classId, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/classes/${classId}`,
        method: "GET",
        ...params,
    });
    /**
   * No description
   *
   * @tags Query
   * @name CosmosNftV1Beta1NfTs
   * @summary NFTs queries all NFTs of a given class or owner,choose at least one of the two, similar to tokenByIndex in
  ERC721Enumerable
   * @request GET:/cosmos/nft/v1beta1/nfts
   */
    cosmosNftV1Beta1NFTs = (query, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/nfts`,
        method: "GET",
        query: query,
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Nft
     * @summary NFT queries an NFT based on its class and id.
     * @request GET:/cosmos/nft/v1beta1/nfts/{class_id}/{id}
     */
    cosmosNftV1Beta1NFT = (classId, id, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/nfts/${classId}/${id}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Owner
     * @summary Owner queries the owner of the NFT based on its class and id, same as ownerOf in ERC721
     * @request GET:/cosmos/nft/v1beta1/owner/{class_id}/{id}
     */
    cosmosNftV1Beta1Owner = (classId, id, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/owner/${classId}/${id}`,
        method: "GET",
        ...params,
    });
    /**
     * No description
     *
     * @tags Query
     * @name CosmosNftV1Beta1Supply
     * @summary Supply queries the number of NFTs from the given class, same as totalSupply of ERC721.
     * @request GET:/cosmos/nft/v1beta1/supply/{class_id}
     */
    cosmosNftV1Beta1Supply = (classId, params = {}) => this.request({
        path: `/cosmos/nft/v1beta1/supply/${classId}`,
        method: "GET",
        ...params,
    });
}
//# sourceMappingURL=rest.js.map