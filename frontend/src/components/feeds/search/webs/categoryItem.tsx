import React from 'react';

type Props = {
  category: Category;
  onClickHandler: (categoryID: number) => void;
};

const CategoryItem: React.FC<Props> = (props: Props) => {
  return (
    <div
      className="categoryItem"
      style={{
        backgroundImage:
          "url('" +
          props.category.illustration.replace(/(\r\n|\n|\r)/gm, '') +
          "')",
      }}
      onClick={() => props.onClickHandler(props.category.id)}
    >
      {props.category.name}
    </div>
  );
};

export default CategoryItem;
