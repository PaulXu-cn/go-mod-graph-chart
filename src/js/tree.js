import * as d3 from 'd3';
import CSS from '../css/styles.css';

var treeWidth = 1;
var treeDepth = 1;

var width = 4000;
var height = 2500;

var innerWidth = 0
var innerHeight = 0;

var root;
var color;

const margin = { top: 5, right: 200, bottom: 5, left: 200};

export default function tree() {
    const svg = d3.select('#app').append('svg').attr('id', 'svg')
    .attr('width', width).attr('height', height);

    const g = svg.append('g')
        .attr('transform', `translate(${margin.left}, ${margin.top})`);

    const fill = d => {
        if (d.depth === 0)
            return color(d.data.name)
        else
            return color(d.parent.data.name)
    }

    const render = function (data) {
        color = d3.scaleOrdinal()
            .domain(root.descendants().filter(d => d.depth <= 1).map(d => d.data.name))
            .range(d3.schemeCategory10);

        g.selectAll("path")
            .data(root.links())
            .join("path")
            .classed('tree-path',true)
            .attr("d", d3.linkHorizontal().x(d => d.y).y(d => d.x));

        g.selectAll('circle').data(root.descendants()).join('circle')
            // optionally, we can use stroke-linejoin to beautify the path connection; 
            //.attr("stroke-linejoin", "round")
            .classed('tree-circle', true)
            .attr("fill", fill)
            .attr('id', (d, i) => `circle-${i}`)
            .attr('cx', d => d.y)
            .attr('cy', d => d.x)
            .attr("r", function (d) {
                if (d.children) {
                    return 6 + Math.sqrt(d.children.length) * 3;
                } else {
                    return 7;
                }
            }).on('mouseover', function (d, i) {
                d3.select(this).classed('on', true);
                d3.select("#text-" + i).classed('on', true);
            }).on('mouseout', function (d, i) {
                d3.select(this).classed('on', false).attr('stroke', fill(d));
                d3.select("#text-" + i).classed('on', false);
            });

        g.selectAll('text').data(root.descendants()).join('text')
            .classed('tree-text', true)
            .attr("text-anchor", d => d.children ? "end" : "start")
            .attr('id', (d, i) => `text-${i}`)
            // note that if d is a child, d.children is undefined which is actually false! 
            .attr('x', d => (d.children ? -6 : 6) + d.y)
            .attr('y', function (d) {return d.depth % 2 ? d.x - 3: d.x + 14})
            .text(d => d.data.name)
            .on('mouseover', function (d, i) {
                d3.select(this).classed('on', true);
                d3.select("#circle-" + i).classed('on', true);
            }).on('mouseout', function (d, i) {
                d3.select(this).classed('on', false);
                d3.select("#circle-" + i).classed('on', false).attr('stroke', fill(d));
            });
    }

    d3.json('/tree.json').then(data => {
        let treeData = data.data.tree;
        treeWidth = data.data.width;
        treeDepth = data.data.depth;

        innerWidth = treeDepth * 450;
        innerHeight = treeWidth * 70;
    
        width = innerWidth + margin.left + margin.right;
        height = innerHeight + margin.top + margin.bottom;

        d3.select('#app').select('svg').attr('width', width).attr('height', height);

        root = d3.hierarchy(treeData);
        // alternatively, we can set size of each node; 
        // root = d3.tree().nodeSize([30, width / (root.height + 1)])(root);
        root = d3.tree().size([innerHeight, innerWidth])(root);
        render(root);

        setTimeout(function () {
            window.scrollTo(0, height / 2)
        }, 1000);
    });
}
