# OpenAI Integration Guide

Learn how to integrate opencomply with your OpenAI project.

This guide shows you how to connect opencomply to an OpenAI project using a project-level API key with read-only access. Each integration corresponds to a single OpenAI project. Once integrated, opencomply will discover and assess key OpenAI resources—such as assistants, files, models, projects, project API keys, rate limits, service accounts, users, and vector stores—enabling compliance and visibility across your OpenAI environment.

## Prerequisites

- opencomply installed and running
- An OpenAI project with the appropriate read-only API key permissions

## Create a Read-Only Project API Key

1. Log in to the OpenAI Dashboard and select the relevant project.
2. Go to **Settings** or **Manage Keys** (location may vary depending on the dashboard version).
3. Click **Create New Key**, set it to **Read-Only**, and copy the generated key.

## Configure Integration in opencomply

1. In the opencomply dashboard, go to **Integrations > OpenAI**.
2. Select **API Key integration**.
3. Paste in your OpenAI project-level API key.
4. Specify the OpenAI project to govern.
5. Click **Save**.