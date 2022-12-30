export const ValidateProps = {
  user: {
    username: { type: "string", minLength: 4, maxLength: 20 },
    name: { type: "string", minLength: 1, maxLength: 50 },
    password: { type: "string", minLength: 8 },
    email: { type: "string", minLength: 1 },
    githubApiToken: { type: "string", minLength: 0, maxLength: 40 },
    railwayApiToken: { type: "string", minLength: 0, maxLength: 36 },
    renderApiToken: { type: "string", minLength: 0, maxLength: 32 },
    renderUserEmail: { type: "string", minLength: 0 },
    isAdmin: { type: "boolean" },
    projects: { type: "array" },
  },
  project: {
    name: { type: "string", minLength: 1, maxLength: 350 },
    visibility: { type: "string", minLength: 6, maxLength: 7 },
    platform: { type: "string", minLength: 5, maxLength: 8 },
    lang: { type: "string", minLength: 1, maxLength: 50 },
    packageManager: { type: "string", minLength: 1, maxLength: 50 },
    hostService: { type: "string", minLength: 1, maxLength: 50 },
    botToken: { type: "string", minLength: 0, maxLength: 100 },
    botAppToken: { type: "string", minLength: 0, maxLength: 100 },
    botSecretToken: { type: "string", minLength: 0, maxLength: 100 },
    railwayProjectId: { type: "string", minLength: 0, maxLength: 100 },
    railwayServiceId: { type: "string", minLength: 0, maxLength: 100 },
    railwayEnvId: { type: "string", minLength: 0, maxLength: 100 },
    renderProjectId: { type: "string", minLength: 0, maxLength: 24 },
  },
};
