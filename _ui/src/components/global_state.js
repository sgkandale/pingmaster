import { createStore } from "redux";
import { ACTION_LOGIN, ACTION_LOGOUT, ACTION_REMOVE_TARGETS } from "./state_actions";

var initialState = {
    loggedIn: false,
    user: {},
    targets: [],
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

        case ACTION_REMOVE_TARGETS:
            return {
                ...state,
                targets: [],
            }

        default:
            return state
    }
};

export const store = createStore(rootReducer);
