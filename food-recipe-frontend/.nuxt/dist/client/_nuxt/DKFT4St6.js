function e({store:t,redirect:a}){if(!t.state.auth.isAuthenticated||t.state.auth.user.role!=="admin")return a("/")}export{e as default};
