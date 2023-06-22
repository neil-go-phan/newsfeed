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

const LIST_CATEGORY_REQUEST_FAIL_MESSAGE = 'reqeust failed';
const SEARCH_ARTICLES_SOURCE_API = '/search/article-source?';
const INPUT_PLACE_HOLDER =
  'Follow your favorite source and nerver miss a story';

function SearchFeedsComponent() {
  const router = useRouter();
  const [path, setpath] = useState<string>();
  const [categories, setCategories] = useState<Categories>([]);
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
      default:
        return <SearchWebs />;
    }
  };
  const getSearchResult = (result: any) => {
    console.log(result);
  };
  return (
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
                  path === _ROUTES.FEEDS_SEARCH_WEBS || path === _ROUTES.FEEDS_SEARCH_WEBS_CATEGORY
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
            api={SEARCH_ARTICLES_SOURCE_API}
            placeholder={INPUT_PLACE_HOLDER}
            getSearchResult={getSearchResult}
          />
        </div>
      </div>
      <CategoriesContext.Provider value={{ categories, setCategories }}>
        <div className="searchLayout__content">{render()}</div>
      </CategoriesContext.Provider>
    </div>
  );
}

export default SearchFeedsComponent;
