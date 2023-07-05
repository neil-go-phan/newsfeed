import React from 'react';
import DeleteBtn from './deleteBtn';
import UpdateBtn from './updateBtn';
// import PreviewBtn from './preview';


type Props = {
  id: number;
  articlesSource: ArticlesSourceInfo;
  handleDeleteArticlesSource: (id: number) => void;
  handleUpdate: (articlesSource: UpdateArticleSourcePayload) => void;
};


const AdminArticlesSourcesAction: React.FC<Props> = (props: Props) => {

  return (
    <div className="action">
      <div className="d-flex flex-column">
        <DeleteBtn
          id={props.id}
          handleDeleteArticlesSource={props.handleDeleteArticlesSource}
        />
        <UpdateBtn articlesSource={props.articlesSource} handleUpdate={props.handleUpdate}/>
      </div>
    </div>
  );
};

export default AdminArticlesSourcesAction;
