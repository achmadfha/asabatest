export function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let ca = decodedCookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

export function setCookie(cname, cvalue, days) {
    let date = new Date();
    date.setDate(date.getDate() + days);
    let value = cvalue + ((days == null) ? "" : "; expires=" + date.toUTCString());
    document.cookie = cname + "=" + value;
}

export function deleteCookie(cname) {
    document.cookie = cname + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
}

export function getUser() {
    let accessToken = getCookie('access_token');
    return !!accessToken;
}

export function setUser(user) {
    if (user && user.access_token) {
        setCookie('access_token', user.access_token, 1);
    }
}

export function clearUser() {
    deleteCookie('access_token');
}

export function getAccessToken() {
    return getCookie('access_token');
}