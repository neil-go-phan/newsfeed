import React, { useEffect, useState } from 'react';

type Props = {
  article: Article;
  sourceTitle: string | undefined;
  sourceLink: string | undefined;
  handleContentModalClose: () => void;
  doc: any;
};

const ContentModal: React.FC<Props> = (props: Props) => {
  const [renderedContent, setRenderedContent] = useState<
    Array<JSX.Element | string | null>
  >([]);
  const [str, setStr] = useState<string>('');
  const renderNode = (node: any): JSX.Element | string | null => {
    if (node.type === 'tag') {
      const { name, children, attribs } = node;
      const props: { [key: string]: string } = Object.entries(attribs).reduce(
        (props, [key, value]) => ({ ...props, [key]: value }),
        {}
      );
      const childElements = children.map((child: any) => renderNode(child));

      return React.createElement(name, props, ...childElements);
    } else if (node.type === 'text') {
      return node.data;
    }

    return null;
  };
  const renderDate = (dateString: string): string => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };
  useEffect(() => {
    const temp: Array<JSX.Element | string | null> = [];
    props.doc.childNodes.forEach((node: any) => {
      temp.push(renderNode(node));
    });
    setRenderedContent(temp);
    if (props.article.published) {
      const dateString = renderDate(props.article.published);
      setStr(dateString);
    }
  }, []);
  return (
    <div className="articleCard__contentModal">
      <div className="title">
        <a href={props.article.link} target="_blank">
          {props.article.title}
        </a>
      </div>
      <div className="info">
        <a href={props.sourceLink} target="_blank" className="source">
          {props.sourceTitle},{' '}
        </a>
        <span className="authors">by {props.article.authors}, </span>
        <span className="published">{str}</span>
      </div>
      {renderedContent}
    </div>
  );
};

export default ContentModal;
