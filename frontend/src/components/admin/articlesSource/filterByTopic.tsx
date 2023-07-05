import { alertError } from '@/helpers/alert';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import SelectSearch, {
  SelectSearchOption,
  SelectedOptionValue,
} from 'react-select-search';

const GET_CATEGORIES_NAME_FAIL_MESSAGE = 'get categories names fail';
const GET_TOPICS_FAIL_MESSAGE = 'get topics fail';

type Props = {
  handleChooseTopicID: (id: number) => void;
};

const SEARCH_CATEGORY_RESULT_EMPTY = 'Not found categories match keyword';
const SEARCH_TOPIC_RESULT_EMPTY = 'Not found topics match keyword';

const NOT_FILTER: SelectSearchOption = {
  name: 'All',
  value: 0,
};

const FilterByTopic: React.FC<Props> = (props: Props) => {
  const [categoryOptions, setCategoryOptions] = useState<
    Array<SelectSearchOption>
  >([]);
  const [topicOptions, setTopicOptions] = useState<Array<SelectSearchOption>>(
    []
  );
  const [categories, setCategories] = useState<Categories>([]);
  const [categoryID, setCategoryID] = useState<number>(0)
  const [topics, setTopics] = useState<Topics>([]);

  useEffect(() => {
    requestListCategoriesNames();
    requestListTopics();
  }, []);

  useEffect(() => {
    castCategoryToOptions();
  }, [categories]);

  useEffect(() => {
    castTopicToOptions();
  }, [topics]);

  useEffect(() => {
    castTopicToOptions();
  }, [categoryID]);

  const requestListCategoriesNames = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('category/list/name');
      if (!data.success) {
        throw GET_CATEGORIES_NAME_FAIL_MESSAGE;
      }
      setCategories(data.categories);
    } catch (error: any) {
      alertError(error);
    }
  };

  const requestListTopics = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('topic/list');
      if (!data.success) {
        throw GET_TOPICS_FAIL_MESSAGE;
      }
      setTopics(data.topics);
    } catch (error: any) {
      alertError(error);
    }
  };

  const handleChooseTopic = (
    e: SelectedOptionValue | SelectedOptionValue[]
  ) => {
    const select = e as SelectedOptionValue;
    const num = +select;
    props.handleChooseTopicID(num);
  };

  const handleChooseCategry = (
    e: SelectedOptionValue | SelectedOptionValue[]
  ) => {
    const select = e as SelectedOptionValue;
    const num = +select;
    setCategoryID(num);
  };

  const castCategoryToOptions = () => {
    const tempOptions: Array<SelectSearchOption> = [];
    tempOptions.push(NOT_FILTER);
    categories.forEach((category) => {
      const tempOption: SelectSearchOption = {
        name: category.name,
        value: category.id,
      };
      tempOptions.push(tempOption);
    });

    setCategoryOptions([...tempOptions]);
  };

  const castTopicToOptions = () => {
    const tempOptions: Array<SelectSearchOption> = [];
    if (categoryID === 0) {
      tempOptions.push(NOT_FILTER);
    }
    topics.forEach((topic) => {
      if (categoryID === 0) {
        const tempOption: SelectSearchOption = {
          name: topic.name,
          value: topic.id,
        };
        tempOptions.push(tempOption);
      } else {
        if (topic.category_id === categoryID) {
          const tempOption: SelectSearchOption = {
            name: topic.name,
            value: topic.id,
          };
          tempOptions.push(tempOption);
        }
      }
    });

    setTopicOptions([...tempOptions]);
  };

  return (
    <div className='d-flex'>
      <div className="category mx-3">
        <SelectSearch
          options={categoryOptions}
          search={true}
          onChange={(e) => handleChooseCategry(e)}
          emptyMessage={SEARCH_CATEGORY_RESULT_EMPTY}
          placeholder="Choose category..."
        />
      </div>
      <div className="topic">
        <SelectSearch
          options={topicOptions}
          search={true}
          onChange={(e) => handleChooseTopic(e)}
          emptyMessage={SEARCH_TOPIC_RESULT_EMPTY}
          placeholder="Choose topic..."
        />
      </div>
    </div>
  );
};

export default FilterByTopic;
