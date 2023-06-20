import React from 'react';
import DeleteBtn from './deleteBtn';
import UpdateBtn from './updateBtn';

type Props = {
  id: number;
  topicName: string;
  category: Category;
  categories: Categories;
  isDisabled: boolean;
  handleDeleteTopic: (id: number, name: string, category_id: number) => void;
  handleUpdateTopic: (id: number, newName: string, newCategoryID: number) => void;
};

const AdminTopicAction: React.FC<Props> = (props: Props) => {
  return (
    <div className="action">
      <div className="d-flex">
        <DeleteBtn
          id={props.id}
          topicName={props.topicName}
          isDisabled={props.isDisabled}
          category_id={props.category.id}
          handleDeleteTopic={props.handleDeleteTopic}
        />
        <UpdateBtn
          id={props.id}
          topicName={props.topicName}
          isDisabled={props.isDisabled}
          category={props.category}
          categories={props.categories}
          handleUpdateTopic={props.handleUpdateTopic}
        />
      </div>
    </div>
  );
};

export default AdminTopicAction;
