/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';

export default function useBlitScript() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.BlitScript.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryScript = (address: string,  options: any) => {
    const key = { type: 'QueryScript',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.BlitScript.query.queryScript(address).then( res => res.data );
    }, options);
  }
  
  const QueryScriptAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryScriptAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.BlitScript.query.queryScriptAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryEval = (script_address: string, query: any, options: any) => {
    const key = { type: 'QueryEval',  script_address, query };    
    return useQuery([key], () => {
      const { script_address,query } = key
      return  client.BlitScript.query.queryEval(script_address, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryWeb = (address: string, query: any, options: any) => {
    const key = { type: 'QueryWeb',  address, query };    
    return useQuery([key], () => {
      const { address,query } = key
      return  client.BlitScript.query.queryWeb(address, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryScript,QueryScriptAll,QueryEval,QueryWeb,
  }
}