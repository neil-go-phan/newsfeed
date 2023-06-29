import React from 'react';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import StarBorderIcon from '@mui/icons-material/StarBorder';
import { CARD_IMG_HEIGHT, CARD_MAX_WIDTH } from '.';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircle, faCircleCheck } from '@fortawesome/free-regular-svg-icons';

type Props = {
  article: Article;
  articlesSource: ArticlesSourceInfo | ArticlesSource | undefined;
  isAdmin: boolean;
  handleModal: () => void;
  base64Img: string;
  readStatus: boolean;
  handleChangeReadStatus: () => void;
};

const ImgCard: React.FC<Props> = (props: Props) => {
  return (
    <Card sx={{ maxWidth: CARD_MAX_WIDTH }} className={props.readStatus ? "articleCard alreadyRead" : "articleCard"}>
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
        <IconButton aria-label="read later">
          <StarBorderIcon />
        </IconButton>
        <IconButton aria-label="status" onClick={props.handleChangeReadStatus}>
          {props.readStatus ? (
            <FontAwesomeIcon icon={faCircleCheck} />
          ) : (
            <FontAwesomeIcon icon={faCircle} />
          )}
        </IconButton>
      </CardActions>
    </Card>
  );
};

export default ImgCard;
