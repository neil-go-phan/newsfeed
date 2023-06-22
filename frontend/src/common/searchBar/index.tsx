import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { useRouter } from 'next/router';
import React, { useEffect, useRef, useState } from 'react';

type Props = {
  placeholder: string;
  api: string;
  getSearchResult: (result: any) => void;
};

const SEARCH_FAIL_MESSAGE = 'search fail';

const SearchBar: React.FC<Props> = (props: Props) => {
  const [keyword, setKeyword] = useState<string>('');
  const router = useRouter();
  const abortControllerRef = useRef<AbortController>(new AbortController());
  const resources = useRef<Map<string, any>>(new Map());

  const handleTypeKeyword = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newKeyword = event.target.value;
    setKeyword(newKeyword);
    // cancle pending request
    abortControllerRef.current.abort();

    const cacheResult = resources.current.get(props.api);
    if (cacheResult) {
      // setResult(cacheResult)
      props.getSearchResult(cacheResult);
      return;
    }
    requestSearch(props.api);
  };

  useEffect(() => {
    resources.current.clear();
  }, [router.asPath]);

  const requestSearch = async (api: string) => {
    try {
      const { data } = await axiosProtectedAPI.get(`${api}q=${keyword}`, {
        signal: abortControllerRef.current.signal,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw SEARCH_FAIL_MESSAGE;
      }
      props.getSearchResult(data);
    } catch (error: any) {
      props.getSearchResult(error);
    }
  };

  return (
    <div className="searchBar">
      <input
        placeholder={props.placeholder}
        value={keyword}
        type='text'
        onChange={(event) => handleTypeKeyword(event)}
      />
    </div>
  );
};

export default SearchBar;
