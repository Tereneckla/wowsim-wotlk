export const CURRENT_PHASE = 5;

// Github pages serves our site under the /wowsim-wotlk directory (because the repo name is wowsim-wotlk)
export const REPO_NAME = 'wowsim-wotlk';

// Get 'elemental_shaman', the pathname part after the repo name
const pathnameParts = window.location.pathname.split('/');
const repoPartIdx = pathnameParts.findIndex(part => part == REPO_NAME);
export const SPEC_DIRECTORY = repoPartIdx == -1 ? '' : pathnameParts[repoPartIdx + 1];
