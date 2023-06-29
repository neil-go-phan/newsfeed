import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import CardActions from '@mui/material/CardActions';
import IconButton from '@mui/material/IconButton';
import StarBorderIcon from '@mui/icons-material/StarBorder';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircle, faCircleCheck } from '@fortawesome/free-regular-svg-icons';
import { CARD_MAX_WIDTH } from '.';

type Props = {
  article: Article;
  articlesSource: ArticlesSourceInfo | ArticlesSource | undefined;
  isAdmin: boolean;
  handleModal: () => void;
  content: string;
  readStatus: boolean;
  handleChangeReadStatus: () => void;
};

const ContentCard: React.FC<Props> = (props: Props) => {
  return (
    <Card sx={{ maxWidth: CARD_MAX_WIDTH }} className="articleCard">
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
        <Typography
          className="articleCard__shortContent"
          variant="body1"
          color="text.secondary"
        >
          {props.content}
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

export default ContentCard;
