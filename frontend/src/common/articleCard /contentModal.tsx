import { alertNotFoundLink } from '@/helpers/alert';
import React, { useEffect, useRef, useState } from 'react';
import { ColorRing } from 'react-loader-spinner';

type Props = {
  article: Article;
  articlesSource: ArticlesSourceInfo | ArticlesSource | undefined;

  handleContentModalClose: () => void;
  doc: any;
};

const ContentModal: React.FC<Props> = (props: Props) => {
  const [renderedContent, setRenderedContent] = useState<
    Array<JSX.Element | string | null>
  >([]);
  const [customCrawlerArticleDescription, setCustomCrawlerArticleDescription] =
    useState<string>('Not found');
  const [isCustomCrawler, setIsCustomCrawler] = useState<boolean>(false);
  const [str, setStr] = useState<string>('');

  const renderNode = (node: any): JSX.Element | string | null => {
    const handleImageLoad = (event: React.SyntheticEvent<HTMLImageElement>) => {
      const imgElement = event.target as HTMLImageElement;
      imgElement.style.display = 'block';
      const loadingDivs = document.getElementsByClassName(imgElement.src);
      for (let i = 0; i < loadingDivs.length; i++) {
        const loadingDiv = loadingDivs[i] as HTMLElement;
        loadingDiv.style.display = 'none';
      }
    };
    if (node.type === 'tag') {
      const { name, children, attribs } = node;
      const props: { [key: string]: string } = Object.entries(attribs).reduce(
        (props, [key, value]) => ({ ...props, [key]: value }),
        {}
      );
      const childElements = children.map((child: any) => renderNode(child));
      if (name === 'img') {
        return React.createElement('div', null, [
          React.createElement('div', { className: props.src }, loadingImg),
          React.createElement(
            name,
            { ...props, onLoad: handleImageLoad },
            ...childElements
          ),
        ]);
      }
      return React.createElement(name, props, ...childElements);
    } else if (node.type === 'text') {
      return node.data;
    }

    return null;
  };

  // const handleImageLoad = (event: React.SyntheticEvent<HTMLImageElement>) => {
  //   const imgElement = event.target as HTMLImageElement;
  //   imgElement.style.display = 'block';
  // };

  const renderDate = (dateString: string): string => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  const handleOpenUrl = () => {
    if (props.article.link) {
      window.open(props.article.link, '_blank');
    } else {
      alertNotFoundLink('Not found link to article');
    }
  };

  useEffect(() => {
    if (props.doc) {
      if (props.doc.childNodes) {
        const temp: Array<JSX.Element | string | null> = [];
        props.doc.childNodes.forEach((node: any) => {
          temp.push(renderNode(node));
        });
        setRenderedContent(temp);
        if (props.article.published) {
          const dateString = renderDate(props.article.published);
          setStr(dateString);
        }
      }
    } else {
      if (props.article.description) {
        setIsCustomCrawler(true);
        setCustomCrawlerArticleDescription(props.article.description);
      }
    }
  }, []);

  return (
    <div className="articleCard__contentModal">
      <div className="title">
        <p onClick={handleOpenUrl}>{props.article.title}</p>
      </div>
      <div className="info">
        <a href={props.articlesSource?.link} target="_blank" className="source">
          {props.articlesSource?.title ? `${props.articlesSource.title}, ` : ''}
        </a>
        <span className="authors">
          {props.article?.authors ? `by ${props.article.authors}, ` : ''}
        </span>
        <span className="published">{str}</span>
      </div>
      {isCustomCrawler ? (
        <p>{customCrawlerArticleDescription}</p>
      ) : (
        renderedContent
      )}
    </div>
  );
};

export default ContentModal;

const loadingImg = (
  <div
    className="loadingImg"
    style={{
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      minHeight: '100px',
    }}
  >
    <ColorRing
      visible={true}
      height="100"
      width="100"
      ariaLabel="blocks-loading"
      wrapperStyle={{}}
      wrapperClass="blocks-wrapper"
      colors={['#e15b64', '#f47e60', '#f8b26a', '#abbd81', '#849b87']}
    />
  </div>
);
