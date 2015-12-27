using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Web.Http;
using DoExample.Core;

namespace DoExample.Api.Controllers
{
    public class StatsController : ApiController
    {
        public string GetStats()
        {
            var cfg = CoreConfig.Load();
            return cfg.DB;
        }
    }
}
