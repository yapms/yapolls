import z from "zod";
import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";
import { URL } from "url";

const pollsters_schema = z.array(z.string());
const subjects_schema = z.array(z.object({
	subject: z.string(),
	poll_types: z.array(z.string()),
}));
const poll_types_schema = z.array(z.string());
const polls_schema = z.array(z.object({
	id: z.string(),
	poll_type: z.string(),
	sample_size: z.number().nullable(),
	population: z.string().nullable(),
	url: z.string().nullable(),
	created_at: z.string(),
	start_date: z.string(),
	end_date: z.string(),
	pollster: z.string(),
	answers: z.array(z.object({
		choice: z.string(),
		pct: z.number(),
	})),
	seat_name: z.string().nullable(),
	sponsors: z.array(z.string()),
	internal: z.boolean(),
	partisan: z.string().nullable(),
	subject: z.string(),
}));

export const load: PageServerLoad = async ({ params, url }) => {

	console.log(params, url);

	const pollsters_response = await fetch("https://api.votehub.com/pollsters");
	const pollsters_json = await pollsters_response.json();
	const pollsters = pollsters_schema.safeParse(pollsters_json);
	if (pollsters.success === false) {
		error(500, { message: "Failed to get pollsters." });
	}

	const subjects_response = await fetch("https://api.votehub.com/subjects");
	const subjects_json = await subjects_response.json();
	const subjects = subjects_schema.safeParse(subjects_json);
	if (subjects.success === false) {
		error(500, { message: "failed to get poll subjects." });
	}

	const poll_types_response = await fetch("https://api.votehub.com/poll-types");
	const poll_types_json = await poll_types_response.json();
	const poll_types = poll_types_schema.safeParse(poll_types_json);
	if (poll_types.success === false) {
		error(500, { message: "Failed to get poll types." });
	}

	const pollster = url.searchParams.get("pollster");
	const subject = url.searchParams.get("subject");
	const poll_type = url.searchParams.get("poll_type");

	const polls_url = new URL('https://api.votehub.com/polls');
	polls_url.searchParams.set('pollster', pollster ?? '');
	polls_url.searchParams.set('subject', subject ?? '');
	polls_url.searchParams.set('poll_type', poll_type ?? '');

	const polls_response = await fetch(polls_url);
	const polls_json = await polls_response.json();
	const polls = polls_schema.safeParse(polls_json);
	if (polls.success === false) {
		console.log(polls.error);
		error(500, { message: "Failed to get polling data." });
	}

	return {
		pollsters: pollsters.data,
		subjects: subjects.data,
		poll_types: poll_types.data,
		polls: polls.data,
	};
}

