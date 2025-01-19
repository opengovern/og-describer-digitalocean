package maps



var ResourceTypesToTables = map[string]string{
  "OpenAI/Project": "openai_project",
  "OpenAI/Project/ApiKey": "openai_project_api_key",
  "OpenAI/Project/RateLimit": "openai_project_rate_limit",
  "OpenAI/Project/ServiceAccount": "openai_project_service_account",
  "OpenAI/Project/User": "openai_project_user",
  "OpenAI/Model": "openai_model",
  "OpenAI/File": "openai_file",
  "OpenAI/VectorStore": "openai_vector_store",
  "OpenAI/Assistant": "openai_assistant",
}

var ResourceTypeToDescription = map[string]interface{}{
 
}

var TablesToResourceTypes = map[string]string{
  "openai_project": "OpenAI/Project",
  "openai_project_api_key": "OpenAI/Project/ApiKey",
  "openai_project_rate_limit": "OpenAI/Project/RateLimit",
  "openai_project_service_account": "OpenAI/Project/ServiceAccount",
  "openai_project_user": "OpenAI/Project/User",
  "openai_model": "OpenAI/Model",
  "openai_file": "OpenAI/File",
  "openai_vector_store": "OpenAI/VectorStore",
  "openai_assistant": "OpenAI/Assistant",
}
