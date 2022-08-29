import { BrowserRouter, Routes, Route } from "react-router-dom";
import Auth from "./auth";
import Homepage from "./homepage";
import Menu from "./menu";

export default function Router() {

    return <BrowserRouter>
        <Routes>
            <Route
                path="/login"
                element={<Auth />}
            >
            </Route>
            <Route
                path="/"
                element={
                    <Menu viewElement={<Homepage />} />
                }
            >
            </Route>
            <Route exact path="*" element={<h1>Not Found</h1>}>
            </Route>
        </Routes>
    </BrowserRouter>
}