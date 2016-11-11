# Perms

An extremely flexible permissions system

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Nodes](#nodes)
  - [Important Considerations](#important-considerations)
  - [Wildcards](#wildcards)
  - [Negations](#negations)
- [Basic Usage](#basic-usage)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Nodes

A permission node grants a particular kind of access to a resource.

It should be formatted as follows

`[-]<resource>[.<sub_resource>...].<verb>`

The following are examples of good nodes

- `projects.webserver.build`
- `projects.webserver.chat.use`
- `projects.webserver.chat.moderate`


### Important Considerations

- Nodes are case sensitive
- Whitespaces are not allowed

### Wildcards

An asterisk or `*` may be used to signify a wildcard match.

The `*` node would match every permission


The `projects.*` node would match for

- `projects.*`
- `projects.webserver.test`
-  etc.

The `projects.*.chat.use` node would match for

- `projects.webserver.chat.use`
- `projects.database.chat.use`
- `projects.client.chat.use`

### Negations

A `-` prefixing a node signifies it's negated. Negation will be only be overwritten if the
user has the node explicitely defined. If the negation lies directly on the user, it must be removed
for the user to have access to the permission.


if Bob has 
- `*`
- `-projects.*`

He has access to

- `billing.budget.manage`

But not 

- `projects.webserver.use`


## Basic Usage

Permissions takes one or more permissions configuration or `PConf`s and assembles a 
`PermissionsWeb`


`PConf`s are JSON and can look like any of the following.


```js
{
    "groups": {
        "project_lead": {
            "nodes": [
                "analytics.*"
            ]
        }
        "manager": {
            "parents": [
                "project_lead"
            ],
            "nodes": [
                "projects.*"
            ]
        }
    }
}
```

```js
{
    "users": {
        "ammar": {
            "groups": [
                "manager",
            ],
            "nodes": [
                "-projects.*.chat.moderate",
            ],
        }
    }
}
```

or even

```js
	{
    "groups": {
        "project_lead": {
            "nodes": [
                "analytics.*"
            ]
        },
        "manager": {
            "parents": [
                "project_lead"
            ],
            "nodes": [
                "projects.*"
            ]
        }
    },
    "users": {
        "ammar": {
            "groups": [
                "manager"
            ],
            "nodes": [
                "-projects.*.chat.moderate"
            ]
        }
    }
}
```

Pconfs can stack up. If multiple pconfs provide `users` over and over again, the loaded users 
will __not__ be overwritten. Users and groups must be explicitely deleted.