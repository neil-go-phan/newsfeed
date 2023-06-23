import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { _ROUTES } from '@/helpers/constants';
import SearchWebs from './webs';
import Link from 'next/link';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faNewspaper, faRss } from '@fortawesome/free-solid-svg-icons';
import SearchBar from '@/common/searchBar';
import FilterByCategory from './webs/category';
import { CategoriesContext } from '@/common/contexts/categoriesContext';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import SearchWebsResult from './webs/searchWebsResult';
import { SearchKeywordContext } from '@/common/contexts/searchKeywordContext';

const LIST_CATEGORY_REQUEST_FAIL_MESSAGE = 'reqeust failed';
const INPUT_PLACE_HOLDER =
  'Follow your favorite source and nerver miss a story';

function SearchFeedsComponent() {
  const router = useRouter();
  const [path, setpath] = useState<string>();
  const [categories, setCategories] = useState<Categories>([]);
  const [keyword, setKeyword] = useState<string>('');

  useEffect(() => {
    const path = router.asPath;
    const beforeQuestionMark = path.split('?')[0];
    setpath(beforeQuestionMark);
  }, [router.asPath]);

  useEffect(() => {
    requestListCategory();
  }, []);

  const requestListCategory = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('category/list-all');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw LIST_CATEGORY_REQUEST_FAIL_MESSAGE;
      }
      setCategories(data.categories);
    } catch (error: any) {
      setCategories([]);
    }
  };

  const render = () => {
    switch (path) {
      case _ROUTES.FEEDS_SEARCH_WEBS:
        return <SearchWebs />;
      case _ROUTES.FEEDS_SEARCH_WEBS_CATEGORY:
        return <FilterByCategory />;
      case _ROUTES.FEEDS_SEARCH_WEBS_RESULT:
        return <SearchWebsResult />;
      default:
        return <SearchWebs />;
    }
  };

  const pushToResultPage = (keyword: string) => {
    if (keyword === '') {
      router.push(_ROUTES.FEEDS_SEARCH_WEBS);
      return;
    }
    if (router.asPath !== _ROUTES.FEEDS_SEARCH_WEBS_RESULT) {
      router.push(_ROUTES.FEEDS_SEARCH_WEBS_RESULT);
    }
    // router.push({
    //   pathname: _ROUTES.FEEDS_SEARCH_WEBS_RESULT,
    //   query: { q: keyword },
    // });
  };

  return (
    <SearchKeywordContext.Provider value={{ keyword, setKeyword }}>
      <div className="searchLayout">
        <div className="searchLayout__title">
          <h1>Find the best information sources</h1>
          <p>You can search for articles, website feeds.</p>
        </div>
        <div className="searchLayout__sticky">
          <div className="searchLayout__tabs">
            <div className="searchLayout__tabs--btns">
              <Link href={_ROUTES.FEEDS_SEARCH_ARTICLES}>
                <div
                  className={
                    path === _ROUTES.FEEDS_SEARCH_ARTICLES
                      ? 'routeBtn active'
                      : 'routeBtn'
                  }
                >
                  <FontAwesomeIcon icon={faNewspaper} />
                  <span>Articles</span>
                </div>
              </Link>
              <Link href={_ROUTES.FEEDS_SEARCH_WEBS}>
                <div
                  className={
                    path === _ROUTES.FEEDS_SEARCH_WEBS ||
                    path === _ROUTES.FEEDS_SEARCH_WEBS_CATEGORY
                      ? 'routeBtn active'
                      : 'routeBtn'
                  }
                >
                  <FontAwesomeIcon icon={faRss} />
                  <span>Feeds</span>
                </div>
              </Link>
            </div>
          </div>
          <div className="searchLayout__searchBar">
            <SearchBar
              placeholder={INPUT_PLACE_HOLDER}
              handleAPI={pushToResultPage}
            />
          </div>
        </div>
        <CategoriesContext.Provider value={{ categories, setCategories }}>
          <div className="searchLayout__content">{render()}</div>
        </CategoriesContext.Provider>
      </div>
    </SearchKeywordContext.Provider>
  );
}

export default SearchFeedsComponent;
