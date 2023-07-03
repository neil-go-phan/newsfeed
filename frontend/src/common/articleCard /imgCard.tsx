import React from 'react';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import {
  faCircle,
  faCircleCheck,
  faStar,
} from '@fortawesome/free-regular-svg-icons';
import { faStar as starSolid } from '@fortawesome/free-solid-svg-icons';
import { CARD_IMG_HEIGHT, CARD_MAX_WIDTH, CARD_MIN_HEIGHT } from '.';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Popup from 'reactjs-popup';

type Props = {
  article: Article;
  articlesSource: ArticlesSourceInfo | ArticlesSource | undefined;
  isAdmin: boolean;
  handleModal: () => void;
  base64Img: string;
  readStatus: boolean;
  handleChangeReadStatus: () => void;
  isReadLater: boolean;
  handleReadLater: () => void;
};

const ImgCard: React.FC<Props> = (props: Props) => {
  return (
    <Card
      sx={{ maxWidth: CARD_MAX_WIDTH, minHeight: CARD_MIN_HEIGHT }}
      className={props.readStatus ? 'articleCard alreadyRead' : 'articleCard'}
    >
      <CardMedia
        component="img"
        height={CARD_IMG_HEIGHT}
        image={props.base64Img}
        alt="thumbnail"
        onClick={props.handleModal}
      />
      <CardContent onClick={props.handleModal}>
        <Typography
          className="articleCard__title"
          gutterBottom
          variant="h5"
          component="div"
        >
          {props.article.title}
        </Typography>
        <Typography
          className="articleCard__source"
          variant="body2"
          color="text.secondary"
        >
          {props.articlesSource?.title}
        </Typography>
      </CardContent>
      <CardActions
        disableSpacing
        className={props.isAdmin ? 'd-none' : 'd-block'}
      >
        <IconButton aria-label="read later" onClick={props.handleReadLater}>
          {props.isReadLater ? (
            <Popup
              trigger={() => (
                <FontAwesomeIcon icon={starSolid} className="starSolid" />
              )}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Remove article to read later list</span>
            </Popup>
          ) : (
            <Popup
              trigger={() => <FontAwesomeIcon icon={faStar} />}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Add article from read later list</span>
            </Popup>
          )}
        </IconButton>
        <IconButton aria-label="status" onClick={props.handleChangeReadStatus}>
          {props.readStatus ? (
            <Popup
              trigger={() => <FontAwesomeIcon icon={faCircleCheck} />}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Mark article as unread</span>
            </Popup>
          ) : (
            <Popup
              trigger={() => <FontAwesomeIcon icon={faCircle} />}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Mark article as read</span>
            </Popup>
          )}
        </IconButton>
      </CardActions>
    </Card>
  );
};

export default ImgCard;
