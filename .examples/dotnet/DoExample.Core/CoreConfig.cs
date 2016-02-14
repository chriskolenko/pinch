using ConfigLoader;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace DoExample.Core
{
    public class CoreConfig
    {
        public String DB { get; set; }
        public String DBDriver { get; set; }∫

        public static CoreConfig Load()
        {
            return Config.New<CoreConfig>();
        }
    }
}
