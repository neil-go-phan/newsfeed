import React, { useEffect, useRef, useState } from 'react';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import { parse } from 'parse5';


type Props = {
  article: Article;
  handleContentModalClose: () => void;
};

const ContentModal: React.FC<Props> = (props: Props) => {
  const contentRef = useRef<any>(null);
  const [dom, setDom] = useState<Document>();
  useEffect(() => {
    if (props.article.description !== '') {
      const doc = parse(props.article.description)
      const nodes = doc.childNodes;
      
      // const newdom = htmlparser2.parseDocument(props.article.description);
      // const fragment = document
      //   .createRange()
      //   .createContextualFragment(props.article.description);
      // if (contentRef.current) {
      //   contentRef.current.appendChild(fragment);
      // }
      // const links = domutils.getElementsByTagName('a', fragment);
      // links.forEach((link) => {
      //   // Kiểm tra xem phần tử a có thuộc tính href không
      //   const href = link.getAttributeValue(link, 'href');
      //   if (href) {
      //     // Thêm thuộc tính target="_blank"
      //     link.setAttribute(link, 'target', '_blank');
      //   }
      // });
    }
  }, [props.article]);

  useEffect(() => {
    if (dom) {
    }
  }, [dom]);

  return (
    <div className="articleCard__contentModal">
      <div className="content">
        <div ref={contentRef}></div>
      </div>
    </div>
  );
};

export default ContentModal;
