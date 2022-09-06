
export function HandleFetchErrors(response){
    if(!response.ok){
        throw Error(response.statusText);
    }
    return response;
}