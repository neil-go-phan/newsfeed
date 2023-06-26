import React, { useContext } from 'react';
import CategoryItem from './categoryItem';
import { useRouter } from 'next/router';
import { _ROUTES } from '@/helpers/constants';
import { CategoriesContext } from '@/common/contexts/categoriesContext';

function SearchWebs() {
  const {categories} = useContext(CategoriesContext)
  const router = useRouter()
  const onClickCategoryItemHandler = (categoryID:number) => {
    router.push({
      pathname: _ROUTES.FEEDS_SEARCH_WEBS_CATEGORY,
      query: { category_id: categoryID },
    })
  }

  return (
    <div className="searchWebs">
      <div className="searchWebs__categories">
        <div className="searchWebs__categories--title">
          <h2>Or try our featured collections</h2>
        </div>
        <div className="searchWebs__categories--listCategory">
          {categories.map((category) => <CategoryItem key={`categoryItme-${category.name}`} category={category} onClickHandler={onClickCategoryItemHandler}/>)}
        </div>
      </div>
    </div>
  );
}

export default SearchWebs;
