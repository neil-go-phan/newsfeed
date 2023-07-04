import { alertError } from '@/helpers/alert';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import SelectSearch, {
  SelectSearchOption,
  SelectedOptionValue,
} from 'react-select-search';

const LIST_SOURCES_FAIL_MESSAGE = 'fail';
type Props = {
  handleChooseSourceID: (id: number) => void;
};

const SEARCH_RESULT_EMPTY = 'Not found article source match keyword';

const NOT_FILTER: SelectSearchOption = {
  name: "All",
  value: 0,
};

const DropdownFilterBySources: React.FC<Props> = (props: Props) => {
  const [options, setOptions] = useState<Array<SelectSearchOption>>([]);
  const [sources, setSources] = useState<ArticlesSourceInfoes>([]);

  useEffect(() => {
    requestListAll();
  }, []);

  useEffect(() => {
    castSourceToOptions();
  }, [sources]);

  const handleChooseSource = (
    e: SelectedOptionValue | SelectedOptionValue[]
  ) => {
    const select = e as SelectedOptionValue;
    const num = +select;
    props.handleChooseSourceID(num);
  };

  const castSourceToOptions = () => {
    const tempOptions: Array<SelectSearchOption> = [];
    tempOptions.push(NOT_FILTER);
    sources.forEach((source) => {
      const tempOption: SelectSearchOption = {
        name: source.title,
        value: source.id,
      };
      tempOptions.push(tempOption);
    });

    
    setOptions([...tempOptions]);
  };

  const requestListAll = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/list/all'
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw LIST_SOURCES_FAIL_MESSAGE;
      }
      setSources(data.articles_sources);
    } catch (error: any) {
      alertError(error);
    }
  };

  return (
    <SelectSearch
      options={options}
      search={true}
      onChange={(e) => handleChooseSource(e)}
      emptyMessage={SEARCH_RESULT_EMPTY}
      placeholder="Choose source"
    />
  );
};

export default DropdownFilterBySources;
