let backend = process.env.REACT_APP_BACKEND;

if (!backend) {
    backend = "localhost:8080";
}

export const BACKEND : string = backend.indexOf("http") > -1 ? backend : "http://" + backend;
