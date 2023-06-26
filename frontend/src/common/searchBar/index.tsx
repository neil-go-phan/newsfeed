import React, { useContext, useEffect, useState } from 'react';
import { SearchKeywordContext } from '../contexts/searchKeywordContext';

type Props = {
  placeholder: string;
  handleAPI: (keyword: string) => void;
};


const SearchBar: React.FC<Props> = (props: Props) => {
  // const [keyword, setKeyword] = useState<string>('');
  const { keyword, setKeyword } = useContext(SearchKeywordContext);
  const handleTypeKeyword = (event: React.ChangeEvent<HTMLInputElement>) => {
    const newKeyword = event.target.value;
    setKeyword(newKeyword);
  };

  useEffect(() => {
    const delayDebounceFn = setTimeout(() => {
      props.handleAPI(keyword);
    }, 500);

    return () => clearTimeout(delayDebounceFn);
  }, [keyword]);

  return (
    <div className="searchBar">
      <input
        placeholder={props.placeholder}
        value={keyword}
        type="text"
        onChange={(event) => handleTypeKeyword(event)}
      />
    </div>
  );
};

export default SearchBar;
