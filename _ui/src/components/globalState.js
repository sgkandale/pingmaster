import { createStore } from "redux";

var initialState = {
}

const rootReducer = (state = initialState, action) => {
    switch (action.type) {

        default:
            return state
    }
};

export const store = createStore(rootReducer);
