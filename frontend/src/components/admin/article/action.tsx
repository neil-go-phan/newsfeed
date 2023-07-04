import React from 'react';
import DeleteBtn from './deleteBtn';
import PreviewBtn from './preview';

type Props = {
  id: number;
  article: Article;
  handleDeleteArticle: (id: number) => void;
};

const AdminArticlesAction: React.FC<Props> = (props: Props) => {
  return (
    <div className="action">
      <div className="d-flex flex-column">
        <DeleteBtn
          id={props.id}
          handleDeleteArticle={props.handleDeleteArticle}
        />
        <PreviewBtn article={props.article}/>
      </div>
    </div>
  );
};

export default AdminArticlesAction;
