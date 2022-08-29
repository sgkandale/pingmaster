import { createStore } from "redux";
import { ACTION_LOGIN } from "./state_actions";

var initialState = {
    loggedIn: false,
    user: {},
}

const rootReducer = (state = initialState, action) => {
    switch (action.type) {

        case ACTION_LOGIN:
            console.log("action.payload", action.payload)
            return {
                ...initialState,
                loggedIn: true,
                user: action.payload,
            }

        default:
            return state
    }
};

export const store = createStore(rootReducer);
