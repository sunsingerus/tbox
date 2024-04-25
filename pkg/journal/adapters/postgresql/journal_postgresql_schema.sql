CREATE TABLE IF NOT EXISTS api_journal
(
  /* d is an ordering/grouping key */
  d TIMESTAMP DEFAULT now(),

  /**********************/
  /**** Call section ****/
  /**********************/

  /* EndpointID [MANDATORY] specifies ID of the endpoint (API call handler/Task processor/etc) which produces the entry */
  endpoint_id INTEGER,

  /* EndpointInstanceID [OPTIONAL] specifies ID of the particular endpoint instance (ex.: process) which produces the entry */
  endpoint_instance_id TEXT,

  /* SourceID [OPTIONAL] specifies ID of the source (possibly external) of the entry */
  source_id TEXT,

  /* ContextID [OPTIONAL] specifies ID of the execution/rpc context associated with the entry */
  context_uid TEXT,

  /* TaskID [OPTIONAL] specifies ID of the task associated with the entry */
  task_uid TEXT,

  /* Type [MANDATORY] specifies type of the entry - what this entry is about. */
  type_id INTEGER,

  /* duration is a nanoseconds duration since start, if applicable */
  duration BIGINT,

  /************************/
  /**** Object section ****/
  /************************/

  /* type specifies object type, if any */
  type INTEGER,

  /* size specifies object size, if any */
  size BIGINT,

  /* address specifies object address, if any */
  address TEXT,

  /* domain specifies object domain, if any */
  domain TEXT,

  /* name specifies object name, if any */
  name TEXT,

  /* digest specifies object digestm if any*/
  digest TEXT,

  /* data specifies object data, if any */
  data TEXT,

  /************************/
  /**** Result section ****/
  /************************/

  /* Result [OPTIONAL] specifies result of the operation, specified by Type */
  result TEXT,
  /* Result [OPTIONAL] specifies status of the operation, specified by Type */
  status TEXT,
  /* Error [OPTIONAL] tells about error encountered during the operation, specified by Type */
  error TEXT
)
;

CREATE INDEX IF NOT EXISTS api_journal_index ON api_journal (d);
