type WindowDimentions = {
  width: number | undefined;
  height: number | undefined;
};

type ChangePasswordFormProperty = {
  password: string;
  password_confirmation?: string;
};

type LoginFormProperty = {
  username: string;
  password: string;
};

type RegisterFormProperty = {
  email: string;
  username: string;
  password: string;
  passwordConfirmation?: string;
};

type UrlFormProperty = {
  url: string;
};

type Article = {
  id: number;
  title: string;
  description: string;
  link: string;
  published: string;
  authors: string;
  articles_source_id: number;
};

type Articles = Array<Article>;

type ArticlesSource = {
  title: string;
  description: string;
  link: string;
  feed_link: string;
  image: string;
  topic_id: number;
};

type ArticlesSourceInfo = {
  id: number;
  title: string;
  description: string;
  link: string;
  image: string;
  follower: number;
  topic_id: number;
};

type ArticlesSourceInfoes = Array<ArticlesSourceInfo>;

type Crawler = {
  source_link: string;
  feed_link: string;
  crawl_type: string;
  article_div: string;
  article_title: string;
  article_description: string;
  article_link: string;
  article_published: string;
  article_authors: string;
  schedule: string;
};

type CreateCrawlerPayload = {
  articles_source: ArticlesSource;
  crawler: Crawler;
};

type Category = {
  id: number;
  name: string;
  illustration: string;
};

type Categories = Array<Category>;

type Topic = {
  id: number;
  name: string;
  category_id: number;
};

type Topics = Array<Topic>;
