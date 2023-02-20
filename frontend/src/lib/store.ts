import { writable, type Writable } from 'svelte/store';

interface GitlabData {
  name: string;
  server: string;
  token: string;
}

export const gitlabData: Writable<GitlabData> = writable({
  name: '',
  server: '',
  token: ''
});

export function updateGitlabData(newData: GitlabData): void {
  gitlabData.set(newData);
}
