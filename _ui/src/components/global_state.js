import { createStore } from "redux";
import { ACTION_LOGIN, ACTION_LOGOUT } from "./state_actions";

var initialState = {
    loggedIn: false,
    user: {},
}

const rootReducer = (state = initialState, action) => {
    switch (action.type) {

        case ACTION_LOGIN:
            return {
                ...initialState,
                loggedIn: true,
                user: action.payload,
            }

        case ACTION_LOGOUT:
            return initialState

        default:
            return state
    }
};

export const store = createStore(rootReducer);
