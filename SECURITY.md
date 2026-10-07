# Security Policy

## Reporting a vulnerability

Please do not open a public issue for a security problem. Use GitHub's private
[security advisory](https://github.com/wtb3771c30h/AegisTrace/security/advisories/new) form, or contact a
maintainer directly if the form is unavailable to you.

Include what an attacker gains in one sentence, the version or commit you are running, the smallest
reproduction you can manage, and whether the input is attacker-controlled in your setup. You should
get an acknowledgement within a few working days and a disclosure timeline before anything is
published.

## Scope

AegisTrace compiles policy-bound workflows and records what executing them did. In
scope: a recorded effect that replays differently (clock reads, map ordering, env
access), an artifact that still verifies after being edited, a policy check the
compiler can be talked past, and path handling in the recorder.
