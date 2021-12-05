import { configureStore, combineReducers } from "@reduxjs/toolkit";
import authSlice from "./slices/auth.slice";
import storage from "redux-persist/lib/storage";
import {
  persistStore,
  persistReducer,
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
} from "redux-persist";
import { alertReducer } from "./slices/alert.slice";
import { loaderReducer } from "./slices/loader.slice";

const authPersistConfig = {
  key: "auth",
  storage
}

const settingsPersistConfig = {
  key: "settings",
  storage
}

const rootReducer = combineReducers({
  auth: persistReducer(authPersistConfig, authSlice),
  alert: alertReducer,
  loader: loaderReducer,
});

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export const persistor = persistStore(store);
