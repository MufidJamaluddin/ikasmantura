import { Model, middlewares } from "react-model"

import AboutModel from "./AboutModel";
import ArticleTopicModel from "./ArticleTopicModel";
import ArticleModel from "./ArticleModel";

const models = { AboutModel, ArticleTopicModel, ArticleModel }

middlewares.config.logger.enable = false

export const { getInitialState, useStore, getState, actions, subscribe, unsubscribe } = Model(models)
